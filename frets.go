package chords

import "strings"

type Note int

var noteNames = [12]string{"C", "C#", "D", "Eb", "E", "F", "F#", "G", "G#", "A", "Bb", "B"}

func (n Note) String() string {
	return noteNames[int(n)]
}

func (n Note) Inc() Note {
	return n.Move(1)
}

func (n Note) Dec() Note {
	return n.Move(-1)
}

func (n Note) Move(p int) Note {
	return Note((int(n) + p) % 12)
}

const (
	C Note = iota
	Cs
	D
	Eb
	E
	F
	Fs
	G
	Gs
	A
	Bb
	B
)

var Notes = []Note{C, Cs, D, Eb, E, F, Fs, G, Gs, A, Bb, B}

func FindNote(name string) (Note, bool) {
	switch strings.ToLower(name) {
	case "c":
		return C, true
	case "c#", "db":
		return Cs, true
	case "d":
		return D, true
	case "d#", "eb":
		return Eb, true
	case "e":
		return E, true
	case "f":
		return F, true
	case "f#", "gb":
		return Fs, true
	case "g":
		return G, true
	case "g#", "ab":
		return Gs, true
	case "a":
		return A, true
	case "a#", "bb":
		return Bb, true
	case "b":
		return B, true
	}

	return C, false
}

func noteIndex(note Note) int {
	for i, n := range Notes {
		if n == note {
			return i
		}
	}
	return -1
}

var offsets = [6]int{4, 9, 2, 7, 11, 4}

func fret(note Note, s int) int {
	return (noteIndex(note) - offsets[s] + 12) % 12
}
