# Chaos Monkey Project Structure

```
chaos-monkey/
├── backend/              # FastAPI REST API
│   ├── main.py          # API endpoints
│   ├── requirements.txt # Python dependencies
│   └── README.md
│
├── agent/               # Go chaos agent
│   ├── main.go         # Chaos execution logic
│   ├── go.mod          # Go dependencies
│   └── README.md
│
├── frontend/            # React dashboard
│   ├── src/
│   │   ├── App.jsx     # Main component
│   │   ├── App.css     # Styles
│   │   ├── main.jsx    # Entry point
│   │   └── index.css   # Base styles
│   ├── package.json
│   ├── vite.config.js
│   └── README.md
│
├── k8s/                 # Kubernetes manifests
│   ├── namespace.yaml
│   ├── agent-deployment.yaml
│   └── README.md
│
├── docs/                # Additional documentation
│
├── PROJECT.md           # Detailed project description
└── README.md            # Quick start guide
```

## Component Overview

### Backend (Python/FastAPI)
- REST API for experiment management
- WebSocket support for real-time updates
- Experiment scheduling and orchestration
- Leaderboard and scoring system

### Agent (Go)
- Runs inside Kubernetes clusters
- Executes chaos experiments
- Measures recovery metrics
- Reports results back to backend

### Frontend (React)
- Web dashboard for experiment control
- Real-time monitoring
- Leaderboard visualization
- Cluster management

### K8s Manifests
- Deployment configurations
- RBAC permissions
- Service definitions
- Ingress rules

## Development Workflow

1. Start backend: `cd backend && python main.py`
2. Start frontend: `cd frontend && npm run dev`
3. Build agent: `cd agent && go build`
4. Deploy to K8s: `kubectl apply -f k8s/`

## Next Steps

1. Implement Kubernetes client integration in backend
2. Add WebSocket real-time updates
3. Build out chaos experiment types
4. Implement scoring algorithm
5. Add authentication/authorization
6. Create Docker images for deployment
