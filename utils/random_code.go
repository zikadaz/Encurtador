package utils

import (
	"math/rand/v2"
	"strings"
)


// Escolhe o elemento char em uma slice
func escolherChar(s []string) ([]string){
	if len(s) < 5 {
		return s
	}
	copia := make([]string, len(s))
	copy(copia, s)
	rand.Shuffle(len(copia), func(i, j int){
		copia[i], copia[j] = copia[j], copia[i]
	})
	return copia[:5]
}

func GeradorAleatorio() (string) {
	const defaultCode = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQ0123456789"
	sliced := strings.Split(defaultCode, "")
	slicedRandom := escolherChar(sliced)
	stringParams := strings.Join(slicedRandom, "")
	return stringParams
}