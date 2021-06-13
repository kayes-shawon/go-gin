package models

type Employee struct {
	tableName struct{} `pg:"employee"`
	Base
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Designation string `json:"designation"`
	Mobile      string `json:"mobile"`
	CompanyId   string `json:"company_id"`
}
