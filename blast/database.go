package blast

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
)

// Database is the database of sequences on which we want to query
type Database [][]byte

// Hit is a self-contained struct representing a hit between a db index and a query hash
type Hit struct {
	kmer    *Kmer
	dbIndex [2]int
}

// Pair represents an alignment between a query and a database sequence
type Pair struct {
	Alignment []byte
	Distance  int
	Sequence  int
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

// ExtendHit extends on both sides of the seed (k-mer) matched until:
// 1- we've extended to the full length of the query OR
// 2- the current distance is beyond the threshold
// NB: The distance factor is not the score used in the true BLAST
func (hit *Hit) ExtendHit(k, t int, q Query, db Database, pairs chan<- *Pair, wg *sync.WaitGroup) {
	defer wg.Done()

	sequence := db[hit.dbIndex[0]] // whole sequence hit in the database

	// go through the indexes of the k-mer in the query (in case it's repeated)
	for _, index := range hit.kmer.indexes {
		iQ, jQ, iDB, jDB, dist := index, index+k-1, hit.dbIndex[1], hit.dbIndex[1]+k-1, 0

		// we extend while all conditions are satisfied
		for dist <= t && (iQ > 0 || jQ < len(q)-1) {
			if iQ-1 >= 0 {
				iQ--
				iDB--
			}
			if jQ+1 < len(q) {
				jQ++
				jDB++
			}
			// extending goes outside of database sequence bounds, alignment is bad. stop!
			if iDB < 0 || jDB >= len(sequence) {
				dist = t + 1 // make dist bigger than the threshold
				break
			}
			// update distance score (not the real score in BLAST)
			if q[iQ] != sequence[iDB] {
				dist++
			}
			if q[jQ] != sequence[jDB] {
				dist++
			}
		}

		// distance is beyond threshold, go to next index
		if dist > t {
			continue
		}

		// this pair is good
		pairs <- &Pair{
			sequence[iDB : jDB+1],
			dist,
			hit.dbIndex[0],
		}
	}
}

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

func (hit *Hit) String() string {
	hitStr := []string{"", "Hit:", "--------------", hit.kmer.kmer, fmt.Sprintf("DB index: (%d,%d)", hit.dbIndex[0], hit.dbIndex[1])}
	return strings.Join(hitStr, "\n")
}

func (p *Pair) String() string {
	pairStr := []string{"", "Pair:", "--------------", fmt.Sprintf("%s\nDistance: %d\nSequence: %d", p.Alignment, p.Distance, p.Sequence)}
	return strings.Join(pairStr, "\n")
}
