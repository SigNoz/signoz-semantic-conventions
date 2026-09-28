package after_resolution

import rego.v1

# Metric naming and unit rules from the OpenTelemetry metric guidelines.
# A metric can opt out of a rule with annotations.signoz.policy_exceptions.

signoz_metric_finding(metric, rule, message) := {
	"id": rule,
	"message": message,
	"level": "violation",
	"context": {},
	"signal_type": "metric",
	"signal_name": metric.name,
}

signoz_duration_units := {"ms", "us", "ns", "min", "h"}

signoz_prefixed_byte_units := {"KiBy", "MiBy", "GiBy", "kBy", "MBy"}

deny contains finding if {
	some metric in input.registry.metrics
	regex.match(`(_total|\.total)$`, metric.name)
	not signoz_excepted(metric, "metric_total_suffix")
	finding := signoz_metric_finding(metric, "metric_total_suffix", sprintf(
		"Metric '%s' ends in 'total'. Counters are sums already; drop the suffix.",
		[metric.name],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	regex.match(`^\{[a-z_]+s\}$`, metric.unit)
	not endswith(metric.unit, "ss}")
	not signoz_excepted(metric, "unit_annotation_plural")
	finding := signoz_metric_finding(metric, "unit_annotation_plural", sprintf(
		"Metric '%s' has unit '%s'. Annotation units are singular, like '{request}'.",
		[metric.name, metric.unit],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	metric.unit in signoz_duration_units
	not signoz_excepted(metric, "unit_duration_seconds")
	finding := signoz_metric_finding(metric, "unit_duration_seconds", sprintf(
		"Metric '%s' has unit '%s'. Durations are measured in 's'.",
		[metric.name, metric.unit],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	metric.unit in signoz_prefixed_byte_units
	not signoz_excepted(metric, "unit_prefixed")
	finding := signoz_metric_finding(metric, "unit_prefixed", sprintf(
		"Metric '%s' has unit '%s'. Use the unprefixed unit 'By'.",
		[metric.name, metric.unit],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	metric.instrument == "counter"
	metric.unit == "1"
	not signoz_excepted(metric, "counter_unit_one")
	finding := signoz_metric_finding(metric, "counter_unit_one", sprintf(
		"Counter '%s' has unit '1'. A count of things uses an annotation unit, like '{request}'.",
		[metric.name],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	endswith(metric.name, ".duration")
	metric.instrument != "histogram"
	not signoz_excepted(metric, "duration_instrument")
	finding := signoz_metric_finding(metric, "duration_instrument", sprintf(
		"Metric '%s' is a '%s'. '*.duration' metrics are histograms.",
		[metric.name, metric.instrument],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	endswith(metric.name, ".duration")
	metric.unit != "s"
	not signoz_excepted(metric, "duration_unit")
	finding := signoz_metric_finding(metric, "duration_unit", sprintf(
		"Metric '%s' has unit '%s'. '*.duration' metrics are measured in 's'.",
		[metric.name, metric.unit],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	endswith(metric.name, ".time")
	metric.instrument != "counter"
	not signoz_excepted(metric, "time_instrument")
	finding := signoz_metric_finding(metric, "time_instrument", sprintf(
		"Metric '%s' is a '%s'. '*.time' metrics are counters; per-operation timings are '*.duration' histograms.",
		[metric.name, metric.instrument],
	))
}

deny contains finding if {
	some metric in input.registry.metrics
	endswith(metric.name, ".utilization")
	metric.unit != "1"
	not signoz_excepted(metric, "utilization_unit")
	finding := signoz_metric_finding(metric, "utilization_unit", sprintf(
		"Metric '%s' has unit '%s'. Utilization is a ratio with unit '1'.",
		[metric.name, metric.unit],
	))
}
