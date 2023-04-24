package main

import "errors"

func Distance(a, b string) (int, error) {
	uno := []byte(a)
	dos := []byte(b)
	var err error
	var cont = 0
	if len(a) == len(b) {

		for i, _ := range uno {
			if uno[i] != dos[i] {
				cont++
			}
		}
	} else {
		err = errors.New("Error")
	}

	return cont, err
}

func main() {
	Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
}
