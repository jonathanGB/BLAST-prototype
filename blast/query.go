package blast

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

type Query []byte
type Kmer struct {
	kmer  string
	index []int
}
type Hash [][]*Kmer

func getHashIndex(q Query, h Hash) int {
	hashFunc := fnv.New32()
	hashFunc.Write(q)
	return int(hashFunc.Sum32()) % len(h)
}

func HashQuery(q Query, k int) Hash {
	n := len(q) - k + 1
	var h Hash = make([][]*Kmer, n*2)

	for i := 0; i < n; i++ {
		kmer := q[i : i+k]
		index, foundKmer := h.get(kmer)
		if foundKmer != nil {
			foundKmer.index = append(foundKmer.index, i)
			continue
		}

		h[index] = append(h[index], &Kmer{
			string(kmer),
			[]int{i},
		})
	}

	return h
}

func (h Hash) get(q Query) (int, *Kmer) {
	index := getHashIndex(q, h)
	for _, currKmer := range h[index] {
		if currKmer.kmer == string(q) {
			return index, currKmer
		}
	}

	return index, nil
}

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
