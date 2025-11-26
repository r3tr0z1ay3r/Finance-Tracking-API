//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // Import SQLite3 driver
)

func connectDb() (*sql.DB, error) {

	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {

		return nil, err

	}

	if err = db.Ping(); err != nil {

		return nil, err

	}
	return db, err

}

func createDb(db *sql.DB) error {

	//Create a table
	createTable := `CREATE TABLE IF NOT EXISTS transactions (
						id	INTEGER PRIMARY KEY AUTOINCREMENT,
						time DATETIME NOT NULL,
						amt REAL NOT NULL,
						flow TEXT NOT NULL,
						mode TEXT 
	);`

	_, err := db.Exec(createTable)
	return err

}

func insertDb(db *sql.DB, amt float64, mode string, flow string) error {

	currTime := time.Now()
	formattedTime := currTime.Format("2006-01-02 15:04:04")
	fmt.Printf("Current Time is %v\n", formattedTime)

	insertCmd := `INSERT INTO transactions(time, amt, flow, mode) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(insertCmd, formattedTime, amt, flow, mode)
	return err
}

func getValDb(db *sql.DB, month time.Month, year int) (*sql.Rows, error) {
	//Creating a map to covert month in english to integer {String(Month) -> Int(month)}
	months := map[time.Month]int{
		time.January:   1,
		time.February:  2,
		time.March:     3,
		time.April:     4,
		time.May:       5,
		time.June:      6,
		time.July:      7,
		time.August:    8,
		time.September: 9,
		time.October:   10,
		time.November:  11,
		time.December:  12,
	}

	Month := months[month]
	yr_mth := fmt.Sprintf("%v-%v", year, Month)
	cmd := fmt.Sprintf(`
		SELECT *
		FROM transactions
		WHERE strftime('%%Y-%%m',time) = '%v'`, yr_mth)

	val, err := db.Query(cmd)
	if err != nil {

		log.Fatalf("Error while fetching rows : %v ", err)

	}
	defer val.Close()

	for val.Next() {

		var id int
		var time time.Time
		var amt float64
		var flow string
		var mode string

		err := val.Scan(&id, &time, &amt, &flow, &mode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v | %v | %v | %v | %v | %v | %v |\n", id, time, amt, flow, mode)

	}
	return val, err

}

func printVals(rows *sql.Rows) {

	for rows.Next() {

		var id int
		var time time.Time
		var amt float64
		var flow string
		var mode string

		err := rows.Scan(&id, &time, &amt, &flow, &mode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v |\n", id, time, amt, flow, mode)
	}

}

func deleteValDb(db *sql.DB, id int) error {

	delCmd := fmt.Sprintf(`DELETE FROM transactions WHERE id = %d`, id)
	_, err := db.Exec(delCmd)
	return err

}
func main() {

	db, err := connectDb()
	if err != nil {

		log.Fatalf("Error in connecting to DB :%v\n", err)

	}
	defer db.Close()

	if err = createDb(db); err != nil {

		log.Fatalf("Error while creating db: %v\n", err)

	}

	// if err = insertDb(db, 40, "UPI", "Expense"); err != nil {

	// 	log.Fatalf("Error in inserting value to db : %v\n", err)

	// }

	// currMonth, Year := time.Now().Month(), time.Now().Year()
	// rows, err := getValDb(db, currMonth, Year)

	// if err != nil {

	// 	log.Fatalf("Error while fetching rows : %v ", err)

	// }
	// defer rows.Close()

	if err = deleteValDb(db, 8); err != nil {
		log.Fatal(err)
	}

}
