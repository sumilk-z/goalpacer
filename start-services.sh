#!/bin/bash
set -e

echo "启动后端服务..."
cd /workspace/backend
go run main.go database.go models.go handlers.go &
BACKEND_PID=$!
echo "后端 PID: $BACKEND_PID"

sleep 3

echo "启动前端服务..."
cd /workspace/frontend
npm start
