package main

import (
	"fmt"
	"os"

	"hawx.me/code/chords"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: chords-progression KEY PROGRESSION...")
		return
	}

	key, ok := chords.FindNote(os.Args[1])
	if !ok {
		fmt.Println("No known note:", os.Args[1])
		return
	}
	progression := os.Args[2:]

	cs := make([]chords.Chord, len(progression))
	for i, p := range progression {
		switch p {
		case "I":
			cs[i] = chords.Major7th(key)
		case "II":
			cs[i] = chords.Major7th(key.Move(1))
		case "III":
			cs[i] = chords.Major7th(key.Move(2))
		case "IV":
			cs[i] = chords.Major7th(key.Move(3))
		case "V":
			cs[i] = chords.Major7th(key.Move(4))
		case "VI":
			cs[i] = chords.Major7th(key.Move(5))
		case "VII":
			cs[i] = chords.Major7th(key.Move(6))

		case "i":
			cs[i] = chords.Minor7th(key)
		case "ii":
			cs[i] = chords.Minor7th(key.Move(1))
		case "iii":
			cs[i] = chords.Minor7th(key.Move(2))
		case "iv":
			cs[i] = chords.Minor7th(key.Move(3))
		case "v":
			cs[i] = chords.Minor7th(key.Move(4))
		case "vi":
			cs[i] = chords.Minor7th(key.Move(5))
		case "vii":
			cs[i] = chords.Minor7th(key.Move(6))
		}
	}

	fmt.Println("Key of", key)
	fmt.Println(progression)
	fmt.Println(cs)

	frets := &chords.Fretboard{}

	for _, c := range cs {
		frets.Add(c.Fretboard())
		frets.Add(chords.Spacer)
	}

	fmt.Println(frets)
}
