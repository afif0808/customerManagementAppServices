package interfaces

type IDBHandler interface {
	Execute(statement string, arguments ...interface{}) error
	Query(statement string, arguments ...interface{}) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}
