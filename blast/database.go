package blast

import (
	"bufio"
	"fmt"
	"strings"
)

// Database is the database of sequences on which we want to query
type Database [][]byte

/*func (db Database) kmerLoop(k int, cb func(Query)) {
	for i, entry := range db {
		for j := 0; j <= len(entry)-k; i++ {
			var q Query = entry[j : j+k]
			cb(q)
		}
	}
}*/

// PopulateDB builds an in-memory database from a file scanner
// Returns the database and an error if there was any during the scanning
func PopulateDB(sc *bufio.Scanner) (Database, error) {
	var db Database

	for sc.Scan() {
		db = append(db, sc.Bytes())
	}

	return db, sc.Err()
}

func (db Database) String() string {
	dbStr := []string{"", "Database:", "--------------"}

	for i, entry := range db {
		dbStr = append(dbStr, fmt.Sprintf("%d: %s", i+1, entry))
	}

	return strings.Join(dbStr, "\n")
}
