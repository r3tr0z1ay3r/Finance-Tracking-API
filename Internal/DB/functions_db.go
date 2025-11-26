package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Internal/Model"
	_ "modernc.org/sqlite" // Import SQLite3 driver
)

func createDb(db *sql.DB) error {
	// Function to create a table called transactions
	// Expects Database connection to be provided as input and returns any err if occured
	// Expected usecase : first run

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

	// Function to insert value onto the database
	// Expects database, (Amount , Mode of payment, Expense/Income) -> Should be parsed from API
	// Should return error if any else Adds value to database
	model := Model.Transaction
	currTime := time.Now()
	formattedTime := currTime.Format("2006-01-02 15:04:04")
	fmt.Printf("Current Time is %v\n", formattedTime)

	insertCmd := `INSERT INTO transactions(time, amt, flow, mode) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(insertCmd, formattedTime, amt, flow, mode)
	return err
}

func getValDb(db *sql.DB, month time.Month, year int) (*sql.Rows, error) {

	// Function to fetch values from the database
	// Designed to get values based on the month and year
	// Designed into months and years to get monthly transactions on a page

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

	//Helper function to print the rows gathered from getValDb()
	//Used while debugging
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

	// Function to delete values from the database
	// Expects the database connection (db) and id to delete from
	// Planning to delete val from database based on selection on the front-end,
	// Front-End should only show necessary info, ID will be abstracted but used here

	delCmd := fmt.Sprintf(`DELETE FROM transactions WHERE id = %d`, id)
	_, err := db.Exec(delCmd)
	return err

}
