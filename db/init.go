package db

import (
	"database/sql"
	"log"

	"battlebit_admin_panel/migration"

	_ "github.com/go-sql-driver/mysql" // import mysql
	"github.com/jmoiron/sqlx"
	_ "github.com/kevinburke/go-bindata/v4" // so it's detected by `dep ensure`
	errors "github.com/pludderio/errs"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Init initializes a database connection based on the dsn provided. It also sets it as the global db connection.
func Init(dsn string, debug bool) error {
	dsn += "?parseTime=1&collation=utf8mb4_unicode_ci&loc=Local"
	dbConn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return errors.Err(err)
	}

	err = dbConn.Ping()
	if err != nil {
		return errors.Err(err)
	}

	if debug {
		boil.DebugMode = true
	}

	boil.SetDB(dbConn)

	migrations := &migrate.AssetMigrationSource{
		Asset:    migration.Asset,
		AssetDir: migration.AssetDir,
		Dir:      "migration",
	}
	n, migrationErr := migrate.Exec(dbConn.DB, "mysql", migrations, migrate.Up)
	if migrationErr != nil {
		return errors.Err(migrationErr)
	}
	log.Printf("Applied %d migrations", n)

	return nil
}

func dbInitConnection(dsn string, driverName string, debug bool) (*sqlx.DB, error) {
	dsn += "?parseTime=1&collation=utf8mb4_unicode_ci&loc=Local"
	dbConn, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		return nil, errors.Err(err)
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, errors.Err(err)
	}

	return dbConn, nil
}

// CloseDB is a wrapper function to allow error handle when it is usually deferred.
func CloseDB(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logrus.Error("Closing DB Error: ", err)
	}
}

// CloseRows Closes SQL Rows for custom SQL queries.
func CloseRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		logrus.WithError(err).Errorf("could not close DB rows")
	}
}
