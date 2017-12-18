package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"

	"./blast"
)

func main() {
	f, err := os.Open("db.txt")
	if err != nil {
		fmt.Println("Error reading db file")
		return
	}

	db, err := blast.PopulateDB(bufio.NewScanner(f))
	if err != nil {
		fmt.Printf("Problem populating database: %v\n", err)
		return
	}
	fmt.Println(db)

	k := 3
	var q blast.Query = []byte("TTTACAGG")
	fmt.Println(q)
	t := (len(q) - k) / 4 // threshold
	fmt.Printf("\nThreshold:\n--------------\n%d .. based on (q-k)/4 = (%d-%d)/4\n", t, len(q), k)
	h := blast.HashQuery(q, k)
	//fmt.Println(h)

	hits := make(chan *blast.Hit)
	pairs := make(chan *blast.Pair)
	var hitsWg sync.WaitGroup

	go db.Scan(h, k, hits)
	go func() {
		// go through all the hits and extend them
		for hit := range hits {
			//fmt.Println(hit)

			hitsWg.Add(1)
			go hit.ExtendHit(k, t, q, db, pairs, &hitsWg)
		}

		// once all hits are extended, close the "pairs" channel
		hitsWg.Wait()
		close(pairs)
	}()

	// go over all the pairs; we know we've seen all of them when "pairs" is closed
	var allPairs []*blast.Pair
	for pair := range pairs {
		allPairs = append(allPairs, pair)
	}

	// remove duplicate pairs, and sort them in increasing order of distance
	uniquePairs := blast.GetUniquePairs(allPairs)
	sort.Slice(uniquePairs, func(i, j int) bool {
		return uniquePairs[i].Distance < uniquePairs[j].Distance
	})
	for _, uniquePair := range uniquePairs {
		fmt.Println(uniquePair)
	}
}
