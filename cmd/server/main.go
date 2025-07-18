package main

import (
	"log"
	"github.com/sumayu/testovoe2/internal/api"
	"github.com/sumayu/testovoe2/internal/bd"
)

func main() {
	db,err := bd.Database()
	if err != nil {
		log.Fatalf("Failed to connected db", err)
	}
	defer db.Close()
	r := api.Router(db)
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}