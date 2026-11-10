# Code generated from semantic convention specification. DO NOT EDIT.

from collections.abc import Callable, Generator, Iterable, Sequence
from typing import Final

from opentelemetry.metrics import (
    CallbackOptions,
    Counter,
    Meter,
    ObservableGauge,
    Observation,
)

# pylint: disable=invalid-name
CallbackT = (
    Callable[[CallbackOptions], Iterable[Observation]]
    | Generator[Iterable[Observation], CallbackOptions, None]
)

SIGNOZ_METER_CPU_NODE_ALLOCATABLE_USAGE: Final = (
    "signoz.meter.cpu.node.allocatable.usage"
)
"""
The amount of CPU allocatable on a Kubernetes node that runs SigNoz.
Instrument: gauge
Unit: {cpu}
Note: The value is in CPU cores and is recorded once per node every hour.
"""


def create_signoz_meter_cpu_node_allocatable_usage(
    meter: Meter, callbacks: Sequence[CallbackT] | None
) -> ObservableGauge:
    """The amount of CPU allocatable on a Kubernetes node that runs SigNoz"""
    return meter.create_observable_gauge(
        name=SIGNOZ_METER_CPU_NODE_ALLOCATABLE_USAGE,
        callbacks=callbacks,
        description="The amount of CPU allocatable on a Kubernetes node that runs SigNoz.",
        unit="{cpu}",
    )


SIGNOZ_METER_LOG_COUNT: Final = "signoz.meter.log.count"
"""
The number of log records observed.
Instrument: counter
Unit: 1
Note: Each log record counts as one.
"""


def create_signoz_meter_log_count(meter: Meter) -> Counter:
    """The number of log records observed"""
    return meter.create_counter(
        name=SIGNOZ_METER_LOG_COUNT,
        description="The number of log records observed.",
        unit="1",
    )


SIGNOZ_METER_LOG_SIZE: Final = "signoz.meter.log.size"
"""
The size of log records observed.
Instrument: counter
Unit: By
Note: The size of a log record is the size of its resource attributes, its
attributes and its body. `signoz.workspace.*` attributes are not
counted.
"""


def create_signoz_meter_log_size(meter: Meter) -> Counter:
    """The size of log records observed"""
    return meter.create_counter(
        name=SIGNOZ_METER_LOG_SIZE,
        description="The size of log records observed.",
        unit="By",
    )


SIGNOZ_METER_METRIC_DATAPOINT_COUNT: Final = "signoz.meter.metric.datapoint.count"
"""
The number of metric data points observed.
Instrument: counter
Unit: 1
Note: A gauge or sum data point counts as one. A histogram data point counts
as its number of buckets plus one for its count and one each for its
sum, min and max when present.
A summary data point counts as its number of quantiles plus two.
Exponential histogram data points, and metrics whose name starts with
`signoz` or `otelcol`, are not counted.
"""


def create_signoz_meter_metric_datapoint_count(meter: Meter) -> Counter:
    """The number of metric data points observed"""
    return meter.create_counter(
        name=SIGNOZ_METER_METRIC_DATAPOINT_COUNT,
        description="The number of metric data points observed.",
        unit="1",
    )


SIGNOZ_METER_METRIC_DATAPOINT_SIZE: Final = "signoz.meter.metric.datapoint.size"
"""
The size of metric data points observed.
Instrument: counter
Unit: By
Note: The value is always `0`.
"""


def create_signoz_meter_metric_datapoint_size(meter: Meter) -> Counter:
    """The size of metric data points observed"""
    return meter.create_counter(
        name=SIGNOZ_METER_METRIC_DATAPOINT_SIZE,
        description="The size of metric data points observed.",
        unit="By",
    )


SIGNOZ_METER_SPAN_COUNT: Final = "signoz.meter.span.count"
"""
The number of spans observed.
Instrument: counter
Unit: 1
Note: Each span counts as one.
"""


def create_signoz_meter_span_count(meter: Meter) -> Counter:
    """The number of spans observed"""
    return meter.create_counter(
        name=SIGNOZ_METER_SPAN_COUNT,
        description="The number of spans observed.",
        unit="1",
    )


SIGNOZ_METER_SPAN_SIZE: Final = "signoz.meter.span.size"
"""
The size of spans observed.
Instrument: counter
Unit: By
Note: The size of a span is the size of its fields, its resource attributes,
its attributes, its events and its links. `signoz.workspace.*` and
`signoz.gen_ai.*` attributes are not counted.
"""


def create_signoz_meter_span_size(meter: Meter) -> Counter:
    """The size of spans observed"""
    return meter.create_counter(
        name=SIGNOZ_METER_SPAN_SIZE,
        description="The size of spans observed.",
        unit="By",
    )
