package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func InitDb() {
	var err error
	dbContext, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, err = dbContext.Exec(`
		CREATE TABLE IF NOT EXISTS receipts (
			id TEXT,
			points INTEGER 
 		);
	`)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

//Mock store of data. Data stores in memory
func GetDataById(id string) int {

	dbContext, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	query, err := dbContext.Query("SELECT * FROM receipts WHERE id = '" + id + "'")
	if err != nil {
		log.Println("Query error")
		log.Fatal(err)
		panic(err)
	}
	
	var returnedId string
	var points int = -1
	for query.Next() {
		err = query.Scan(&returnedId, &points)
		if err != nil {
			log.Println(err)
			if err == sql.ErrNoRows {
				fmt.Println("No row matching ID")
			} else {

				panic(err)
			}
		}
	}
	
	return points
}

func StoreProcessedReceipt(id string, points int) {
	
	dbContext, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	insertStatement := `INSERT INTO receipts (id, points) VALUES (?, ?)`
	statement, _ := dbContext.Prepare(insertStatement)
	_, err = statement.Exec(id, points)
	if err != nil {
		log.Println("Insert failed")
		log.Fatal(err)
		panic(err)
	}

}