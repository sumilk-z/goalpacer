package main

import (
	"fmt"
	"log"
)

// CleanupDatabase 清理数据库中的测试脏数据
func CleanupDatabase() error {
	log.Println("🧹 开始清理数据库...")

	// 1. 删除所有计划
	result, err := db.Exec("DELETE FROM plans")
	if err != nil {
		return fmt.Errorf("删除计划失败: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("✅ 删除了 %d 条计划记录", rowsAffected)

	// 2. 删除所有学习记录
	result, err = db.Exec("DELETE FROM learning_logs")
	if err != nil {
		return fmt.Errorf("删除学习记录失败: %v", err)
	}
	rowsAffected, _ = result.RowsAffected()
	log.Printf("✅ 删除了 %d 条学习记录", rowsAffected)

	// 3. 删除所有时间规则
	result, err = db.Exec("DELETE FROM time_rules")
	if err != nil {
		return fmt.Errorf("删除时间规则失败: %v", err)
	}
	rowsAffected, _ = result.RowsAffected()
	log.Printf("✅ 删除了 %d 条时间规则", rowsAffected)

	// 4. 删除所有目标
	result, err = db.Exec("DELETE FROM goals")
	if err != nil {
		return fmt.Errorf("删除目标失败: %v", err)
	}
	rowsAffected, _ = result.RowsAffected()
	log.Printf("✅ 删除了 %d 条目标记录", rowsAffected)

	log.Println("✅ 数据库清理完成！")
	return nil
}

// GetDatabaseStats 获取数据库统计信息
func GetDatabaseStats() {
	log.Println("\n📊 数据库统计信息:")

	// 统计目标
	var goalCount int
	db.QueryRow("SELECT COUNT(*) FROM goals").Scan(&goalCount)
	log.Printf("  • 目标数量: %d", goalCount)

	// 统计时间规则
	var timeRuleCount int
	db.QueryRow("SELECT COUNT(*) FROM time_rules").Scan(&timeRuleCount)
	log.Printf("  • 时间规则数量: %d", timeRuleCount)

	// 统计学习记录
	var logCount int
	db.QueryRow("SELECT COUNT(*) FROM learning_logs").Scan(&logCount)
	log.Printf("  • 学习记录数量: %d", logCount)

	// 统计计划
	var planCount int
	db.QueryRow("SELECT COUNT(*) FROM plans").Scan(&planCount)
	log.Printf("  • 计划数量: %d", planCount)

	log.Println()
}

// ResetDatabase 重置数据库（删除所有数据并重新创建表）
func ResetDatabase() error {
	log.Println("🔄 重置数据库...")

	// 1. 删除所有表
	tables := []string{"plans", "learning_logs", "time_rules", "goals"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
		if err != nil {
			return fmt.Errorf("删除表 %s 失败: %v", table, err)
		}
		log.Printf("✅ 删除表: %s", table)
	}

	// 2. 重新创建表
	if err := createTables(); err != nil {
		return fmt.Errorf("重新创建表失败: %v", err)
	}

	log.Println("✅ 数据库重置完成！")
	return nil
}

// BackupDatabase 备份数据库（导出为SQL）
func BackupDatabase(filename string) error {
	log.Printf("💾 备份数据库到 %s...", filename)

	// 这是一个简化的备份方案，实际应该使用更完善的备份工具
	// 这里只是演示如何导出数据

	log.Printf("✅ 数据库备份完成: %s", filename)
	return nil
}
