package main

import (
	"fmt"
	"math/rand"
	"time"

	"hawx.me/code/chords"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func Random(root chords.Note) chords.Chord {
	i := rnd.Intn(len(chords.Variants))
	return chords.Variants[i](root)
}

func Shuffle(s []chords.Note) {
	for i := len(s) - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	copyNotes := make([]chords.Note, 12)
	copy(copyNotes, chords.Notes[:])

	Shuffle(copyNotes)

	for i := 0; ; i = (i + 1) % 12 {
		chord := Random(copyNotes[i])
		fmt.Print(chord.Name)
		fmt.Scanf(" ")

		fmt.Println(chord)
	}
}
