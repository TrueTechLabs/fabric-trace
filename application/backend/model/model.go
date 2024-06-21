package model

// user_id INT PRIMARY KEY ,
// username VARCHAR(50) UNIQUE NOT NULL,
// `password` VARCHAR(50) NOT NULL,
// RealInfo VARCHAR(100)

type MysqlUser struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RealInfo string `json:"real_info"`
}
