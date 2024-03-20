package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "testdb"
)

func Connect() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func InitDB(config *configs.Config) (*sql.DB, error) {
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode))
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		fmt.Println("Open err")
		return nil, err
	}
	defer db.Close()
	
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
