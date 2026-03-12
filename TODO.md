# Chaos Monkey Project TODO

## Phase 1: Get it Working ✨

### Basic Infrastructure
- [ ] Test backend server (start and verify endpoints)
- [ ] Install and run frontend (npm install && npm run dev)
- [ ] Verify backend-frontend communication
- [ ] Create first experiment through UI

### Kubernetes Setup
- [ ] Set up local K8s cluster (minikube/kind/k3s)
- [ ] Deploy a test application (nginx or similar)
- [ ] Test agent can connect to cluster
- [ ] Verify agent can list pods

### First Chaos Experiment
- [ ] Add HTTP endpoint to agent to accept commands
- [ ] Wire backend to trigger agent's PodAssassin
- [ ] Kill first pod successfully
- [ ] Watch pod recover automatically
- [ ] Display experiment status in UI

## Phase 2: Make it Useful 🔧

### Real-time Updates
- [ ] Add WebSocket support to backend
- [ ] Stream chaos events to frontend
- [ ] Show live pod status during experiments
- [ ] Add experiment progress indicator

### Recovery Metrics
- [ ] Implement MeasureRecovery function
- [ ] Calculate recovery time
- [ ] Calculate basic resilience score
- [ ] Display metrics in UI
- [ ] Store experiment results

### Safety Features
- [ ] Validate blast radius limits
- [ ] Implement emergency stop button
- [ ] Add dry-run mode
- [ ] Add protected namespaces whitelist
- [ ] Prevent accidental production chaos

## Phase 3: Make it Fun 🎮

### Gamification
- [ ] Design scoring algorithm
- [ ] Build leaderboard UI
- [ ] Implement achievements system
  - [ ] "Unkillable" achievement
  - [ ] "Circuit Breaker Champion"
  - [ ] "Chaos Veteran"
  - [ ] "Phoenix" achievement
- [ ] Track experiment history
- [ ] Add user profiles (optional)

### Additional Chaos Types
- [ ] Network Gremlin (latency injection)
- [ ] Network Gremlin (packet loss)
- [ ] Resource Hog (CPU stress)
- [ ] Resource Hog (memory stress)
- [ ] Disk Filler
- [ ] Time Traveler (clock skew)
- [ ] DNS Chaos

### Visualization
- [ ] Service dependency graph
- [ ] Real-time blast radius visualization
- [ ] Recovery time charts
- [ ] Chaos timeline view
- [ ] Cluster health dashboard

## Phase 4: Production Ready 🚀

### Persistence
- [ ] Set up PostgreSQL database
- [ ] Store experiment history
- [ ] Store scores and achievements
- [ ] Track resilience trends over time
- [ ] Export reports

### Multi-cluster Support
- [ ] Connect multiple clusters
- [ ] Cluster selection in UI
- [ ] Compare resilience across environments
- [ ] Cross-cluster chaos scenarios

### Scheduling & Automation
- [ ] Cron-like experiment scheduling
- [ ] Recurring chaos experiments
- [ ] Chaos days (sustained periods)
- [ ] CI/CD pipeline integration
- [ ] Slack/Discord notifications

### Advanced Features
- [ ] Authentication/authorization
- [ ] Role-based access control
- [ ] Audit logging
- [ ] Chaos budgets (rate limiting)
- [ ] Custom chaos experiment templates
- [ ] API documentation
- [ ] Prometheus metrics export

## Infrastructure & DevOps

### Deployment
- [ ] Create Dockerfile for backend
- [ ] Create Dockerfile for agent
- [ ] Create Dockerfile for frontend
- [ ] Helm chart for deployment
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Automated testing

### Documentation
- [ ] API documentation
- [ ] User guide
- [ ] Architecture diagrams
- [ ] Contributing guidelines
- [ ] Security best practices

## Ideas for Later 💡

- [ ] Community chaos experiment marketplace
- [ ] ML-based resilience predictions
- [ ] Automated chaos recommendations
- [ ] Integration with observability tools (Grafana, Datadog)
- [ ] Mobile app for monitoring
- [ ] Chaos game mode (competitive)
- [ ] Replay past experiments
- [ ] A/B testing for resilience strategies

---

## Current Status

**Last Updated:** 2026-03-12

**Completed:**
- ✅ Project scaffolding
- ✅ Backend API skeleton
- ✅ Agent skeleton with PodAssassin
- ✅ Frontend UI skeleton
- ✅ Kubernetes manifests
- ✅ Documentation
- ✅ Git repository initialized
- ✅ Pushed to GitHub

**Next Up:** Test basic infrastructure and get UI running
