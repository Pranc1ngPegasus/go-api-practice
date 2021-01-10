package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"

	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/configuration"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

var _ RDBConnector = (*rdbConnector)(nil)

type (
	RDBConnector interface {
		GetDB() *sql.DB
		GetContext() context.Context
		Close()
	}

	rdbConnector struct {
		config  configuration.Config
		db      *sql.DB
		context context.Context
	}
)

func NewRDBConnector(config configuration.Config) RDBConnector {
	dns := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?parseTime=true",
		config.DatabaseUser,
		config.DatabasePass,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
	)

	sqltrace.Register(
		"mysql",
		&mysql.MySQLDriver{},
		sqltrace.WithServiceName(configuration.Get().DatabaseHost),
	)

	db, err := sqltrace.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	// Connection Poolの設定
	// db.SetMaxIdleConns(1)

	return &rdbConnector{
		config:  config,
		db:      db,
		context: context.Background(),
	}
}

func (dbc *rdbConnector) GetDB() *sql.DB {
	return dbc.db
}

func (dbc *rdbConnector) GetContext() context.Context {
	return dbc.context
}

func (dbc *rdbConnector) Close() {
	dbc.db.Close()
}
