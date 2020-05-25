package main

import (
	"github.com/stretchr/testify/assert"
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

func TestCalculatePercentage(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		expected float64
	}{
		{"Devuelve el porcentaje de vocales que hay en la palabra", args{"television"}, 0.5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			assert.Equal(t, test.expected, CalculatePercentage(test.args.word))
		})
	}
}

func TestConvertVowelsToMumy(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		expected string
	}{
		{"Devuelve la misma palabra si tiene menos de 30% de vocales", args{"Fred"}, "Fred"},
		{"Devuelve la misma palabra reemplazando cada vocal por mumy, si contiene más de 30% de vocales", args{"casa"}, "cmumysmumy"},

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			assert.Equal(t, test.expected, ConvertVowelsToMumy(test.args.word))
		})
	}
}



