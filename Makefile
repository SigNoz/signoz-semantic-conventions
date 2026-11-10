.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# The Python formatter for gen/python/.
RUFF ?= ruff

.PHONY: yaml-lint
yaml-lint: ## Lint the yaml files
	@yamllint .

# Upstream rego policies, pinned. names.rego is replaced by
# policies/after_resolution/names.rego, which supports attribute exceptions.
WEAVER_PACKAGES := https://github.com/open-telemetry/opentelemetry-weaver-packages.git@92eba8ea6d06c6b60a3386917858443816977810
WEAVER_PACKAGES_POLICIES := \
	policies/check/stability/stability.rego \
	policies/check/stability/deprecation.rego \
	policies/check/entity_associations/entity_associations.rego \
	policies/check/naming_conventions/attribute_name_collisions.rego \
	policies/check/naming_conventions/attribute_types.rego \
	policies/check/naming_conventions/constant_name_collisions.rego \
	policies/check/naming_conventions/enum_member_collisions.rego \
	policies/check/naming_conventions/metric_brief_format.rego \
	policies/check/naming_conventions/metrics_collisions.rego

.PHONY: weaver-check
weaver-check: ## Runs weaver check with the rego policies.
	@dups="$$(grep -rhE '^- id: |^  metric_name: ' model | sed -E 's/ *#.*//; s/["'"'"']//g; s/ +$$//' | sort | uniq -d)"; \
		if [ -n "$$dups" ]; then echo "Defined more than once in model/:"; echo "$$dups"; exit 1; fi
	@weaver registry check --registry model/ --policy policies/before_resolution
	@weaver registry check --registry model/ --v2 --policy policies/after_resolution \
		$(foreach policy,$(WEAVER_PACKAGES_POLICIES),--policy '$(WEAVER_PACKAGES)[$(policy)]')

.PHONY: weaver-docs
weaver-docs: ## Generate docs/ from the registry.
	@rm -rf docs
	@weaver registry generate --registry model/ --v2 --quiet markdown docs

.PHONY: weaver-gen
weaver-gen: weaver-go weaver-python ## Generate gen/ from the registry.

# A template that has nothing to render for an area leaves an empty file, so
# empty files are deleted after generating.
.PHONY: weaver-go
weaver-go: ## Generate gen/go/ from the registry.
	@rm -rf gen/go/sigsemconv
	@weaver registry generate --registry model/ --quiet go gen/go/sigsemconv
	@find gen/go/sigsemconv -empty -delete
	@gofmt -w gen/go/sigsemconv

.PHONY: weaver-python
weaver-python: ## Generate gen/python/ from the registry.
	@rm -rf gen/python/sigsemconv
	@weaver registry generate --registry model/ --quiet python gen/python/sigsemconv
	@find gen/python/sigsemconv -empty -delete
	@touch gen/python/sigsemconv/py.typed
	@$(RUFF) format --quiet gen/python/sigsemconv
