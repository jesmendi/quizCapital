package main

import (
	"net/http"

	"github.com/jesmendi/quizCapital/config"
	"github.com/jesmendi/quizCapital/quiz"
)

func getQuestions(w http.ResponseWriter, r *http.Request) {

	if len(quiz.Res) == 0 {
		quiz.GetQuestions(quiz.Countries)
	}
	quiz.GetNextQuestion(w)
}

func handleAnswer(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	var result quiz.Result
	result.FillValue(r.FormValue("capital"))
	if len(quiz.Res) == config.C.Conf.QuestionsPerGame {
		var stand *quiz.Table
		stand.GetTable(w)
	} else {
		quiz.GetNextQuestion(w)
	}

}
