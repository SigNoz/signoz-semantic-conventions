package after_resolution

import rego.v1

# SigNoz attributes and signals live under `signoz.`. A key upstream already
# defines is referenced, never redefined, and upstream's namespaces (`http.`,
# `db.`, ...) are left to upstream.

signoz_vendor_namespaces := {"signoz"}

signoz_upstream_attribute_keys := {attr.key |
	some dependency in input.dependencies
	some attr in dependency.registry.attributes
}

signoz_upstream_namespaces := {split(key, ".")[0] | some key in signoz_upstream_attribute_keys} | {split(metric.name, ".")[0] |
	some dependency in input.dependencies
	some metric in dependency.registry.metrics
}

signoz_root(name) := split(name, ".")[0]

# Attributes

deny contains finding if {
	some attr in input.registry.attributes
	attr.key in signoz_upstream_attribute_keys
	not signoz_excepted(attr, "shadows_upstream_attribute")
	finding := signoz_attribute_finding(attr, "shadows_upstream_attribute", sprintf(
		"Attribute '%s' is already defined upstream. Reference it with '- ref: %s' instead of redefining it.",
		[attr.key, attr.key],
	))
}

deny contains finding if {
	some attr in input.registry.attributes
	not attr.key in signoz_upstream_attribute_keys
	root := signoz_root(attr.key)
	root in signoz_upstream_namespaces
	not root in signoz_vendor_namespaces
	not signoz_excepted(attr, "namespace_collides_with_upstream")
	finding := signoz_attribute_finding(attr, "namespace_collides_with_upstream", sprintf(
		"Attribute '%s' is under the upstream '%s.' namespace. SigNoz attributes belong under 'signoz.'.",
		[attr.key, root],
	))
}

deny contains finding if {
	some attr in input.registry.attributes
	root := signoz_root(attr.key)
	not root in signoz_upstream_namespaces
	not root in signoz_vendor_namespaces
	not signoz_excepted(attr, "namespace_not_vendor")
	finding := signoz_attribute_finding(attr, "namespace_not_vendor", sprintf(
		"Attribute '%s' is outside the 'signoz.' namespace.",
		[attr.key],
	))
}

# Signals

signoz_signals contains {"type": "metric", "name": metric.name, "signal": metric} if {
	some metric in input.registry.metrics
}

signoz_signals contains {"type": "event", "name": event.name, "signal": event} if {
	some event in input.registry.events
}

signoz_signals contains {"type": "entity", "name": entity.type, "signal": entity} if {
	some entity in input.registry.entities
}

deny contains finding if {
	some s in signoz_signals
	root := signoz_root(s.name)
	root in signoz_upstream_namespaces
	not root in signoz_vendor_namespaces
	not signoz_excepted(s.signal, "namespace_collides_with_upstream")
	finding := signoz_signal_finding(s, "namespace_collides_with_upstream", sprintf(
		"%s '%s' is under the upstream '%s.' namespace. SigNoz signals belong under 'signoz.'.",
		[s.type, s.name, root],
	))
}

deny contains finding if {
	some s in signoz_signals
	root := signoz_root(s.name)
	not root in signoz_upstream_namespaces
	not root in signoz_vendor_namespaces
	not signoz_excepted(s.signal, "namespace_not_vendor")
	finding := signoz_signal_finding(s, "namespace_not_vendor", sprintf(
		"%s '%s' is outside the 'signoz.' namespace.",
		[s.type, s.name],
	))
}
