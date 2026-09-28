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

# Weaver reads `deprecated: "text"` as reason `unspecified`; a deprecation needs a structured reason.
signoz_deprecations contains {"what": sprintf("Group '%s'", [group.id]), "deprecated": group.deprecated, "group": group} if {
	some group in input.groups
	group.deprecated
}

signoz_deprecations contains {"what": sprintf("Attribute '%s'", [attr.id]), "deprecated": attr.deprecated, "group": group} if {
	some group in input.groups
	some attr in object.get(group, "attributes", [])
	attr.id
	attr.deprecated
}

deny contains finding if {
	some d in signoz_deprecations
	d.deprecated.reason == "unspecified"
	finding := signoz_finding(d.group, "deprecated_unstructured", sprintf(
		"%s has an unstructured 'deprecated'. Use a mapping with 'reason: renamed | obsoleted | uncategorized'.",
		[d.what],
	))
}
