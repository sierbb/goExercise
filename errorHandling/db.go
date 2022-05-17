package errorHandling

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func getDB(drive string, source string) (*sql.DB, error) {
	// Opening a driver typically will not attempt to connect to the database.
	return sql.Open(drive, source)
}

func queryRowbyName(db *sql.DB, name string) (bool, error) {
	// from current DB, check whether the provided name exists
	err := db.QueryRow(`
select
	p.name
from
	people as p
;`).Scan(&name)
	if err != nil{
		// we expect this error when there is no match in DB so deals with it
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}else {
			// wraps the original error
			return false, fmt.Errorf("error when scanning db for name %s: %w", name, err)
		}
	}
	return true, nil
}

func main(){
	db, err := getDB("drive-name", "databsase-test1")
	if err != nil{
		log.Fatal(err)
		return
	}
	name := "mary"
	hasRow, err := queryRowbyName(db, name)
	if err != nil{
		log.Fatal(err)
		return
	}
	if hasRow{
		log.Printf("Name %s is found\n", name)
	}else {
		log.Printf("Name %s is not found\n", name)
	}
}

