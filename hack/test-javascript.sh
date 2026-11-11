#!/usr/bin/env bash
# Checks that the generated JavaScript package works for a consumer.
#
# Installs gen/javascript into a throwaway project, then for every entry point
# in its package.json checks that `import` and `require` give the same names
# and values, and type-checks the declarations of both formats.
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

readonly package="${PWD}/gen/javascript"
readonly typescript="typescript@${TYPESCRIPT_VERSION:-5.9.3}"

dir="$(mktemp -d)"
trap 'rm -rf "${dir}"' EXIT
cd "${dir}"

printf '{"name": "consumer", "private": true}\n' >package.json
npm install --silent --no-audit --no-fund "${package}" "${typescript}"

cat >check.mjs <<'SCRIPT'
import { readFileSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';

const require = createRequire(import.meta.url);
const { name, exports } = require('sigsemconv/package.json');
const entries = Object.keys(exports)
  .filter((path) => path !== './package.json')
  .map((path) => name + path.slice(1));

// A template attribute is a function of the key; compare what it returns.
const values = (module) =>
  JSON.stringify(
    Object.entries(module)
      .filter(([key]) => key !== 'default' && key !== 'module.exports')
      .map(([key, value]) => [key, typeof value === 'function' ? value('key') : value])
      .sort(),
  );

for (const entry of entries) {
  const esm = values(await import(entry));
  const cjs = values(require(entry));
  if (esm === '[]') throw new Error(`${entry} exports nothing`);
  if (esm !== cjs) throw new Error(`${entry}: import and require give different names or values`);
}

// One file per declaration format that imports every entry point.
const imports = (format) =>
  entries
    .map((entry, i) =>
      format === 'mts' ? `import * as e${i} from '${entry}';` : `import e${i} = require('${entry}');`,
    )
    .concat(`export const all = [${entries.map((_, i) => `e${i}`).join(', ')}];`)
    .join('\n');
writeFileSync('types.mts', imports('mts') + '\n');
writeFileSync('types.cts', imports('cts') + '\n');
console.log(`checked ${entries.length} entry points`);
SCRIPT

node check.mjs
npx --no-install tsc --noEmit --strict --skipLibCheck false \
  --module node16 --moduleResolution node16 types.mts types.cts
