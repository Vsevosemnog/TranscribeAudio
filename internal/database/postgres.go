package database

import (
	"context"
	"fmt"

	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DBInContext = "ConnectToDatabaseInContext"

func GetDBFromContext(ctx context.Context) (*gorm.DB, error) {

	db := ctx.Value(DBInContext)
	if db == nil {
		return nil, fmt.Errorf("no connection to DB id in context")
	}

	return db.(*gorm.DB), nil
}

func ConnectPG(ctx context.Context, dsn string) *gorm.DB {
	//log := logging.GetLoggerFromCtx(ctx)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func ClosePG(ctx context.Context, db *gorm.DB) {

	dbConnection, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Close()
	if err != nil {
		log.Fatal(err)
	}
}
