#!/bin/bash

echo "🧪 Testing Chaos Monkey Components..."
echo ""

# Test backend
echo "1️⃣  Testing Backend..."
cd backend
python main.py > /tmp/chaos-backend-test.log 2>&1 &
BACKEND_PID=$!
sleep 2

# Test API endpoints
API_TEST=$(curl -s http://localhost:8000/)
if [[ $API_TEST == *"Chaos Monkey API"* ]]; then
    echo "   ✅ Backend API responding"
else
    echo "   ❌ Backend API not responding"
    kill $BACKEND_PID 2>/dev/null
    exit 1
fi

# Test experiments endpoint
EXPERIMENTS=$(curl -s http://localhost:8000/experiments)
if [[ $EXPERIMENTS == "[]" ]]; then
    echo "   ✅ Experiments endpoint working"
else
    echo "   ❌ Experiments endpoint failed"
fi

# Test creating experiment
CREATE_RESULT=$(curl -s -X POST http://localhost:8000/experiments \
    -H "Content-Type: application/json" \
    -d '{"name":"test","type":"pod_assassin","target_namespace":"default","target_labels":{},"blast_radius":50,"duration_seconds":60}')
if [[ $CREATE_RESULT == *"exp-"* ]]; then
    echo "   ✅ Can create experiments"
else
    echo "   ❌ Cannot create experiments"
fi

kill $BACKEND_PID 2>/dev/null
cd ..

# Test agent
echo ""
echo "2️⃣  Testing Agent..."
if [ -f "agent/chaos-agent" ]; then
    echo "   ✅ Agent binary exists"
else
    echo "   ❌ Agent binary not found (run: cd agent && go build)"
fi

# Test frontend
echo ""
echo "3️⃣  Testing Frontend..."
if [ -d "frontend/node_modules" ]; then
    echo "   ✅ Frontend dependencies installed"
else
    echo "   ❌ Frontend dependencies missing (run: cd frontend && npm install)"
fi

if [ -f "frontend/dist/index.html" ]; then
    echo "   ✅ Frontend builds successfully"
else
    echo "   ⚠️  Frontend not built yet (run: cd frontend && npm run build)"
fi

echo ""
echo "🎉 Component tests complete!"
echo ""
echo "To start development environment:"
echo "   ./start-dev.sh"
