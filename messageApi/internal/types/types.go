package types

// Message represents a data struct for passing between modules
type Message struct {
	Id           int    `db:"id" json:"id"`
	Message      string `db:"message" json:"message"`
	IsPalindrome bool   `db:"ispalindrome" json:"ispalindrome"`
}

// Config represents the configuration for the service
type Config struct {
	Db DbConnection
}

// DbConnection represents the values needed to connect to a database
type DbConnection struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
}
