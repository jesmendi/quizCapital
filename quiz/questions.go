package quiz

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/jesmendi/quizCapital/config"
	"github.com/jesmendi/quizCapital/helper"
)

var Countries []CountryData

var QuesData []QuestionData

type QuestionData struct {
	Country string
	Correct string
	Capital []string
}

// Creating the Questions
// Create all questions when the game is started
func GetQuestions(cty []CountryData) {

	clearData()
	random := helper.GetRandomSlice(config.C.Conf.NumberAnswer, len(cty))
	split := helper.ArrayChunk(random, config.C.Conf.AnswerPerQuestion)
	QuesData = make([]QuestionData, 0)
	for _, seeds := range split {
		rand.Seed(time.Now().UnixNano())
		permutations := rand.Perm(config.C.Conf.AnswerPerQuestion)
		var answerCapitals = make([]string, 0)

		for _, p := range permutations {
			var answerCapital = cty[seeds[p]].Capital
			answerCapitals = append(answerCapitals, answerCapital)
		}

		q := QuestionData{
			Country: cty[seeds[0]].Country,
			Correct: cty[seeds[0]].Capital,
			Capital: answerCapitals,
		}
		QuesData = append(QuesData, q)
	}
}

// Send the next question to the player
func GetNextQuestion(w http.ResponseWriter) {
	data := QuesData[len(Res)]
	data.Correct = "xxxx"
	b, err := json.Marshal(data)
	helper.SendData(w, b, err)
}
