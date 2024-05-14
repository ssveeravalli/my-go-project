package main

import (
	"database/sql"
	"fmt"
	"log"

	// adding underscore because we are not explicitly using it
	_ "github.com/lib/pq"
	_ "github.com/qustavo/dotsql"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type TeddyBear struct {
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	Occupation     string  `json:"occupation"`
	Characteristic string  `json:"characteristic"`
	Age            float32 `json:"age"`
}

type PicnicLocation struct {
	Location     string
	MaxOccupancy float64
	HasMusic     bool
}

// picnic locations
var picnicLocations = []PicnicLocation{
	{Location: "Monroe Park", MaxOccupancy: 6, HasMusic: true},
	{Location: "Golden Gate Bridge", MaxOccupancy: 5, HasMusic: true},
	{Location: "Bramberly Park", MaxOccupancy: 4, HasMusic: true},
	{Location: "Treehouse", MaxOccupancy: 3, HasMusic: false},
	{Location: "Virginia Beach", MaxOccupancy: 2, HasMusic: true},
}

// getPicnicLocations responds with the list of all PicnicLocations as JSON.
func getPicnicLocations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, picnicLocations)
}

func main() {

	/*
		func Open(driverName, dataSourceName string) (*DB, error)
		Open opens a database specified by its database driver name
		and returns a pointer to an sql.db object that represents a db
	*/
	connStr := "postgres://postgres:secret@localhost:5432/picnicdb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	/*
		deferring the close function so when the main function finishes its execution
		it will run the db.Close function to close off the database
	*/

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := tableExists(db)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !result {
		createTeddyBearTableIfDoesNotExist(db)
	}

	fmt.Printf("Welcome to the Teddy Bear Picnic! Select 1 to browse the Teddy Bear Databse or select 2 to add your own: ")
	var i int
	fmt.Scan(&i)

	if i == 1 {
		displayTable(db)
		fmt.Printf("Enter the ID of the teddy bear you'd like to select\n")
		var id int
		fmt.Scan(&id)

		userTeddyBear, err := getTeddyBearByID(db, id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Your teddy bear is: %s", userTeddyBear.Name)

	} else if i == 2 {
		fmt.Printf("Give your teddy bear a name\n")
		var name string
		fmt.Scan(&name)
		fmt.Printf("Give your teddy bear a color\n")
		var color string
		fmt.Scan(&color)
		fmt.Printf("Give your teddy bear a occupation\n")
		var occupation string
		fmt.Scan(&occupation)
		fmt.Printf("Give your teddy bear a characteristic\n")
		var characteristic string
		fmt.Scan(&characteristic)
		fmt.Printf("Give your teddy bear a age\n")
		var age float32
		fmt.Scan(&age)

		teddyBear := TeddyBear{name, color, occupation, characteristic, age}
		insertTeddyBear(db, teddyBear)
		displayTable(db)

	} else {
		fmt.Printf("Please select either 1 or 2")
	}
}

// db is a pointer to a sql db object
func createTeddyBearTableIfDoesNotExist(db *sql.DB) {
	query := "SELECT EXISTS (SELECT 1 FROM picnicdb WHERE teddy_bears = $1)"

	// Execute the query
	row := db.QueryRow(query, "teddy_bears")

	// Scan the result into a boolean variable
	var exists bool
	err := row.Scan(&exists)
	if err != nil {

	}

	sqlScript, err := ioutil.ReadFile("database/teddy_bear_database_setup.sql")
	if err != nil {
		log.Fatal(err)
	}

	sqlCommands := strings.Split(string(sqlScript), ";")

	// Execute SQL commands
	for _, cmd := range sqlCommands {
		if _, err := db.Exec(cmd); err != nil {
			log.Println("Error executing SQL command:", err)
			continue
		}
	}
}

/*
creates a query that inserts the teddy bear into the db table and we parameterize the values that were passing in the columns. call the query row function and pass in the values from the structural argument and scan the result into the primary key variable that we set up
*/
func insertTeddyBear(db *sql.DB, teddyBear TeddyBear) int {
	query :=
		`INSERT INTO teddy_bears (teddy_bear_name, color, occupation, characteristic, age)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var pk int
	err := db.QueryRow(query, teddyBear.Name, teddyBear.Color, teddyBear.Occupation, teddyBear.Characteristic, teddyBear.Age).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func getTeddyBearByID(db *sql.DB, id int) (*TeddyBear, error) {
	// Prepare the SQL query
	query := "SELECT teddy_bear_name, color, occupation, characteristic, age FROM teddy_bears WHERE id = $1"

	// Execute the query
	row := db.QueryRow(query, id)

	var teddyBear TeddyBear

	err := row.Scan(&teddyBear.Name, &teddyBear.Color, &teddyBear.Occupation, &teddyBear.Characteristic, &teddyBear.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("teddy bear with ID %d not found", id)
		}
		return nil, err
	}
	return &teddyBear, nil
}

// DisplayTable function to display a table from the database
func displayTable(db *sql.DB) error {
	// Execute query to fetch data from the table
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", "teddy_bears"))
	if err != nil {
		return err
	}
	defer rows.Close()

	// gets column names
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	for _, col := range columns {
		fmt.Printf("%s\t", col)
	}
	fmt.Println()

	// iterate over the rows
	for rows.Next() {
		rowData := make([]interface{}, len(columns))
		rowPointers := make([]interface{}, len(columns))
		for i := range rowData {
			rowPointers[i] = &rowData[i]
		}
		if err := rows.Scan(rowPointers...); err != nil {
			return err
		}
		// Print each field value
		for _, value := range rowData {
			fmt.Printf("%v\t", value)
		}
		fmt.Println()
	}
	return nil
}

func tableExists(db *sql.DB) (bool, error) {
	// Prepare the SQL query
	query := "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)"

	// Execute the query
	row := db.QueryRow(query, "teddy_bears")

	// Scan the result into a boolean variable
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
