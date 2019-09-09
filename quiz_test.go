package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadQuizKey(t *testing.T) {
	want := []QuestionAnswer{{"1 + 1", "2"}, {"2 + 2", "4"}, {"3 + 3", "6"}, {"4 + 4", "8"}}
	got, err := readQuizKey("sample.csv")
	if err == nil {
		fmt.Println("want: ", want)
		fmt.Println("got: ", got)
		result := reflect.DeepEqual(got, want)
		if !result {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestGiveQuestion(t *testing.T) {
	quiz := QuestionAnswer{"1 + 1", "2"}
	var stdin bytes.Buffer

	stdin.Write([]byte("special\n"))

	badResult, err := GiveQuestion(quiz, &stdin)

	assert.NoError(t, err)
	assert.Equal(t, false, badResult)

	stdin.Reset()
	stdin.Write([]byte("2\n"))

	goodResult, err := GiveQuestion(quiz, &stdin)
	assert.NoError(t, err)
	assert.Equal(t, true, goodResult)
}

func TestReadUserResponse(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("special\n"))
	goodResult, goodErr := ReadUserResponse(&stdin)

	assert.NoError(t, goodErr)
	assert.Equal(t, "special\n", goodResult)

	stdin.Reset()
	stdin.Write([]byte(""))
	_, badErr := ReadUserResponse(&stdin)

	assert.Error(t, badErr)
}
