package infrastructures

import (
	"customerManagementAppServices/interfaces"
	"database/sql"
	"fmt"
)

type MysqlHandler struct {
	Conn *sql.DB
}

func (handler *MysqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) Query(statement string, arguments ...interface{}) (interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement, arguments...)
	if err != nil {
		fmt.Println(err)
		return new(MysqlRow), err
	}
	row := new(MysqlRow)
	row.Rows = rows

	return row, nil
}

type MysqlRow struct {
	Rows *sql.Rows
}

func (r MysqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MysqlRow) Next() bool {
	return r.Rows.Next()
}
