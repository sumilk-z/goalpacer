package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB 初始化数据库
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./goalpacer.db")
	if err != nil {
		return err
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("✅ 数据库连接成功")

	// 创建表
	if err := createTables(); err != nil {
		return err
	}

	return nil
}

// createTables 创建所有数据表
func createTables() error {
	schema := `
	-- 学习目标表
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		description TEXT,
		status TEXT NOT NULL DEFAULT 'active',
		deadline DATE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- 时间规则表
	CREATE TABLE IF NOT EXISTS time_rules (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		day_of_week INTEGER NOT NULL,
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- 学习记录表
	CREATE TABLE IF NOT EXISTS learning_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		goal_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		duration INTEGER,
		log_date DATE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (goal_id) REFERENCES goals(id)
	);

	-- 学习计划表
	CREATE TABLE IF NOT EXISTS plans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		goal_id INTEGER NOT NULL,
		plan_date DATE NOT NULL,
		content TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (goal_id) REFERENCES goals(id)
	);
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Printf("❌ 创建表失败: %v", err)
		return err
	}

	log.Println("✅ 数据表创建成功")
	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
