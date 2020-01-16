package interfaces

type IDBHandler interface {
	Execute(statement string)
	Query(statement string, arguments ...interface{}) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
