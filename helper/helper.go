package helper

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// get Random slice with the id of the questions(id fron csv file )
func GetRandomSlice(number int, total int) []int {
	rand.Seed(time.Now().UnixNano())
	q := rand.Perm((total))
	q = q[:number]
	fmt.Println(q)
	return q

}

// Split the slice
func ArrayChunk(arr []int, limit int) [][]int {
	batchTotal := make([][]int, 0)
	for i := 0; i < len(arr); i += limit {
		batch := arr[i:min(i+limit, len(arr))]
		batchTotal = append(batchTotal, batch)
	}

	return batchTotal
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Send Data to  the client side
func SendData(w http.ResponseWriter, b []byte, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
