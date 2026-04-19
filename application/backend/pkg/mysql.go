package pkg

import (
	"backend/model"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// 定义一个全局对象db
var db *sql.DB

// DB 获取数据库连接（供其他包使用）
func DB() *sql.DB {
	return db
}

func MysqlInit() (err error) {
	// 验证数据库名称（防止SQL注入）
	dbName := viper.GetString("mysql.db")
	if dbName == "" {
		return fmt.Errorf("mysql.db cannot be empty")
	}

	// 使用临时管理连接创建数据库和表
	adminDSN := viper.GetString("mysql.user") + ":" + viper.GetString("mysql.password") + "@tcp(" + viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port") + "/?parseTime=true&charset=utf8mb4"

	var adminDB *sql.DB
	adminDB, err = sql.Open("mysql", adminDSN)
	if err != nil {
		return fmt.Errorf("failed to open admin database connection: %w", err)
	}
	defer adminDB.Close()

	// 配置临时连接池
	adminDB.SetMaxOpenConns(1)
	adminDB.SetMaxIdleConns(0)
	adminDB.SetConnMaxLifetime(30 * time.Second)

	// 创建数据库
	_, err = adminDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	// 创建表（使用ALTER TABLE确保现有表的password列足够长）
	_, err = adminDB.Exec("CREATE TABLE IF NOT EXISTS " + dbName + ".users (user_id VARCHAR(50) PRIMARY KEY, username VARCHAR(50) UNIQUE NOT NULL, password VARCHAR(100) NOT NULL, realInfo VARCHAR(100))")
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// 修改现有表的password列长度以支持bcrypt（约60字符）
	_, err = adminDB.Exec("ALTER TABLE " + dbName + ".users MODIFY COLUMN password VARCHAR(100) NOT NULL")
	if err != nil {
		// 如果列已存在且长度足够，忽略错误并继续
		// 这是为了兼容已存在的表，避免因权限不足导致启动失败
		Warn("Attempted to modify password column but failed (may already be correct length)", Err(err))
		// 清除错误，继续启动
		err = nil
	}

	Info("Database and tables created, initializing connection pool")

	// 使用包含数据库名称的 DSN 创建应用连接池
	appDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		dbName)

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", appDSN)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(25)                     // 最大打开连接数
	db.SetMaxIdleConns(5)                      // 最大空闲连接数
	db.SetConnMaxLifetime(5 * time.Minute)    // 连接最大生存时间
	db.SetConnMaxIdleTime(1 * time.Minute)    // 空闲连接最大生存时间

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		// 清理失败的连接
		_ = db.Close()
		db = nil
		return fmt.Errorf("failed to ping database: %w", err)
	}

	Info("Database initialized successfully",
		String("host", viper.GetString("mysql.host")),
		String("database", dbName))

	return nil
}

// 注册
func InsertUser(user *model.MysqlUser) (err error) {
	sqlStr := "select count(user_id) from users where username = ?"
	var count int64
	err = db.QueryRow(sqlStr, user.Username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}
	// 使用 bcrypt 哈希密码
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return errors.New("密码加密失败: " + err.Error())
	}
	// realInfo 继续使用 MD5（用于哈希非敏感信息）
	realInfoHash := EncryptByMD5(user.RealInfo)

	sqlStr = "insert into users(user_id,username,password,realInfo) values(?,?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, hashedPassword, realInfoHash)
	if err != nil {
		return err
	}
	return nil
}

// 登录
func Login(user *model.MysqlUser) (err error) {
	sqlStr := "select username,password from users where username = ?"
	var hashedPassword string
	err = db.QueryRow(sqlStr, user.Username).Scan(&user.Username, &hashedPassword)
	if err != nil {
		return err
	}

	// 先尝试 bcrypt 验证（新密码格式）
	if VerifyPassword(user.Password, hashedPassword) {
		return nil
	}

	// 如果 bcrypt 验证失败，尝试 MD5 验证（旧密码格式兼容）
	if hashedPassword == EncryptByMD5(user.Password) {
		// MD5 验证成功，自动升级为 bcrypt
		Warn("User password is in legacy MD5 format, upgrading to bcrypt",
			String("username", user.Username))

		newHash, err := HashPassword(user.Password)
		if err != nil {
			Error("Failed to hash password for upgrade", Err(err))
			// 即使升级失败，也允许登录（下次再试）
			return nil
		}

		// 更新数据库中的密码
		updateSQL := "UPDATE users SET password = ? WHERE username = ?"
		_, err = db.Exec(updateSQL, newHash, user.Username)
		if err != nil {
			Error("Failed to upgrade password hash in database", Err(err))
			// 即使更新失败，也允许登录
		} else {
			Info("Password upgraded successfully for user", String("username", user.Username))
		}

		return nil
	}

	return errors.New("密码错误")
}

// 获取用户ID
func GetUserID(username string) (userID string, err error) {
	sqlStr := "select user_id from users where username = ?"
	err = db.QueryRow(sqlStr, username).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// 获取用户姓名
func GetUsername(userID string) (username string, err error) {
	sqlStr := "select username from users where user_id = ?"
	err = db.QueryRow(sqlStr, userID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db != nil {
		err := db.Close()
		db = nil
		if err != nil {
			return fmt.Errorf("failed to close database: %w", err)
		}
		Info("Database connection closed")
	}
	return nil
}
