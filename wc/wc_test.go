package main

import (
	"testing"
)

func Test_byteCounts(t *testing.T) {

	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "it should return correct length on ascii strings",
			text: "hello",
			want: 5,
		}, {name: "it should take into account byte count in non ascii strings",
			text: "👨🏻‍💻😅",
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteCounts(tt.text); got != tt.want {
				t.Errorf("byteCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lineCounts(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "it should return correct line length",
			text: "hello this is a test \n There is one line here \n\n and another line there",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lineCounts(tt.text); got != tt.want {
				t.Errorf("lineCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_characterCounts(t *testing.T) {

	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "it should return correct length on ascii strings",
			text: "hello",
			want: 5,
		}, {name: "it should give character count even in byte heavy characters",
			text: "áè",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterCounts(tt.text); got != tt.want {
				t.Errorf("characterCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordCounts(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "it should return correct word length on simple sentences",
			text: "hello this is a test",
			want: 5,
		}, {name: "it should return correct word length on sentences with line breaks",
			text: "hello this is a test \n and we continue one line below",
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordCounts(tt.text); got != tt.want {
				t.Errorf("wordCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
