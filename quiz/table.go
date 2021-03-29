package quiz

import (
	"encoding/json"
	"net/http"

	"github.com/jesmendi/quizCapital/helper"
)

type Table struct {
	Result   []Result
	Position int
	Correct  int
}

var class []int
var sum = 0

// Get the Classification to send to the client side
func (c *Table) GetTable(w http.ResponseWriter) {
	position := getPosition()
	class = append(class, sum)
	c =
		&Table{
			Result:   Res,
			Position: position,
			Correct:  sum,
		}
	clearData()
	b, err := json.Marshal(c)
	helper.SendData(w, b, err)

}

// returns the percentage of the players that have beaten
func getPosition() int {

	if sum == 0 {
		return 0
	}

	if len(class) == 0 {
		return 100
	}
	var greater = 0

	for _, j := range class {
		if sum > j {

			greater = greater + 1
		}
	}
	total := 100 * float32(greater) / float32(len(class))
	return int(total)
}
