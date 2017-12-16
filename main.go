package main

import (
	"bufio"
	"fmt"
	"os"

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

	k := 11
	var q blast.Query = []byte("ATTACAGGTCAGAGCTAGTCGATATGCAGTAGTCTAGACATGCGTATGCAGTAGTCGCTATCGCGATCGCGCGATATCGATATGTGAC")
	fmt.Println(q)
	h := blast.HashQuery(q, k)
	fmt.Println(h)

	hits := make(chan *blast.Hit)
	go db.Scan(h, k, hits)
	for hit := range hits {
		fmt.Println(hit)
	}
}
