package blast

import (
	"bufio"
	"fmt"
	"strings"
)

// Database is the database of sequences on which we want to query
type Database [][]byte

// Hit is a self-contained struct representing a hit between a db index and a query hash
type Hit struct {
	kmer    *Kmer
	dbIndex [2]int
}

// Scan goes through the database and alerts the `hits` channel for every hit between the query and the database
func (db Database) Scan(h Hash, k int, hits chan<- *Hit) {
	for i, entry := range db { // loop through db sequences
		n := len(entry) - k + 1

		for j := 0; j < n; j++ { // loop k-mers of a sequence
			var q Query = entry[j : j+k]
			if _, hit := h.get(q); hit != nil {
				hits <- &Hit{
					hit,
					[2]int{i, j},
				}
			}
		}
	}
	close(hits)
}

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
		dbStr = append(dbStr, fmt.Sprintf("%d: %s", i, entry))
	}

	return strings.Join(dbStr, "\n")
}

func (h *Hit) String() string {
	hitStr := []string{"", "Hit:", "--------------", h.kmer.kmer, fmt.Sprintf("DB index: (%d,%d)", h.dbIndex[0], h.dbIndex[1])}
	return strings.Join(hitStr, "\n")
}
