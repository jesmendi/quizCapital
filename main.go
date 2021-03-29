package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesmendi/quizCapital/config"
)

func main() {

	handleRequests()

}

// Handle the request received from the client side
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc(config.C.Conf.QuestionsUrl, getQuestions).Methods(http.MethodGet)
	myRouter.HandleFunc(config.C.Conf.AnswerUrl, handleAnswer).Methods(http.MethodPost)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
