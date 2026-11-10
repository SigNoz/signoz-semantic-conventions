# Code generated from semantic convention specification. DO NOT EDIT.

from enum import Enum
from typing import Final

ADDRESS: Final = "address"
"""
The remote endpoint of an external call.
Note: `<rpc.service>/<rpc.method>` for RPC calls, otherwise the remote host
as `host` or `host:port`, or the span's `peer.service`.
"""

OPERATION: Final = "operation"
"""
The name of the aggregated spans.
Note: Set to `overflow_operation` when the service has more distinct span
names than are tracked.
"""

RESOURCE_DEPLOYMENT_ENVIRONMENT: Final = "resource_deployment.environment"
"""
The value of `deployment.environment` on the span's Resource.
"""

RESOURCE_SERVICE_NAMESPACE: Final = "resource_service.namespace"
"""
The value of `service.namespace` on the span's Resource.
"""

RESOURCE_SIGNOZ_COLLECTOR_ID: Final = "resource_signoz.collector.id"
"""
A copy of `signoz.collector.id`.
"""

SIGNOZ_COLLECTOR_ID: Final = "signoz.collector.id"
"""
The identifier of the collector instance that produced the metric.
Note: The collector's `service.instance.id`, or a UUID assigned when the
collector starts if it has none.
"""

SPAN_KIND: Final = "span.kind"
"""
The kind of the aggregated spans.
"""

STATUS_CODE: Final = "status.code"
"""
The status code of the aggregated spans.
"""


class SpanKindValues(Enum):
    UNSPECIFIED = "SPAN_KIND_UNSPECIFIED"
    """The span kind is unspecified."""
    INTERNAL = "SPAN_KIND_INTERNAL"
    """An internal operation."""
    SERVER = "SPAN_KIND_SERVER"
    """The server side of a remote call."""
    CLIENT = "SPAN_KIND_CLIENT"
    """The client side of a remote call."""
    PRODUCER = "SPAN_KIND_PRODUCER"
    """The producer of an asynchronous message."""
    CONSUMER = "SPAN_KIND_CONSUMER"
    """The consumer of an asynchronous message."""


class StatusCodeValues(Enum):
    UNSET = "STATUS_CODE_UNSET"
    """The span status is unset."""
    OK = "STATUS_CODE_OK"
    """The span completed successfully."""
    ERROR = "STATUS_CODE_ERROR"
    """The span completed with an error."""
