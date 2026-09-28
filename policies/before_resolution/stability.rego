package before_resolution

import rego.v1

# Every signal, attribute definition and enum member states its stability.
# Weaver only warns when one is missing, so the warning never fails a check.

deny contains finding if {
	some group in input.groups
	group.type in signoz_signal_types
	not group.stability
	finding := signoz_finding(group, "missing_stability", sprintf(
		"Group '%s' has no 'stability'. New conventions start at 'development'.",
		[group.id],
	))
}

deny contains finding if {
	some group in input.groups
	signoz_is_registry(group)
	some attr in object.get(group, "attributes", [])
	attr.id
	not attr.stability
	finding := signoz_finding(group, "missing_stability", sprintf(
		"Attribute '%s' has no 'stability'. New conventions start at 'development'.",
		[attr.id],
	))
}

deny contains finding if {
	some group in input.groups
	signoz_is_registry(group)
	some attr in object.get(group, "attributes", [])
	some member in object.get(attr, ["type", "members"], [])
	not member.stability
	finding := signoz_finding(group, "missing_stability", sprintf(
		"Value '%s' of attribute '%s' has no 'stability'.",
		[member.id, attr.id],
	))
}
