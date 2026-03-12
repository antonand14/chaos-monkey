# Chaos Monkey as a Service

A web platform where teams can test their Kubernetes cluster resilience through controlled chaos experiments with gamification elements.

## Core Concept

Users connect their K8s clusters, configure chaos experiments through a web UI, and watch in real-time as chaos is unleashed. The system measures how well applications recover and assigns scores based on resilience metrics.

## Key Features

### 1. Chaos Experiments Library
- **Pod Assassin**: Randomly kills pods matching labels (deployment, namespace, or cluster-wide)
- **Network Gremlin**: Injects latency, packet loss, or bandwidth limits between services
- **Resource Hog**: Consumes CPU/memory on nodes to simulate resource pressure
- **Time Traveler**: Manipulates system time in containers to test time-dependent logic
- **DNS Chaos**: Returns wrong IPs or timeouts for service discovery
- **Disk Filler**: Fills up persistent volumes to test storage failure handling
- **Certificate Expirer**: Simulates expired TLS certificates
- **Zone Outage**: Takes down entire availability zones

### 2. Experiment Scheduler
- One-time chaos runs
- Recurring schedules (daily chaos at 2pm, random times during business hours)
- "Chaos Days" - sustained multi-hour chaos events
- Blast radius controls (max % of pods, specific namespaces only)

### 3. Real-time Dashboard
- Live view of chaos being injected
- Service health metrics (response times, error rates, pod restarts)
- Recovery time tracking
- Blast radius visualization showing affected services

### 4. Gamification & Leaderboards

**Resilience Score** calculated from:
- Mean Time To Recovery (MTTR)
- Blast radius containment (did failures cascade?)
- Zero-downtime deployments during chaos
- Auto-healing speed

**Achievements**:
- "Unkillable" - survived 100 pod deletions with <1s downtime
- "Circuit Breaker Champion" - prevented cascading failures
- "Chaos Veteran" - ran 50+ experiments
- "Phoenix" - recovered from total namespace wipeout in <30s

**Leaderboards**:
- Most resilient cluster (by score)
- Fastest recovery times
- Most chaos experiments survived
- Team rankings

### 5. Safety Features
- Dry-run mode to preview impact
- Emergency stop button
- Automatic rollback if critical services fail
- Whitelist/blacklist for protected resources
- Chaos budgets (max failures per hour)

## Example Scenarios

### Scenario 1: E-commerce Checkout Resilience
```
Experiment: "Black Friday Stress Test"
- Kill 50% of payment service pods
- Inject 500ms latency to database
- Fill up Redis cache storage to 95%

Expected behavior:
- Checkout still works (degraded performance OK)
- No lost transactions
- Auto-scaling kicks in
- Circuit breakers prevent cascade

Score factors:
- Did orders complete? (+100 points)
- Recovery time < 10s? (+50 points)
- No manual intervention? (+25 points)
```

### Scenario 2: Microservices Cascade Test
```
Experiment: "Domino Effect"
- Kill authentication service pods
- Watch which services fail

Good architecture:
- Services degrade gracefully
- Cached tokens keep working
- Retry logic with backoff
- Clear error messages to users

Bad architecture:
- Everything returns 500
- Infinite retry loops
- Cascading timeouts
- Database connection pool exhaustion

Resilience score shows the difference
```

### Scenario 3: Data Persistence Challenge
```
Experiment: "Database Disaster"
- Terminate database pods
- Corrupt 10% of disk blocks
- Network partition between replicas

Tests:
- Do you have proper backups?
- Can you failover to replica?
- Is data consistency maintained?
- How long until full recovery?

Leaderboard category: "Data Guardian"
```

## Tech Stack

### Backend
- Go for chaos agents (lightweight, runs in cluster)
- Python/FastAPI for web API
- Kubernetes client libraries for cluster interaction

### Frontend
- React for dashboard
- WebSocket for real-time updates
- D3.js or similar for service mesh visualization

### Infrastructure
- Kubernetes operator pattern for chaos agents
- Prometheus for metrics collection
- PostgreSQL for experiment history and scores

### Chaos Injection
- Kubernetes client-go for direct pod manipulation
- Custom CRDs for experiment definitions
- Admission webhooks for safety controls

## MVP Scope

Start with:
1. Simple web UI to connect a K8s cluster (kubeconfig upload)
2. One chaos type: random pod deletion
3. Real-time view of pods being killed
4. Basic resilience score (recovery time only)
5. Simple leaderboard (single user for now)

## Project Structure

```
chaos-monkey/
├── backend/           # FastAPI backend
├── agent/            # Go chaos agent (runs in K8s)
├── frontend/         # React dashboard
├── k8s/              # Kubernetes manifests
└── docs/             # Additional documentation
```

## Getting Started

See individual README files in each directory for setup instructions.
