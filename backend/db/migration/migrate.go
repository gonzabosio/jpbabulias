package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func execQuery(db *sql.DB, queryPath string, wg *sync.WaitGroup) error {
	defer wg.Done()
	q, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(q))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load("../backend/.env"); err != nil {
		log.Fatalf("Failed to load enviroment variables: %v", err)
	}
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STR"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	basePath := "./db/migration/query"

	querieFiles := []string{"/create_table_patient.sql", "/create_table_appt.sql"}

	var wg sync.WaitGroup
	for _, f := range querieFiles {
		wg.Add(1)
		if err := execQuery(db, basePath+f, &wg); err != nil {
			log.Fatalf("%s query failed: %v", f, err)
		}
	}
	wg.Wait()
	log.Println("Database migration/update was successful")
}
