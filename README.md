## Iris
A WIP drop in replacement for Prometheus Alertmanager

## Components
Distributor: Receives incoming alerts and distributes them across the dispatcher ring using consistent hashing
Dispatcher: Core component responsible for handling alerts, implementing alert grouping, and routing
NotifyFlow: Manages the notification workflow including silences, inhibition rules, and grouping
Notifier: Handles sending notifications to various integrations (email, Slack, etc.)
AlertStore: Persists alerts and their states
Analyzer: Provides analytics and querying capabilities for stored alerts

## Storage supported
| Name      | Status   |
|-----------|:--------:|
| Redis     |Supported |
| Dragonfly |Supported |
| Milvus    |Supported |