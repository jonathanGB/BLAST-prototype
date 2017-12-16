package blast

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

// Query represents a query to the Database
type Query []byte

// Kmer represents the k-mer (or repeats at different indices of that k-mer) in the Hash
type Kmer struct {
	kmer  string
	index []int
}

// Hash is a custom-made hash table of Kmers
type Hash [][]*Kmer

// getHashIndex returns the index in the hash of the given query
func getHashIndex(q Query, h Hash) int {
	hashFunc := fnv.New32()
	hashFunc.Write(q)
	return int(hashFunc.Sum32()) % len(h)
}

// HashQuery builds a hash from a query and the size of the k-mers
func HashQuery(q Query, k int) Hash {
	n := len(q) - k + 1               // number of k-mers
	var h Hash = make([][]*Kmer, n*2) // init hash to twice the number of k-mers

	for i := 0; i < n; i++ {
		kmer := q[i : i+k]
		index, foundKmer := h.get(kmer)

		// we found an existing k-mer, add our index to the slice of indices (means this k-mer is repeated)
		if foundKmer != nil {
			foundKmer.index = append(foundKmer.index, i)
			continue
		}

		// no existing k-mer found, create a new one and store it in the hash
		h[index] = append(h[index], &Kmer{
			string(kmer),
			[]int{i},
		})
	}

	return h
}

// get finds the query in the hash
// if it's found, it returns the Kmer and the index in the underlying hash where it is
// if it's not found, returns nil
func (h Hash) get(q Query) (int, *Kmer) {
	index := getHashIndex(q, h)
	for _, currKmer := range h[index] {
		if currKmer.kmer == string(q) {
			return index, currKmer
		}
	}

	return index, nil
}

// has tells if the hash contains the query
func (h Hash) has(q Query) bool {
	_, kmer := h.get(q)
	return kmer != nil
}

func (h Hash) String() string {
	hashStr := []string{"", "Hash:", "--------------"}
	numCollisions := 0

	for i, kmers := range h {
		collidedKmers := []string{fmt.Sprintf("%d:", i)}
		if len(kmers) > 1 {
			numCollisions += len(kmers) - 1
		}
		for _, kmer := range kmers {
			var indexes []string

			for _, index := range kmer.index {
				indexes = append(indexes, strconv.Itoa(index))
			}

			collidedKmers = append(collidedKmers, fmt.Sprintf("(%s)", strings.Join(indexes, ", ")))
		}

		hashStr = append(hashStr, strings.Join(collidedKmers, " "))
	}

	hashStr = append(hashStr, fmt.Sprintf("Number of collisions: %d", numCollisions))
	return strings.Join(hashStr, "\n")
}

func (q Query) String() string {
	return strings.Join([]string{"", "Query:", "--------------", string(q)}, "\n")
}
