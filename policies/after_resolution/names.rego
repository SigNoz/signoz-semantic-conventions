# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0
#
# Copied from open-telemetry/opentelemetry-weaver-packages
# policies/check/naming_conventions/names.rego at
# 92eba8ea6d06c6b60a3386917858443816977810, and modified:
# - each rule skips a name that lists the rule id in
#   annotations.signoz.policy_exceptions (upstream has no attribute exceptions);
# - naming_convention_attr_name checks the full format of attribute names.
# Drop this copy for upstream's once it supports attribute exceptions.

package after_resolution

import rego.v1

# checks attribute name has a namespace
deny contains finding if {
	some attr in input.registry.attributes
	name := attr.key

	# some deprecated attributes have no namespace and need to be ignored
	not attr.deprecated
	not regex.match(has_namespace_regex, name)
	not signoz_excepted(attr, "naming_convention_attr_has_namespace")

	finding := {
		"id": "naming_convention_attr_has_namespace",
		"context": {"attribute": attr.key},
		"message": sprintf("Attribute name '%s' should have a namespace. Attribute name %s", [name, invalid_name_helper]),
		"level": "violation",
	}
}

# checks attribute name format
deny contains finding if {
	some attr in input.registry.attributes
	name := attr.key
	not attr.deprecated
	not regex.match(name_regex, name)
	not signoz_excepted(attr, "naming_convention_attr_name")

	finding := {
		"id": "naming_convention_attr_name",
		"context": {"attribute": attr.key},
		"message": sprintf("Attribute name '%s' is invalid. Attribute name %s", [name, invalid_name_helper]),
		"level": "violation",
	}
}

# checks attribute enum member id format
deny contains finding if {
	some attr in input.registry.attributes
	some member in attr.type.members
	name := member.id
	not regex.match(name_regex, name)
	not signoz_excepted(attr, "naming_convention_enum_id")

	finding := {
		"id": "naming_convention_enum_id",
		"context": {
			"attribute": attr.key,
			"member_id": name,
		},
		"message": sprintf("Member id '%s' on attribute '%s' is invalid. Member id %s", [name, attr.key, invalid_name_helper]),
		"level": "violation",
	}
}

# checks entity type format
deny contains finding if {
	some entity in input.registry.entities
	name := entity.type

	name != null
	not regex.match(name_regex, name)
	not signoz_excepted(entity, "naming_convention_entity_type")

	finding := {
		"id": "naming_convention_entity_type",
		"context": {},
		"message": sprintf("Entity type '%s' is invalid. Entity type %s", [name, invalid_name_helper]),
		"level": "violation",
		"signal_type": "entity",
		"signal_name": name,
	}
}

# checks span type format
deny contains finding if {
	some span in input.registry.spans
	name := span.type

	name != null
	not regex.match(name_regex, name)
	not signoz_excepted(span, "naming_convention_span_type")

	finding := {
		"id": "naming_convention_span_type",
		"context": {},
		"message": sprintf("Span type '%s' is invalid. Span type %s", [name, invalid_name_helper]),
		"level": "violation",
		"signal_type": "span",
		"signal_name": name,
	}
}

# checks metric name format
deny contains finding if {
	some metric in input.registry.metrics
	name := metric.name

	name != null
	not regex.match(name_regex, name)
	not signoz_excepted(metric, "naming_convention_metric_name")

	finding := {
		"id": "naming_convention_metric_name",
		"context": {},
		"message": sprintf("Metric name '%s' is invalid. Metric name %s", [name, invalid_name_helper]),
		"level": "violation",
		"signal_type": "metric",
		"signal_name": name,
	}
}

# checks event name format
deny contains finding if {
	some event in input.registry.events
	name := event.name

	name != null
	not regex.match(name_regex, name)
	not signoz_excepted(event, "naming_convention_event_name")

	finding := {
		"id": "naming_convention_event_name",
		"context": {},
		"message": sprintf("Event name '%s' is invalid. Event name %s", [name, invalid_name_helper]),
		"level": "violation",
		"signal_type": "event",
		"signal_name": name,
	}
}

# not valid: '1foo.bar', 'foo.bar.', 'foo.bar_', 'foo..bar', 'foo._bar' ...
# valid: 'foo.bar', 'foo.1bar', 'foo.1_bar'
name_regex := "^[a-z][a-z0-9]*([._][a-z0-9]+)*$"

has_namespace_regex := "^[a-z0-9_]+\\.([a-z0-9._]+)+$"

invalid_name_helper := "must consist of lowercase alphanumeric characters separated by '_' and '.'"

is_empty_or_null(obj, property) if {
	prop := object.get(obj, property, null)
	{prop == null, prop == ""}[_]
}
