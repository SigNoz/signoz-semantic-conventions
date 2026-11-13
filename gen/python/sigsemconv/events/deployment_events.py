# Code generated from semantic convention specification. DO NOT EDIT.

from typing import Final

SIGNOZ_DEPLOYMENT: Final = "signoz.deployment"
"""
A deployment of a service.
Note: `service.name`, `service.version` and `deployment.environment.name` are set
on the OpenTelemetry Resource; every other attribute is set on the log
record. Severity is `INFO` (9) for `succeeded` and `ERROR` (17) for `failed`.
"""
