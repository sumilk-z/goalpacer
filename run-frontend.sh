#!/bin/bash
cd /Users/zhucui/CodeBuddy/20251027000233/frontend
nohup npm start > /tmp/frontend.log 2>&1 &
echo $! > /tmp/frontend.pid
echo "前端已启动，PID: $(cat /tmp/frontend.pid)"
sleep 5
tail -20 /tmp/frontend.log
