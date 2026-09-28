package before_resolution

import rego.v1

# A group or attribute is excused from `rule` when its annotations list it:
#
#   annotations:
#     signoz:
#       policy_exceptions:
#       - entity_role_implicit
signoz_excepted(obj, rule) if {
	some exception in object.get(obj, ["annotations", "signoz", "policy_exceptions"], [])
	exception == rule
}

signoz_finding(group, rule, message) := {
	"id": rule,
	"message": message,
	"level": "violation",
	"context": {"group": group.id},
}

# Only `registry.*` attribute groups define attributes; every other group refs them.
signoz_is_registry(group) if {
	group.type == "attribute_group"
	startswith(group.id, "registry.")
}

# Weaver reads `type: resource` as `entity` before policies run.
signoz_signal_types := {"span", "metric", "event", "entity"}

signoz_name_regex := `^[a-z][a-z0-9]*([._][a-z0-9]+)*$`
