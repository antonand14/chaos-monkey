# Backend API

Simple HTTP server-based REST API for managing chaos experiments (no external dependencies needed).

## Setup

```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

No dependencies required for basic functionality!

## Run

```bash
python main.py
```

API will be available at http://localhost:8000

## Endpoints

- `GET /` - API info
- `POST /experiments` - Create chaos experiment
- `GET /experiments` - List all experiments
- `GET /experiments/{id}` - Get experiment details
- `POST /experiments/{id}/run` - Execute experiment
- `POST /experiments/{id}/stop` - Emergency stop
- `GET /leaderboard` - Get resilience scores

## TODO

- [ ] Kubernetes client integration (will need `pip install kubernetes`)
- [ ] Database persistence (PostgreSQL)
- [ ] Authentication/authorization
- [ ] Experiment scheduling
- [ ] Metrics collection
- [ ] WebSocket support for real-time updates
