package repositories

import (
	"database/sql"

	// load mysql driver
	"github.com/fguy/helloworld-go/config"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

const driverName = "mysql"

// NewMySQL -
func NewMySQL(logger *zap.Logger, cfg *config.AppConfig) (func() (*sql.DB, error), error) {
	return func() (*sql.DB, error) {
		db, err := sql.Open(driverName, cfg.DSN)
		if err != nil {
			logger.Error("can not connnect to the database", zap.Error(err))
			return nil, err
		}
		return db, nil
	}, nil
}
