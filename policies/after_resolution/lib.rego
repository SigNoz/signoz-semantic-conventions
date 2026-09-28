package after_resolution

import rego.v1

# An attribute or signal is excused from `rule` when its annotations list it:
#
#   annotations:
#     signoz:
#       policy_exceptions:
#       - counter_unit_one
signoz_excepted(obj, rule) if {
	some exception in object.get(obj, ["annotations", "signoz", "policy_exceptions"], [])
	exception == rule
}

signoz_attribute_finding(attr, rule, message) := {
	"id": rule,
	"message": message,
	"level": "violation",
	"context": {"attribute": attr.key},
}

# `s` is {"type": ..., "name": ..., "signal": ...}.
signoz_signal_finding(s, rule, message) := {
	"id": rule,
	"message": message,
	"level": "violation",
	"context": {},
	"signal_type": s.type,
	"signal_name": s.name,
}

signoz_stability_rank := {
	"development": 0,
	"alpha": 1,
	"beta": 2,
	"release_candidate": 3,
	"stable": 4,
}
