# Kubernetes Manifests

Deployment manifests for running Chaos Monkey in Kubernetes.

## Deploy

```bash
# Create namespace
kubectl apply -f namespace.yaml

# Deploy chaos agent
kubectl apply -f agent-deployment.yaml

# Verify deployment
kubectl get pods -n chaos-monkey
```

## RBAC Permissions

The chaos agent requires permissions to:
- List and get pods
- Delete pods (for chaos experiments)
- Read pod logs
- List deployments and replicasets

## Security Considerations

- The agent has cluster-wide pod deletion permissions
- Consider using namespace-scoped roles for production
- Implement additional safety checks in the agent code
- Use admission webhooks to protect critical resources

## TODO

- [ ] Backend deployment manifest
- [ ] Frontend deployment manifest
- [ ] Ingress configuration
- [ ] ConfigMap for configuration
- [ ] Secret management
