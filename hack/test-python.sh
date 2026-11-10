#!/usr/bin/env bash
# Checks that the generated Python package works for a consumer.
#
# Installs gen/python into a throwaway virtual environment, imports every
# module with warnings as errors, and creates every metric instrument.
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

readonly package="${PWD}/gen/python"

dir="$(mktemp -d)"
trap 'rm -rf "${dir}"' EXIT

python3 -m venv "${dir}/venv"
"${dir}/venv/bin/pip" install --quiet --disable-pip-version-check "${package}"

# Run outside the repository so that the installed package is the one imported.
cd "${dir}"
"${dir}/venv/bin/python" -W error - <<'SCRIPT'
import importlib
import importlib.metadata
import inspect
import pkgutil

import sigsemconv
from opentelemetry.metrics import get_meter
from sigsemconv.version import __version__

if importlib.metadata.version("sigsemconv") != __version__:
    raise SystemExit("the installed version and sigsemconv.version differ")

meter = get_meter("test")
modules = 0
instruments = 0
for info in pkgutil.walk_packages(sigsemconv.__path__, "sigsemconv."):
    module = importlib.import_module(info.name)
    modules += 1
    for name, function in inspect.getmembers(module, inspect.isfunction):
        if not name.startswith("create_"):
            continue
        # An observable instrument also takes its callbacks.
        if "callbacks" in inspect.signature(function).parameters:
            function(meter, None)
        else:
            function(meter)
        instruments += 1

if modules == 0:
    raise SystemExit("sigsemconv has no modules")
print(f"checked {modules} modules and {instruments} instruments")
SCRIPT
