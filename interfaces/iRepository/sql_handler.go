// DIPを利用してinfrastructure層を直接呼び出すのではなく、SqlHandlerで定義した振る舞いを呼び出す
package iRepository

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
