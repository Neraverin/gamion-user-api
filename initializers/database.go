package initializers

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectoToDB() {
	caPath := os.Getenv("DB_CA_PATH")
	connstring := os.Getenv("DB_CONNECTION_STRING")

	connstring += fmt.Sprintf(" sslmode=require sslrootcert=%s", caPath)

	sqlDB, err := sql.Open("pgx", connstring)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open DB: %v\n", err)
		os.Exit(1)
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open DB: %v\n", err)
		os.Exit(1)
	}
}
