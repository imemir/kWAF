# kwaf

Lightweight Kubernetes-native Web Application Firewall built in Go. Define security rules with CRDs and protect your services from common attacks like XSS, SQLi, and more.

## Project layout

This repository contains the basic skeleton for the kWAF project. The main components are:

- `cmd/kwaf`: entry point for the CLI application
- `internal/waf`: core WAF engine implementation
- `api/`: custom resource definitions for security rules
- `controllers/`: reconciliation logic for Kubernetes resources

This structure is intended as a starting point and will be expanded with full functionality in future commits.
