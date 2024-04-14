package repo

import (
	"database/sql"
	"fmt"

	configs "github.com/Levash0v/github.com/config"

	"github.com/models"
	_ "github.com/lib/pq"
)

func InitDB(config *configs.Config) (*sql.DB, error) {
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		fmt.Println("Open err")
		return nil, err
	}

	createDB := `DROP TABLE IF EXISTS items;
	CREATE TABLE IF NOT EXISTS items (
		item_id VARCHAR(255),
		value VARCHAR(255)
	);
	`

	_, err = db.Exec(createDB)
	if err != nil {
		fmt.Println("Exec err")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping err")
		return nil, err
	}

	return db, nil
}

func Create(item models.Item, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO items (item_id, value) VALUES ($1, $2)", item.ID, item.Value)
	if err != nil {
		fmt.Println("Exec INSERT")
		return err
	}

	return nil
}

func Read(id string, db *sql.DB) *models.Item {
	var result models.Item
	db.QueryRow("SELECT item_id, value FROM items WHERE item_id = $1", id).Scan(&result)

	return &result
}
