package main

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestIsVowel(t *testing.T) {
	type args struct {
		letter string
	}
	tests := []struct {
		name string
		args args
		expected bool
	}{
		{"Devuelve true si es una vocal", args{"a"}, true},
		{"Devuelve false si no es una vocal", args{"b"}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			assert.Equal(t, test.expected, IsVowel(test.args.letter))
		})
	}
}

func TestVowelConverter(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		expected string
	}{
		{"Devuelve la misma palabra si tiene menos de 30% de vocales", args{"Fred"}, "Fred"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			assert.Equal(t, test.expected, IsVowel(test.args.word))
		})
	}
}

