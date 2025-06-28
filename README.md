# kwaf

Lightweight Kubernetes-native Web Application Firewall built in Go. Define security rules with CRDs and protect your services from common attacks like XSS, SQLi, and more.

## Project layout

This repository contains the basic skeleton for the kWAF project. The main components are:

- `cmd/kwaf`: entry point for the CLI application
- `internal/waf`: core WAF engine implementation
- `api/`: custom resource definitions for security rules
- `controllers/`: reconciliation logic for Kubernetes resources

The repository now includes a minimal rule engine that supports regex and substring based rules as well as "allow" and "block" actions. You can test it with the CLI:

```bash
# example usage
kwaf "DROP TABLE users"
```

This structure is intended as a starting point and will be expanded with full functionality in future commits.

## Project goals

- Deliver a lightweight Web Application Firewall that runs natively on Kubernetes.
- Provide declarative security policies via Custom Resource Definitions.
- Offer a simple command line tool for local rule testing and debugging.

## Upcoming tasks

- Extend the rule engine with additional pattern types and actions.
- Expose rule and rule set CRDs through the controller manager.
- Ship example manifests and Helm charts for quick deployment.
- Establish continuous integration with unit and integration tests.

