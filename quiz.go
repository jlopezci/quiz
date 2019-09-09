package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// QuestionAnswer holds the question and answer
type QuestionAnswer struct {
	Question string
	Answer   string
}

func main() {
	quizKey, err := readQuizKey("problems.csv")
	if err != nil {
		fmt.Println("Error reading quiz file: ", err.Error())
	} else {
		PlayGame(quizKey)
	}
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

//PlayGame - execute the quiz for the user and total their results
func PlayGame(quiz []QuestionAnswer) {

}

// GiveQuestion - list out the question and verify the answer from the user
func GiveQuestion(problem QuestionAnswer, stdin io.Reader) (bool, error) {
	result := false
	var returnError error

	fmt.Print("Question: ", problem.Question)
	answer, err := ReadUserResponse(stdin)
	if err != nil {
		returnError = errors.New(err.Error())
	} else {
		// strip the newline from the answer
		answer = strings.TrimSuffix(answer, "\n")

		if answer == problem.Answer {
			result = true
		}
	}

	fmt.Println("")

	return result, returnError
}

// ReadUserResponse reads user input
func ReadUserResponse(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	text, err := reader.ReadString('\n')
	return text, err
}
