package after_resolution

import rego.v1

# Attribute definitions, and attributes as signals reference them.

signoz_string_types := {"string", "string[]", "template[string]", "template[string[]]"}

deny contains finding if {
	some attr in input.registry.attributes
	attr.type in signoz_string_types
	not attr.examples
	not signoz_excepted(attr, "missing_examples")
	finding := signoz_attribute_finding(attr, "missing_examples", sprintf(
		"Attribute '%s' is a '%s' and has no 'examples'.",
		[attr.key, attr.type],
	))
}

deny contains finding if {
	some attr in input.registry.attributes
	some member in attr.type.members
	signoz_stability_rank[member.stability] > signoz_stability_rank[attr.stability]
	finding := signoz_attribute_finding(attr, "enum_member_stability", sprintf(
		"Value '%s' of attribute '%s' is '%s', more stable than the attribute ('%s').",
		[member.id, attr.key, member.stability, attr.stability],
	))
}

# Every attribute a signal, entity or public attribute group lists.
signoz_signal_attributes contains {"owner": sprintf("%s '%s'", [s.type, s.name]), "type": s.type, "attr": attr} if {
	some s in signoz_signals
	some attr in array.concat(
		object.get(s.signal, "attributes", []),
		array.concat(object.get(s.signal, "identity", []), object.get(s.signal, "description", [])),
	)
}

signoz_signal_attributes contains {"owner": sprintf("span '%s'", [span.type]), "type": "span", "attr": attr} if {
	some span in input.registry.spans
	some attr in object.get(span, "attributes", [])
}

signoz_signal_attributes contains {"owner": sprintf("attribute group '%s'", [group.id]), "type": "attribute_group", "attr": attr} if {
	some group in input.registry.attribute_groups
	some attr in object.get(group, "attributes", [])
}

deny contains finding if {
	some entry in signoz_signal_attributes
	is_object(entry.attr.requirement_level)
	some level, condition in entry.attr.requirement_level
	trim_space(condition) == ""
	finding := {
		"id": "requirement_level_condition",
		"message": sprintf("'%s' on %s is '%s' but doesn't say when.", [entry.attr.key, entry.owner, level]),
		"level": "violation",
		"context": {"attribute": entry.attr.key},
	}
}

deny contains finding if {
	some attr in input.registry.attributes
	some member in attr.type.members
	object.get(member, "brief", "") == ""
	not signoz_excepted(attr, "enum_member_brief")
	finding := signoz_attribute_finding(attr, "enum_member_brief", sprintf(
		"Value '%s' of attribute '%s' has no brief.",
		[member.id, attr.key],
	))
}

signoz_sensitive_words := {"email", "phone", "password", "token", "secret", "ssn", "credit"}

# A key segment (split at `.` and `_`) or a word in the brief or note that suggests personal or secret data.
signoz_looks_sensitive(attr) if {
	some part in regex.split(`[._]`, lower(attr.key))
	part in signoz_sensitive_words
}

signoz_looks_sensitive(attr) if contains(lower(attr.key), "full_name")

signoz_looks_sensitive(attr) if {
	text := concat(" ", [object.get(attr, "brief", ""), object.get(attr, "note", "")])
	regex.match(`(?i)\b(email|phone|password|token|secret|ssn|full_name|credit)\b`, text)
}

# Keys or text that suggest personal or secret data need a warning note.
deny contains finding if {
	some attr in input.registry.attributes
	note := object.get(attr, "note", "")
	signoz_looks_sensitive(attr)
	not contains(note, "[!WARNING]")
	not signoz_excepted(attr, "pii_warning")
	finding := signoz_attribute_finding(attr, "pii_warning", sprintf(
		"Attribute '%s' may carry sensitive data. Add a '> [!WARNING]' note saying so, or an exception if it doesn't.",
		[attr.key],
	))
}

deny contains finding if {
	some event in input.registry.events
	some attr in object.get(event, "attributes", [])
	attr.key == "event.name"
	finding := signoz_signal_finding({"type": "event", "name": event.name}, "event_name_attribute", sprintf(
		"Event '%s' references 'event.name'. The event name is the record's own field; drop the attribute.",
		[event.name],
	))
}

deny contains finding if {
	some attr in input.registry.attributes
	is_object(attr.type)
	count(object.get(attr.type, "members", [])) == 0
	finding := signoz_attribute_finding(attr, "enum_members_empty", sprintf(
		"Enum attribute '%s' has no members.",
		[attr.key],
	))
}

# An array attribute's examples are lists of values, one list per example.
signoz_example_holders contains {"key": attr.key, "type": attr.type, "examples": attr.examples} if {
	some attr in input.registry.attributes
	attr.examples
}

signoz_example_holders contains {"key": entry.attr.key, "type": entry.attr.type, "examples": entry.attr.examples} if {
	some entry in signoz_signal_attributes
	entry.attr.examples
}

deny contains finding if {
	some holder in signoz_example_holders
	is_string(holder.type)
	regex.match(`\[\]\]?$`, holder.type) # string[] and template[string[]] alike
	some example in signoz_as_array(holder.examples)
	not is_array(example)
	finding := {
		"id": "examples_array_shape",
		"message": sprintf("Attribute '%s' is a '%s'; each example must itself be a list (`- - \"a\"`).", [holder.key, holder.type]),
		"level": "violation",
		"context": {"attribute": holder.key},
	}
}

signoz_as_array(value) := value if is_array(value)

signoz_as_array(value) := [value] if not is_array(value)
