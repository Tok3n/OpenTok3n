package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/**
	Database tools
**/
func _getDB()(*sql.DB, error){
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",configData.DBUser, configData.DBPassword,configData.DBAddress,configData.DBPort,configData.DBName))
}

func _DBExcecute(db *sql.DB,query string) error{
	_,err := db.Query(query)
	return err
}
