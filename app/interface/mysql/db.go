package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/ent/migrate"
	"github.com/go-sql-driver/mysql"
)

const defaultMySQLPort = 3306

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string

	Loc       string
	Collation string
}

func NewDB(cfg Config) (*ent.Client, error) {

	connStr, err := buildConnectionString(cfg)
	if err != nil {
		return nil, err
	}

	client, err := ent.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		return nil, err
	}

	return client, nil
}

func buildConnectionString(cfg Config) (string, error) {
	mysqlCfg := mysql.NewConfig()

	if cfg.Host == "" {
		return "", errors.New("db host is not set")
	}

	if cfg.User == "" {
		return "", errors.New("db user is not set")
	}

	mysqlCfg.Net = "tcp"
	port := defaultMySQLPort
	if cfg.Port != 0 {
		port = cfg.Port
	}

	mysqlCfg.Addr = fmt.Sprintf("%s:%d", cfg.Host, port)

	mysqlCfg.DBName = cfg.Database
	mysqlCfg.User = cfg.User
	mysqlCfg.Passwd = cfg.Password

	mysqlCfg.ParseTime = true

	if cfg.Loc != "" {
		loc, err := time.LoadLocation(cfg.Loc)
		if err != nil {
			return "", err
		}
		mysqlCfg.Loc = loc
	}

	if cfg.Collation != "" {
		mysqlCfg.Collation = cfg.Collation
	}

	ret := mysqlCfg.FormatDSN()

	return ret, nil
}
