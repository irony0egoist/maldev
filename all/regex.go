package all

import "github.com/irony0egoist/maldev/src/regex"

func FindSubstringBetween(input string, first_delim string, second_delim string) []string {
	return regex.FindSubstringBetween(input, first_delim, second_delim)
}
