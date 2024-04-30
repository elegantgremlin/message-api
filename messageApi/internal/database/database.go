package database

import (
	"context"
	"fmt"
	"messageApi/internal/types"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database is the interface for the data module
type Database interface {
	GetMessage(int) (types.Message, error)
	ListMessages() ([]types.Message, error)
	CreateMessage(types.Message) (types.Message, error)
	UpdateMessage(types.Message) (types.Message, error)
	DeleteMessage(int) error
}

// database is the implementation of the data module
type database struct {
	config types.Config
	conn   *pgxpool.Pool
}

// NewDatabase creates an instance of the data module
func NewDatabase(cfg types.Config) (Database, error) {
	db, err := initializeDatabase(cfg)
	if err != nil {
		return nil, err
	}

	return &database{cfg, db}, nil
}

// Used to verify connection pool is only initialized once
var (
	pgOnce sync.Once
	dbPool *pgxpool.Pool
	dbErr  error
)

// initializeDatabase creates a connection pool to the database and creates the Messages table
// Will only create the connection pool on first call and return the same pool for all calls
func initializeDatabase(cfg types.Config) (*pgxpool.Pool, error) {
	pgOnce.Do(func() {
		dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database)

		dbPool, dbErr = pgxpool.New(context.Background(), dbUrl)
		if dbErr != nil {
			return
		}

		// TODO: extract this, the API service should not be responsible for migrations
		createTableSql := `CREATE TABLE IF NOT EXISTS public.messages (
			id serial4 NOT NULL,
			message varchar(100) NULL,
			ispalindrome bool NULL,
			CONSTRAINT messages_pk PRIMARY KEY (id)
		);`

		_, dbErr = dbPool.Exec(context.Background(), createTableSql)
	})

	return dbPool, dbErr
}

// CreateMessage will INSERT the Message into the database
func (d *database) CreateMessage(msg types.Message) (types.Message, error) {
	args := pgx.NamedArgs{
		"message":      msg.Message,
		"ispalindrome": strconv.FormatBool(msg.IsPalindrome),
	}
	var id int
	err := d.conn.QueryRow(context.Background(), "INSERT INTO public.messages (message, ispalindrome) VALUES(@message, @ispalindrome) RETURNING id", args).Scan(&id)
	if err != nil {
		return types.Message{}, err
	}

	msg.Id = id

	return msg, nil
}

// GetMessage returns a single Message from the database
func (d *database) GetMessage(id int) (types.Message, error) {
	args := pgx.NamedArgs{
		"id": strconv.Itoa(id),
	}

	rows, err := d.conn.Query(context.Background(), "SELECT id, message, ispalindrome FROM public.messages WHERE id = @id", args)
	if err != nil {
		return types.Message{}, err
	}
	defer rows.Close()

	msgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.Message])
	if err != nil || len(msgs) == 0 {
		return types.Message{}, err
	}

	return msgs[0], nil
}

// ListMessages returns a full list of Messages in the database
func (d *database) ListMessages() ([]types.Message, error) {
	rows, err := d.conn.Query(context.Background(), "SELECT id, message, ispalindrome FROM public.messages")
	if err != nil {
		return []types.Message{}, err
	}
	defer rows.Close()

	msgs, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.Message])
	if err != nil {
		return []types.Message{}, err
	}

	return msgs, nil
}

// UpdateMessage performs an UPDATE on an existing Message in the database
func (d *database) UpdateMessage(msg types.Message) (types.Message, error) {
	args := pgx.NamedArgs{
		"id":           msg.Id,
		"message":      msg.Message,
		"ispalindrome": strconv.FormatBool(msg.IsPalindrome),
	}
	cmd, err := d.conn.Exec(context.Background(), "UPDATE public.messages SET message = @message, ispalindrome = @ispalindrome WHERE id = @id", args)
	if err != nil || cmd.RowsAffected() == 0 {
		return types.Message{}, err
	}

	return msg, nil
}

// DeleteMessage performs a DELETE on an existing message in the database
func (d *database) DeleteMessage(id int) error {
	args := pgx.NamedArgs{
		"id": id,
	}

	cmd, err := d.conn.Exec(context.Background(), "DELETE FROM public.messages WHERE id = @id", args)
	if err != nil {
		return err
	} else if cmd.RowsAffected() == 0 {
		return fmt.Errorf("no record found for id %d", id)
	}

	return nil
}
