# Code generated from semantic convention specification. DO NOT EDIT.

from typing import Final

CONNECTOR_ID: Final = "connector_id"
"""
The identifier of the meter connector instance that produced the metric.
Note: A UUID assigned when the connector instance is created.
"""

SIGNOZ_IO_WORKLOAD: Final = "signoz.io/workload"
"""
The SigNoz workload a Kubernetes node is dedicated to.
Note: The value of the node's `signoz.io/workload` label.

Known values: `clickhouse`, `redpanda`, `consumers`, `zookeeper`,
`gateway`, `signoz`.
"""
