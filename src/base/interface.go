package base

type DBTX interface {
	// ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	// PrepareContext(context.Context, string) (*sql.Stmt, error)
	// QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	// QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Set(query string, args map[string]interface{})
	Fetch(query string, args map[string]interface{})
	Retrieve(query string, args map[string]interface{})
	Create(query string, args map[string]interface{})
}