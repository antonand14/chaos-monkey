# Chaos Monkey as a Service

Test your Kubernetes cluster resilience through controlled chaos experiments with gamification.

## Quick Start

### Prerequisites
- Python 3.11+ (3.14 works!)
- Go 1.21+
- Node.js 18+
- Kubernetes cluster (local or remote)
- kubectl configured

### Setup

1. **Backend Setup**
```bash
cd backend
# No dependencies needed!
python main.py
```

2. **Agent Setup**
```bash
cd agent
go mod download
go build -o chaos-agent
```

3. **Frontend Setup**
```bash
cd frontend
npm install
npm run dev
```

Or use the convenience scripts:
```bash
./start-backend.sh  # Starts backend on port 8000
```

## Architecture

- **Backend**: FastAPI REST API for experiment management
- **Agent**: Go-based chaos agent deployed in K8s clusters
- **Frontend**: React dashboard for experiment control and visualization
- **Database**: PostgreSQL for experiment history and scores

## Development Roadmap

- [x] Project scaffolding
- [ ] Backend API endpoints
- [ ] Kubernetes client integration
- [ ] Pod deletion chaos experiment
- [ ] Real-time WebSocket updates
- [ ] Basic scoring system
- [ ] Frontend dashboard
- [ ] Leaderboard

## Contributing

This is a personal tinkering project. Feel free to fork and experiment!

## License

MIT
