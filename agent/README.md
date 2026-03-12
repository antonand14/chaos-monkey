# Chaos Agent

Go-based agent that executes chaos experiments in Kubernetes clusters.

## Setup

```bash
go mod download
go build -o chaos-agent
```

## Run

```bash
# With kubeconfig
./chaos-agent -kubeconfig ~/.kube/config

# In-cluster (when deployed to K8s)
./chaos-agent
```

## Features

- **Pod Assassin**: Randomly delete pods based on labels
- **Recovery Measurement**: Track how long services take to recover

## TODO

- [ ] Network chaos (latency, packet loss)
- [ ] Resource chaos (CPU/memory pressure)
- [ ] API integration with backend
- [ ] Safety checks and blast radius limits
- [ ] Dry-run mode
- [ ] Metrics export
