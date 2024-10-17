package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input: "Hello world",
			output: []string{
				"hello",
				"world",
			},
		},
		{
			input: "DiD tHiS WorK",
			output: []string{
				"did",
				"this",
				"work",
			},
		},
	}

	for _, cs := range cases {
		actual := CleanInput(cs.input)
		if len(actual) != len(cs.output) {
			t.Errorf("Lengths do not match %v vs %v", actual, cs.output)
			continue
		}
		for idx := range cs.output {
			if cs.output[idx] != actual[idx] {
				t.Errorf("words did not match... %v vs %v", cs.output[idx], actual[idx])
			}

		}

	}
}
