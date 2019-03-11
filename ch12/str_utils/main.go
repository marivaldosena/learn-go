package strutils

import "strings"

/*
	Contains verifica se o parâmetro word possui o
	parâmetro pattern.
*/
func Contains(word, pattern string) bool {
	return strings.Contains(word, pattern)
}

/*
	Count retorna a quantidade de patterns em word.
*/
func Count(word, pattern string) int {
	return strings.Count(word, pattern)
}

/*
	BeginsWith verifica se word começa com pattern.
*/
func BeginsWith(word, pattern string) bool {
	return strings.HasPrefix(word, pattern)
}

/*
	EndsWith verifica se word termina com pattern.
*/
func EndsWith(word, pattern string) bool {
	return strings.HasSuffix(word, pattern)
}

/*
	Index verifica a posição do primeiro pattern em word.
*/
func Index(word, pattern string) int {
	return strings.Index(word, pattern)
}

/*
	Concat returna a junção dos argumentos de arr e utiliza sep
	como separador.
*/
func Concat(arr []string, sep string) string {
	return strings.Join(arr, sep)
}

/*
	Repeat retorna word de acordo com o número de times.
*/
func Repeat(word string, times int) string {
	return strings.Repeat(word, times)
}

/*
	Upper retorna word em letras maiúsculas.
*/
func Upper(word string) string {
	return strings.ToUpper(word)
}
