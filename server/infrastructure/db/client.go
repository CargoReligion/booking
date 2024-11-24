package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DbClient interface {
	NamedGetSingleEntity(dest interface{}, query string, args ...interface{}) error
	NamedSelectEntities(dest interface{}, query string, args ...interface{}) error
	GetSingleEntity(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	ExecuteCommand(cmd string, args ...interface{}) (sql.Result, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

type dbClient struct {
	conn *sqlx.DB
}

func NewDbClient(db *sqlx.DB) DbClient {
	return &dbClient{
		conn: db,
	}
}

func (dc *dbClient) NamedGetSingleEntity(dest interface{}, query string, args ...interface{}) error {
	q, as, err := sqlx.Named(query, args)
	if err != nil {
		return err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	return dc.GetSingleEntity(dest, q, as...)
}
func (pc *dbClient) NamedSelectEntities(dest interface{}, query string, args ...interface{}) error {
	q, as, err := sqlx.Named(query, args)
	if err != nil {
		return err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	return pc.Select(dest, q, as...)
}

func (dc *dbClient) GetSingleEntity(dest interface{}, query string, args ...interface{}) error {
	return dc.conn.Get(dest, query, args...)
}

func (dc *dbClient) Select(dest interface{}, query string, args ...interface{}) error {
	return dc.conn.Select(dest, query, args...)
}

func (dc *dbClient) ExecuteCommand(cmd string, args ...interface{}) (sql.Result, error) {
	return dc.conn.Exec(cmd, args...)
}

func (dc *dbClient) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return dc.conn.NamedExec(query, arg)
}
