package remove_duplicates

import (
	"log"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{1, 1, 2}

	n := RemoveDuplicates(a)

	log.Println(n)

}
