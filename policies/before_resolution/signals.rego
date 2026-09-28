package before_resolution

import rego.v1

# Checked on the raw files because weaver's resolved (v2) registry drops
# event bodies and `sampling_relevant` on non-span attributes.

# Event data goes in attributes. A body is for a display message at most.
deny contains finding if {
	some group in input.groups
	group.type == "event"
	group.body
	not signoz_excepted(group, "event_body")
	finding := signoz_finding(group, "event_body", sprintf(
		"Event '%s' defines a body. Put its data in attributes.",
		[object.get(group, "name", group.id)],
	))
}

# Only signals: `trace.*` and other shared groups that spans extend may set it.
deny contains finding if {
	some group in input.groups
	group.type in {"metric", "event", "entity"}
	some attr in object.get(group, "attributes", [])
	object.get(attr, "sampling_relevant", null) != null
	finding := signoz_finding(group, "sampling_relevant_on_non_span", sprintf(
		"'%s' in '%s' sets 'sampling_relevant', which only applies to spans.",
		[object.get(attr, "ref", object.get(attr, "id", "?")), group.id],
	))
}
