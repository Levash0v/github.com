package main

import (
	"modem/repo"
)

func main() {
	config, err := configs.LoadConfig("/config/config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := repo.InitDB(config)
	if err != nil {
		panic(err)
	}

	cmd.Run(db)
}
