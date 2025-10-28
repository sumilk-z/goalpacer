#!/bin/bash

# 清理数据库脚本

cd /Users/zhucui/CodeBuddy/goalpacer/backend

echo "🧹 清理数据库中的测试脏数据..."
echo ""

# 使用 sqlite3 命令行工具清理数据库
sqlite3 goalpacer.db << EOF
-- 显示清理前的统计信息
.mode column
.headers on
SELECT '清理前的数据统计:' as info;
SELECT 'goals' as table_name, COUNT(*) as count FROM goals
UNION ALL
SELECT 'time_rules', COUNT(*) FROM time_rules
UNION ALL
SELECT 'learning_logs', COUNT(*) FROM learning_logs
UNION ALL
SELECT 'plans', COUNT(*) FROM plans;

-- 清理所有数据
DELETE FROM plans;
DELETE FROM learning_logs;
DELETE FROM time_rules;
DELETE FROM goals;

-- 显示清理后的统计信息
SELECT '' as blank;
SELECT '清理后的数据统计:' as info;
SELECT 'goals' as table_name, COUNT(*) as count FROM goals
UNION ALL
SELECT 'time_rules', COUNT(*) FROM time_rules
UNION ALL
SELECT 'learning_logs', COUNT(*) FROM learning_logs
UNION ALL
SELECT 'plans', COUNT(*) FROM plans;
EOF

echo ""
echo "✅ 数据库清理完成！"
echo ""
echo "📊 所有表已清空，可以开始新的测试了"
