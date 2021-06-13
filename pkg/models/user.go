package models

type User struct {
	tableName struct{} `pg:"users"`
	Base
	UserName string `json:"username" pg:"username"`
	Password string `json:"password" pg:"password"`
}


