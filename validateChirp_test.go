package main

import "testing"

// handlerValidateChirp
// {
//     "body": "This is a kerfuffle opinion I need to share with the world"
// }

// ->
// {
//     "cleaned_body": "This is a **** opinion I need to share with the world"
// }

func TestCensorProfaneWords(t *testing.T) {
	type TestCase struct {
		input string
		want  string
	}

	testCases := []TestCase{
		{"This is a kerfuffle opinion I need to share with the world", "This is a **** opinion I need to share with the world"},
		{"Nothing profane here", "Nothing profane here"},
		{"Fornax is an interesting constellation", "**** is an interesting constellation"},
		{"I love Sharbert ice-cream", "I love **** ice-cream"},
		{"Multiple words: kerfuffle sharbert fornax", "Multiple words: **** **** ****"},
		{"Multiple words: kerfuffle, sharbert, fornax.", "Multiple words: kerfuffle, sharbert, fornax."},
	}

	for _, testcase := range testCases {
		got := censorProfaneWords(testcase.input)
		if got != testcase.want {
			t.Errorf("For input \"%s\": got \"%s\", want \"%s\"", testcase.input, got, testcase.want)
		}
	}
}
