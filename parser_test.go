package main

import "testing"

func TestBadExamples(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name: "five blocks",
			input: "####\n" +
				"...#\n" +
				"....\n" +
				"....\n",
		},
		{
			name: "diagonal blocks",
			input: "...#\n" +
				"..#.\n" +
				".#..\n" +
				"#...\n",
		},
		{
			name: "disconnected pairs",
			input: "...#\n" +
				"...#\n" +
				"#...\n" +
				"#...\n",
		},
		{
			name: "no blocks",
			input: "....\n" +
				"....\n" +
				"....\n" +
				"....\n",
		},
		{
			name: "separated blocks",
			input: "..##\n" +
				"....\n" +
				"....\n" +
				"##..\n",
		},
		{
			name: "incorrect separators",
			input: "...#\n" +
				"...#\n" +
				"...#\n" +
				"...#\n" +
				"....\n" +
				"....\n" +
				"....\n" +
				"####\n" +
				"\n\n" +
				".###\n" +
				"...#\n" +
				"....\n" +
				"....\n" +
				"\n" +
				"....\n" +
				"..##\n" +
				".##.\n" +
				"....\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ParseTetrominoes(test.input)

			if err == nil {
				t.Errorf("expected invalid input to return an error")
			}
		})
	}
}
