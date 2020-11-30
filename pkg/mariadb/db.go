package mariadb

import (
	"database/sql"
	"fmt"
)

func init() {
	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", "dellis:@/shud")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
