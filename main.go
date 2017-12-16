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

	var q blast.Query = []byte("ATTACAGGTCAG")
	fmt.Println(q)
	h := blast.HashQuery(q, 3)
	fmt.Println(h)
}
