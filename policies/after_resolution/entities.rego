package after_resolution

import rego.v1

# Every entity has at least one identifying attribute. Upstream's stability
# package checks this for stable entities only.

deny contains finding if {
	some entity in input.registry.entities
	count(object.get(entity, "identity", [])) == 0
	finding := signoz_signal_finding({"type": "entity", "name": entity.type}, "entity_no_identifying", sprintf(
		"Entity '%s' has no identifying attribute.",
		[entity.type],
	))
}
