from dataclasses import dataclass, asdict
from typing import List
import json
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.parse import urlparse
import threading

# Models
@dataclass
class ChaosExperiment:
    name: str
    type: str
    target_namespace: str
    target_labels: dict
    blast_radius: int
    duration_seconds: int
    dry_run: bool = False

# In-memory storage
experiments = []
results = []

class ChaosAPIHandler(BaseHTTPRequestHandler):
    def _send_json(self, data, status=200):
        self.send_response(status)
        self.send_header('Content-type', 'application/json')
        self.send_header('Access-Control-Allow-Origin', '*')
        self.end_headers()
        self.wfile.write(json.dumps(data).encode())

    def _send_cors(self):
        self.send_response(200)
        self.send_header('Access-Control-Allow-Origin', '*')
        self.send_header('Access-Control-Allow-Methods', 'GET, POST, OPTIONS')
        self.send_header('Access-Control-Allow-Headers', 'Content-Type')
        self.end_headers()

    def do_OPTIONS(self):
        self._send_cors()

    def do_GET(self):
        path = urlparse(self.path).path
        
        if path == '/':
            self._send_json({"message": "Chaos Monkey API", "version": "0.1.0"})
        elif path == '/experiments':
            self._send_json(experiments)
        elif path == '/leaderboard':
            self._send_json({"leaderboard": []})
        elif path.startswith('/experiments/'):
            exp_id = path.split('/')[-1]
            exp = next((e for e in experiments if e["id"] == exp_id), None)
            if exp:
                self._send_json(exp)
            else:
                self._send_json({"error": "Not found"}, 404)
        else:
            self._send_json({"error": "Not found"}, 404)

    def do_POST(self):
        path = urlparse(self.path).path
        content_length = int(self.headers.get('Content-Length', 0))
        body = self.rfile.read(content_length).decode() if content_length > 0 else '{}'
        
        try:
            data = json.loads(body)
        except:
            data = {}

        if path == '/experiments':
            exp_id = f"exp-{len(experiments) + 1}"
            experiments.append({"id": exp_id, **data})
            self._send_json({"experiment_id": exp_id, "status": "created"})
        elif path.endswith('/run'):
            exp_id = path.split('/')[-2]
            self._send_json({"status": "running", "experiment_id": exp_id})
        elif path.endswith('/stop'):
            exp_id = path.split('/')[-2]
            self._send_json({"status": "stopped", "experiment_id": exp_id})
        else:
            self._send_json({"error": "Not found"}, 404)

def run_server(port=8000):
    server = HTTPServer(('0.0.0.0', port), ChaosAPIHandler)
    print(f"🐵 Chaos Monkey API running on http://localhost:{port}")
    print(f"📖 API endpoints:")
    print(f"  GET  / - API info")
    print(f"  GET  /experiments - List experiments")
    print(f"  POST /experiments - Create experiment")
    print(f"  POST /experiments/{{id}}/run - Run experiment")
    print(f"  POST /experiments/{{id}}/stop - Stop experiment")
    print(f"  GET  /leaderboard - Get leaderboard")
    server.serve_forever()

if __name__ == "__main__":
    run_server()
