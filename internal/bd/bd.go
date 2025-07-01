package bd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
func Database () (*sql.DB)  {
	err := godotenv.Load("../../configs/config.env")
	if err != nil {
		 fmt.Printf("error loading config.env file: %v", err)
	return nil
	}
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
	)
	fmt.Println("Connecting to database with connection string:",)
	db, err := sql.Open("postgres",dsn)
	if err != nil {
		 fmt.Printf("failed to connect to database: %v", err)
		return nil
	}
err = db.Ping()
	if err != nil { fmt.Printf("failed to ping database: %v", err)
		return nil
	}
	log.Println("Successfully connected to database")
	
	return db
}
	





// т.к в условии задачи сказанно о асинхронном выполнении оппераций чтобы избежать состояния гонки нужно обязатльно
//  в запросах блокировать другие транзакции до зовершения текущей 
// TO-DO в бд должен быть
// 1 wallet со своим wallet_uuid (ID нужен для get запроса))
// нужно создать таблицу WALLETS 
//функции из handler должны будут импортировать данные отсюда  
// для get запроса нужна функция которая по wallet_uuid выводит количестно денег. select money from wallets where id = 
// нашему id 