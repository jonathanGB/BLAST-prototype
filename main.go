package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

	"./blast"
)

type blastJSON struct {
	Query     string `json:"query"`
	IsVerbose bool   `json:"isVerbose"`
	Kmer      int    `json:"kMer"`
}

func (data *blastJSON) getBlastOutput(db blast.Database) (string, error) {
	var output []string

	k := data.Kmer
	var q blast.Query = blast.Query(data.Query)
	t := (len(q) - k) / 4 // threshold
	h := blast.HashQuery(q, k)

	if data.IsVerbose {
		output = append(output,
			db.String(),
			q.String(),
			fmt.Sprintf("\nThreshold:\n--------------\n%d .. based on (q-k)/4 = (%d-%d)/4\n", t, len(q), k),
			h.String(),
		)
	}

	hits := make(chan *blast.Hit)
	pairs := make(chan *blast.Pair)
	var hitsWg sync.WaitGroup

	go db.Scan(h, k, hits)
	go func() {
		// go through all the hits and extend them
		for hit := range hits {
			output = append(output, hit.String())

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
		output = append(output, uniquePair.String())
	}

	return strings.Join(output, "\n"), nil
}

func processBlast(db blast.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		var data blastJSON
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			fmt.Fprint(w, "Error parsing request!")
			return
		}

		output, err := data.getBlastOutput(db)
		if err != nil {
			fmt.Fprint(w, "Error getting BLAST results")
		} else {
			fmt.Fprint(w, output)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	path := r.URL.Path[1:]
	file := "index.html"
	if path == "blast.js" {
		file = "blast.js"
	}
	http.ServeFile(w, r, file)
}

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

	http.HandleFunc("/", index)
	http.HandleFunc("/blast", processBlast(db))
	http.ListenAndServe(":8080", nil)
}
