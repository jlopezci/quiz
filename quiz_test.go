package main

import (
	"fmt"
	"reflect"
	"testing"
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
