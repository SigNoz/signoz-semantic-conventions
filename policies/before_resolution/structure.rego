package before_resolution

import rego.v1

# How groups are laid out in the model files.

signoz_id_prefixes := {
	"attribute_group": ["registry.", "attributes.", "metric_attributes.", "trace.", "event_attributes."],
	"span": ["span."],
	"metric": ["metric."],
	"event": ["event."],
	"entity": ["entity."],
}

deny contains finding if {
	some group in input.groups
	not signoz_is_registry(group)
	some attr in object.get(group, "attributes", [])
	attr.id
	not signoz_excepted(group, "attr_defined_outside_registry")
	finding := signoz_finding(group, "attr_defined_outside_registry", sprintf(
		"Attribute '%s' is defined inside '%s'. Define attributes in a 'registry.*' group and reference them here with '- ref: %s'.",
		[attr.id, group.id, attr.id],
	))
}

deny contains finding if {
	some group in input.groups
	group.type == "undefined"
	finding := signoz_finding(group, "group_type_missing", sprintf(
		"Group '%s' has no 'type'.",
		[group.id],
	))
}

deny contains finding if {
	some group in input.groups
	object.get(group, "prefix", "") != ""
	finding := signoz_finding(group, "prefix_used", sprintf(
		"Group '%s' uses 'prefix', which is obsolete. Write fully qualified attribute ids.",
		[group.id],
	))
}

deny contains finding if {
	some group in input.groups
	not regex.match(signoz_name_regex, group.id)
	finding := signoz_finding(group, "group_id_format", sprintf(
		"Group id '%s' must be lowercase words separated by '.' or '_'.",
		[group.id],
	))
}

deny contains finding if {
	some group in input.groups
	prefixes := signoz_id_prefixes[group.type]
	not signoz_has_prefix(group.id, prefixes)
	finding := signoz_finding(group, "group_id_convention", sprintf(
		"Group '%s' is a '%s'; its id starts with one of %v.",
		[group.id, group.type, prefixes],
	))
}

signoz_has_prefix(id, prefixes) if {
	some prefix in prefixes
	startswith(id, prefix)
}

deny contains finding if {
	some group in input.groups
	group.type == "span"
	not group.span_kind
	finding := signoz_finding(group, "span_missing_kind", sprintf(
		"Span '%s' has no 'span_kind'. Set one of client, server, producer, consumer or internal.",
		[group.id],
	))
}

# Weaver silently applies a ref's stability or deprecation to the signal, so a ref never sets them.
deny contains finding if {
	some group in input.groups
	some attr in object.get(group, "attributes", [])
	attr.ref
	some field in {"stability", "deprecated"}
	object.get(attr, field, null) != null
	finding := signoz_finding(group, "ref_definition_fields", sprintf(
		"'- ref: %s' in '%s' sets '%s'. Set it on the attribute's definition instead.",
		[attr.ref, group.id, field],
	))
}
