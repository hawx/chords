package chords

type Note string

var Notes = [12]Note{"C", "C#", "D", "Eb", "E", "F", "F#", "G", "G#", "A", "Bb", "B"}

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
