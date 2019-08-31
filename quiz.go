package main

import (
	"encoding/csv"
	"errors"
	"os"
)

// QuestionAnswer holds the question and answer
type QuestionAnswer struct {
	Question string
	Answer   string
}

func main() {

}

func readQuizKey(filename string) ([]QuestionAnswer, error) {
	var quizKey []QuestionAnswer
	var returnError error
	file, openErr := os.Open(filename)
	if openErr != nil {
		returnError = errors.New(openErr.Error())
	}
	defer file.Close()

	lines, readErr := csv.NewReader(file).ReadAll()
	if readErr != nil {
		returnError = errors.New(readErr.Error())
	}

	for i := 0; i < len(lines); i++ {
		var problem QuestionAnswer
		problem.Question = lines[i][0]
		problem.Answer = lines[i][1]
		quizKey = append(quizKey, problem)
	}

	return quizKey, returnError
}
