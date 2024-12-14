package pkg

import (
	"backend/model"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var db *sql.DB

func MysqlInit() (err error) {
	dsn := viper.GetString("mysql.user") + ":" + viper.GetString("mysql.password") + "@tcp(" + viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port") + ")/"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute)
	//goland:noinspection SqlNoDataSourceInspection
	_, _ = db.Exec("CREATE DATABASE IF NOT EXISTS " + viper.GetString("mysql.db"))
	_, _ = db.Exec("USE " + viper.GetString("mysql.db"))
	//goland:noinspection SqlNoDataSourceInspection
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (user_id VARCHAR(50) PRIMARY KEY, username VARCHAR(50) UNIQUE NOT NULL, password VARCHAR(50) NOT NULL, realInfo VARCHAR(100))")
	if err != nil {
		panic(err.Error())
	}
	dsn = viper.GetString("mysql.user") + ":" + viper.GetString("mysql.password") + "@tcp(" + viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port") + ")/" + viper.GetString("mysql.db")
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func InsertUser(user *model.MysqlUser) (err error) {
	//goland:noinspection SqlNoDataSourceInspection
	sqlStr := "select count(user_id) from users where username = ?"
	var count int64
	err = db.QueryRow(sqlStr, user.Username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}
	//goland:noinspection SqlNoDataSourceInspection
	sqlStr = "insert into users(user_id,username,password,realInfo) values(?,?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, EncryptByMD5(user.Password), EncryptByMD5(user.RealInfo))
	if err != nil {
		return err
	}
	return nil
}

func Login(user *model.MysqlUser) (err error) {
	//goland:noinspection SqlNoDataSourceInspection
	sqlStr := "select username,password from users where username = ?"
	var password string
	err = db.QueryRow(sqlStr, user.Username).Scan(&user.Username, &password)
	if err != nil {
		return err
	}
	if EncryptByMD5(user.Password) != password {
		return errors.New("密码错误")
	}
	return nil
}

func GetUserID(username string) (userID string, err error) {
	//goland:noinspection SqlNoDataSourceInspection
	sqlStr := "select user_id from users where username = ?"
	err = db.QueryRow(sqlStr, username).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func GetUsername(userID string) (username string, err error) {
	//goland:noinspection SqlNoDataSourceInspection
	sqlStr := "select username from users where user_id = ?"
	err = db.QueryRow(sqlStr, userID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
