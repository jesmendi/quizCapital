package quiz

import (
	"strconv"
)

type Result struct {
	Answer  string
	Country string
	Capital string
	OK      bool
}

var Res []Result

// Checking if the answer is correct
func (r Result) FillValue(capital string) {

	i, error := strconv.ParseInt(capital, 10, 32)
	var answer = ""
	var cor = false
	if error != nil || i < 1 || i > 3 {
		answer = "Incorret choise"

	} else {
		i = i - 1
		answer = QuesData[len(Res)].Capital[i]
		cor = answer == QuesData[len(Res)].Correct
		if cor {
			sum = sum + 1
		}
	}
	r = Result{
		Answer:  answer,
		Country: QuesData[len(Res)].Country,
		Capital: QuesData[len(Res)].Correct,
		OK:      cor,
	}
	Res = append(Res, r)

}

// Clear the data when the game is over or the player reload the game
func clearData() {
	Res = []Result{}
	sum = 0

}
