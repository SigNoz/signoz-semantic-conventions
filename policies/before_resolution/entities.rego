package before_resolution

import rego.v1

# Every attribute of an entity says whether it identifies the entity or
# describes it. Weaver accepts an attribute without one.

deny contains finding if {
	some group in input.groups
	group.type == "entity"
	some attr in object.get(group, "attributes", [])
	attr.ref
	not attr.role
	not signoz_excepted(group, "entity_role_implicit")
	finding := signoz_finding(group, "entity_role_implicit", sprintf(
		"Attribute '%s' of entity '%s' has no 'role'. Set 'identifying' or 'descriptive'.",
		[attr.ref, object.get(group, "name", group.id)],
	))
}
