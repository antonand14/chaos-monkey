#!/bin/bash

echo "🐵 Starting Chaos Monkey Development Environment..."
echo ""

# Get the script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Start backend
echo "📡 Starting backend on http://localhost:8000"
cd "$SCRIPT_DIR/backend" && python main.py &
BACKEND_PID=$!

# Wait for backend to start
sleep 2

# Start frontend
echo "🎨 Starting frontend on http://localhost:3000"
cd "$SCRIPT_DIR/frontend" && npm run dev &
FRONTEND_PID=$!

echo ""
echo "✅ Services started!"
echo "   Backend:  http://localhost:8000"
echo "   Frontend: http://localhost:3000"
echo ""
echo "Press Ctrl+C to stop all services"

# Trap Ctrl+C and kill both processes
trap "echo ''; echo '🛑 Stopping services...'; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit" INT

# Wait for both processes
wait
