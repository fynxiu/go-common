package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/fynxiu/go-common/log"
	xtime "github.com/fynxiu/go-common/time"
	"strings"
	"time"
)

// Config database config
type Config struct {
	DSN         string         // database source connection
	Active      int            // active pool
	Idle        int            // idle pool
	IdleTimeout xtime.Duration // idle connection max life time
}

type ormLog struct{}

// Print ormLog print
func (r ormLog) Print(v ...interface{}) {
	log.Info(strings.Repeat("%v ", len(v)), v...)
}

func init() {
	// TODO: custom error message
	// gorm.ErrRecordNotFound = ecode.NothingFound
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		_ = log.Error("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	db.SetLogger(ormLog{})
	return
}
