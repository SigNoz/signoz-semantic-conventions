# Code generated from semantic convention specification. DO NOT EDIT.

from typing import Final

from opentelemetry.metrics import Counter, Histogram, Meter

SIGNOZ_CALLS_TOTAL: Final = "signoz_calls_total"
"""
The number of spans observed.
Instrument: counter
Unit: 
"""


def create_signoz_calls_total(meter: Meter) -> Counter:
    """The number of spans observed"""
    return meter.create_counter(
        name=SIGNOZ_CALLS_TOTAL,
        description="The number of spans observed.",
        unit="",
    )


SIGNOZ_DB_LATENCY_COUNT: Final = "signoz_db_latency_count"
"""
The number of database calls observed.
Instrument: counter
Unit: 1
Note: A database call is a span that carries `db.system` and is not of kind
server.
"""


def create_signoz_db_latency_count(meter: Meter) -> Counter:
    """The number of database calls observed"""
    return meter.create_counter(
        name=SIGNOZ_DB_LATENCY_COUNT,
        description="The number of database calls observed.",
        unit="1",
    )


SIGNOZ_DB_LATENCY_SUM: Final = "signoz_db_latency_sum"
"""
The total duration of database calls observed.
Instrument: counter
Unit: 1
Note: Values are in milliseconds. A database call is a span that carries
`db.system` and is not of kind server.
"""


def create_signoz_db_latency_sum(meter: Meter) -> Counter:
    """The total duration of database calls observed"""
    return meter.create_counter(
        name=SIGNOZ_DB_LATENCY_SUM,
        description="The total duration of database calls observed.",
        unit="1",
    )


SIGNOZ_EXTERNAL_CALL_LATENCY_COUNT: Final = "signoz_external_call_latency_count"
"""
The number of external calls observed.
Instrument: counter
Unit: 1
Note: An external call is a client span with a known remote endpoint.
"""


def create_signoz_external_call_latency_count(meter: Meter) -> Counter:
    """The number of external calls observed"""
    return meter.create_counter(
        name=SIGNOZ_EXTERNAL_CALL_LATENCY_COUNT,
        description="The number of external calls observed.",
        unit="1",
    )


SIGNOZ_EXTERNAL_CALL_LATENCY_SUM: Final = "signoz_external_call_latency_sum"
"""
The total duration of external calls observed.
Instrument: counter
Unit: 1
Note: Values are in milliseconds. An external call is a client span with a
known remote endpoint.
"""


def create_signoz_external_call_latency_sum(meter: Meter) -> Counter:
    """The total duration of external calls observed"""
    return meter.create_counter(
        name=SIGNOZ_EXTERNAL_CALL_LATENCY_SUM,
        description="The total duration of external calls observed.",
        unit="1",
    )


SIGNOZ_LATENCY: Final = "signoz_latency"
"""
The duration of spans observed.
Instrument: histogram
Unit: ms
Note: Recorded as an explicit bucket histogram and, when enabled, also as an
exponential histogram under the same name. Values are in milliseconds.
"""


def create_signoz_latency(meter: Meter) -> Histogram:
    """The duration of spans observed"""
    return meter.create_histogram(
        name=SIGNOZ_LATENCY,
        description="The duration of spans observed.",
        unit="ms",
    )
