#!/bin/bash
cd /Users/zhucui/CodeBuddy/20251027000233/backend
nohup go run main.go database.go models.go handlers.go > /tmp/backend.log 2>&1 &
echo $! > /tmp/backend.pid
echo "后端已启动，PID: $(cat /tmp/backend.pid)"
sleep 2
tail -20 /tmp/backend.log
