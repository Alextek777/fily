package storage

import (
	"database/sql"
	"fmt"
	"os/exec"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/Alextek777/fily/src/internal/config"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(cfg *config.Config) (*PostgresStore, error) {
	err := createDbIfExists(cfg)
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s",
		cfg.Db.User,
		cfg.Db.DbName,
		cfg.Db.Password,
		cfg.Db.SslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// TODO ERASE BUILT IN PSQL CALLS
func createDbIfExists(cfg *config.Config) error {
	command := fmt.Sprintf("PGPASSWORD=%s psql -U postgres -l | grep %s | wc -l", cfg.Db.Password, cfg.Db.DbName)

	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(string(output[0:1]))
	if err == nil && i == 1 {
		return nil
	}

	command = fmt.Sprintf("PGPASSWORD=%s createdb -p %s -h %s -U %s -e %s", cfg.Db.Password, cfg.Db.Port, cfg.Db.Ip, cfg.Db.User, cfg.Db.DbName)
	_, err = exec.Command("bash", "-c", command).Output()

	return err
}

func (s *PostgresStore) InitStorage() error {
	return s.CreateTables()
}

func (s *PostgresStore) CreateTables() error {
	err := s.createLinkTables()

	if err != nil {
		return err
	}

	return err
}

func (s *PostgresStore) createLinkTables() error {
	query := `create table if not exists account (
		id 					serial primary key,
		first_name 			varchar(100),
		last_name 			varchar(100),
		number 				serial,
		encrypted_password 	varchar(100),
		balance 			serial CHECK (balance >= 0),
		created_at 			timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}
