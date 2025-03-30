# Kubernetes Operator Demo

This is a demonstration project showing how to build a Kubernetes operator in Go. The operator manages a custom resource called `DemoApp`.

## Project Structure

```
.
├── api/            # API definitions and types
├── controllers/    # Controller implementations
├── main.go         # Entry point
└── go.mod          # Go module definition
```

## Prerequisites

- Go 1.21 or later
- Kubernetes cluster (local or remote)
- kubectl configured to work with your cluster

## Getting Started

1. Install the CRD:
```bash
kubectl apply -f config/crd/bases/demoapp.example.com_demoapps.yaml
```

2. Run the operator:
```bash
go run main.go
```

## Custom Resource Example

```yaml
apiVersion: demoapp.example.com/v1alpha1
kind: DemoApp
metadata:
  name: example-demoapp
spec:
  message: "Hello from DemoApp!"
```

## Learning Objectives

This project demonstrates:
- Creating custom resources in Kubernetes
- Building operators using controller-runtime
- Implementing reconciliation loops
- Managing custom resource states 