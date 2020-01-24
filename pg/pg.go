package pg

import (
	"fmt"
	"log"
	"sync"
	"time"

	// _ "github.com/jackc/pgx/v4/stdlib" // postgres driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ysergeyru/go-task-foxg/config"
	"github.com/ysergeyru/go-task-foxg/logger"
)

// PgDB is a PostgreSQL DB instance
type PgDB struct {
	*sqlx.DB
}

var (
	PGDB *PgDB
	once sync.Once
)

// DB returns Postgre database handle
func DB() *PgDB {
	once.Do(func() {
		var err error
		PGDB, err = connect()
		if err != nil {
			log.Fatal(err)
		}
		KeepAlive()
	})
	return PGDB
}

func connect() (*PgDB, error) {
	cfg := config.Get()
	db, err := sqlx.Open("postgres", fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=%s", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPass))
	db.SetMaxIdleConns(cfg.PostgresMaxIdleConns)
	db.SetMaxOpenConns(cfg.PostgresMaxOpenConns)
	if err != nil {
		log.Fatal("sql.Open failed:", err)
	}
	// Make sure Postgres is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PgDB{db}, nil
}

// KeepAlive makes sure Postgres is alive and reconnects if needed
func KeepAlive() {
	logger := logger.Get()
	var err error
	go func() {
		for {
			time.Sleep(time.Second * 3)
			lostConnect := false
			if PGDB == nil {
				lostConnect = true
			} else if err := PGDB.Ping(); err != nil {
				lostConnect = true
			}
			if lostConnect {
				logger.Errorf("Lost PostgreSQL connection. Restoring...")
				PGDB, err = connect()
				if err != nil {
					logger.Error(err)
					continue
				}
				logger.Notice("Reconnected!")
			}
		}
	}()
}
