package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DbClient interface {
	Ping() error
	NamedGetSingleEntity(dest interface{}, query string, args ...interface{}) error
	NamedSelectEntities(dest interface{}, query string, args ...interface{}) error
	GetSingleEntity(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	ExecuteCommand(cmd string, args ...interface{}) (sql.Result, error)
}

type TransactionalDbClient interface {
	DbClient
	BeginTx() (TransactionalDbClient, error)
	Commit() error
	Rollback() error
}

type dbClient struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
}

func NewDbClient(db *sqlx.DB) DbClient {
	return &dbClient{
		conn: db,
	}
}

func (dc *dbClient) BeginTx() (TransactionalDbClient, error) {
	tx, err := dc.conn.Beginx()
	if err != nil {
		return nil, err
	}
	return &dbClient{tx: tx}, nil
}

func (dc *dbClient) Commit() error {
	if dc.tx != nil {
		return dc.tx.Commit()
	}
	return nil
}

func (dc *dbClient) Rollback() error {
	if dc.tx != nil {
		return dc.tx.Rollback()
	}
	return nil
}

func (pc *dbClient) Ping() error {
	return pc.conn.Ping()
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
	if dc.tx != nil {
		return dc.tx.Get(dest, query, args...)
	}
	return dc.conn.Get(dest, query, args...)
}

func (dc *dbClient) Select(dest interface{}, query string, args ...interface{}) error {
	if dc.tx != nil {
		return dc.tx.Select(dest, query, args...)
	}
	return dc.conn.Select(dest, query, args...)
}

func (dc *dbClient) ExecuteCommand(cmd string, args ...interface{}) (sql.Result, error) {
	if dc.tx != nil {
		return dc.tx.Exec(cmd, args...)
	}
	return dc.conn.Exec(cmd, args...)
}
