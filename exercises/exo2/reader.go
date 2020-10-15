package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

func (x spaceEraser) Read(L []byte) (int, error) {
	length, err := x.r.Read(L)
	j := 0 //compteur
	for i := 0; i < length; i++ {
		if L[i] != 32 { // test si le charactère n'est pas un espace
			L[j] = L[i] // on l'ajoute dans la liste L
			j++
		}
	}
	return j, err
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
