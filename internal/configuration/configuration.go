package configuration

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	postgresmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GetDBConnection(migrationsPath string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		GetConfiguration("DB_HOST"),
		GetConfiguration("DB_PORT"),
		GetConfiguration("DB_USER"),
		GetConfiguration("DB_NAME"),
		GetConfiguration("DB_PASSWORD"),
		GetConfiguration("DB_SSL"),
	)), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,       // Disable color
		},
	)})
	if err != nil {
		return nil, err
	}

	dbInstance, err := db.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgresmigrate.WithInstance(dbInstance, &postgresmigrate.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		GetConfiguration("DB_NAME"), driver)

	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return db, err
}

func GetConfiguration(configuration string) string {
	env := os.Getenv(configuration)
	if env == "" {
		logrus.WithFields(logrus.Fields{
			"err": fmt.Sprintf("%s key not found in the .env file", configuration),
		}).Warnln()
		return env
	}

	return env
}
