package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/carlmjohnson/go-utils/normalizedsort"
)

var notes = [12]string{"C", "C#", "D", "Eb", "E", "F", "F#", "G", "G#", "A", "Bb", "B"}

func NoteName(index int) string {
	return notes[index%12]
}

func NoteIndex(note string) int {
	for i, n := range notes {
		if n == note {
			return i
		}
	}
	return -1
}

var offsets = [6]int{4, 9, 2, 7, 11, 4}

func NotePos(note string, s int) int {
	return (NoteIndex(note) - offsets[s] + 12) % 12
}

type Chord struct {
	Name string
	Pos  [][6]int
}

func (c Chord) String() string {
	s := c.Name

	for i := 5; i >= 0; i-- {
		s += "\n  -"
		for _, pos := range c.Pos {
			note := pos[i]
			if note == -1 {
				s += "-X--"
			} else {
				n := strconv.Itoa(note)
				if len(n) == 1 {
					s += "-" + n + "--"
				} else {
					s += n + "--"
				}
			}
		}
	}

	return s + "\n"
}

func Major7th(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + " Maj7",
		Pos: [][6]int{
			// X R 2 1 2 X
			{-1, index1, index1 + 2, index1 + 1, index1 + 2, -1},
			// X X R 2 2 2
			{-1, -1, index2, index2 + 2, index2 + 2, index2 + 2},
			// R X 1 1 0 X
			{index0, -1, index0 + 1, index0 + 1, index0, -1},
			// X R X 1 2 0
			{-1, index1, -1, index1 + 1, index1 + 2, index1},
		},
	}
}

func Major6th(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "Maj6",
		Pos: [][6]int{
			// X R 2 1 2 X
			{-1, index1, index1 + 2, index1 + 1, index1 + 2, -1},
			// X X R 2 0 2
			{-1, -1, index2, index2 + 2, index2, index2 + 2},
			// R X -1 1 0 X
			{index0, -1, index0 - 1, index0 + 1, index0, -1},
			// X R X -1 2 0
			{-1, index1, -1, index1 - 1, index1 + 2, index1},
		},
	}
}

func Minor7th(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "Min7",
		Pos: [][6]int{
			// X R 2 0 1 X
			{-1, index1, index1 + 2, index1, index1 + 1, index1},
			// X X R 2 1 1
			{-1, -1, index2, index2 + 2, index2 + 1, index2 + 1},
			// R X 0 0 0 X
			{index0, -1, index0, index0, index0, -1},
			// X R X 0 1 0
			{-1, index1, -1, index1, index1 + 1, index1},
		},
	}
}

func Minor6th(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "Min6",
		Pos: [][6]int{
			// X R 2 -1 1 X
			{-1, index1, index1 + 2, index1 - 1, index1 + 1, -1},
			// X X R 2 0 1
			{-1, -1, index2, index2 + 2, index2, index2 + 1},
			// R X -1 0 0 X
			{index0, -1, index0 - 1, index0, index0, -1},
			// X R X -1 1 0
			{-1, index1, -1, index1 - 1, index1 + 1, index1},
		},
	}
}

func Minor9th(root string) Chord {
	index1, index2 := NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "Min9",
		Pos: [][6]int{
			// X R -2 0 0 X
			{-1, index1, index1 - 2, index1, index1, -1},
			// X X R -2 1 0
			{-1, -1, index2, index2 - 2, index2 + 1, index2},
		},
	}
}

func Minor11th(root string) Chord {
	index0, index1 := NotePos(root, 0), NotePos(root, 1)

	return Chord{
		Name: root + "Min11",
		Pos: [][6]int{
			// R X 0 0 -2 X
			{index0, -1, index0, index0, index0 - 2, -1},
			// X R X 0 1 -2
			{-1, index1, -1, index1, index1 + 1, index1 - 2},
		},
	}
}

func Dominant7th(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "7",
		Pos: [][6]int{
			// X R 2 0 2 X
			{-1, index1, index1 + 2, index1, index1 + 2, -1},
			// X X R 2 1 2
			{-1, -1, index2, index2 + 2, index2 + 1, index2 + 2},
			// R X 0 1 0 X
			{index0, -1, index0, index0 + 1, index0, -1},
			// X R X 0 2 0
			{-1, index1, -1, index1, index1 + 2, index1},
		},
	}
}

func Dominant9th(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "9",
		Pos: [][6]int{
			// X R -1 0 0 X
			{-1, index1, index1 - 1, index1, index1, -1},
			// X X R -1 1 0
			{-1, -1, index2, index2 - 1, index2 + 1, index2},
			// R X 0 -1 -3 X
			{index0, -1, index0, index0 - 1, index0 - 3, -1},
			// X R X 0 0 -3
			{-1, index1, -1, index1, index1, index1 - 3},
		},
	}
}

func Dominant13th(root string) Chord {
	index0, index1 := NotePos(root, 0), NotePos(root, 1)

	return Chord{
		Name: root + "13",
		Pos: [][6]int{
			// X R -1 0 0 2
			{-1, index1, index1 - 1, index1, index1, index1 + 2},
			// X R X 0 2 2
			{-1, index1, -1, index1, index1 + 2, index1 + 2},
			// R X 0 1 2 X
			{index0, -1, index0, index0 + 1, index0 + 2, -1},
		},
	}
}

func Minor7thFlat5(root string) Chord {
	index0, index1, index2 := NotePos(root, 0), NotePos(root, 1), NotePos(root, 2)

	return Chord{
		Name: root + "Min7(b5)",
		Pos: [][6]int{
			// X R 1 0 1 0 X
			{-1, index1, index1 + 1, index1, index1 + 1, -1},
			// X X R 1 1 1 1
			{-1, -1, index2, index2 + 1, index2 + 1, index2 + 1},
			// R X 0 0 -1 X
			{index0, -1, index0, index0, index0 - 1, -1},
			// X R X 0 1 -1
			{-1, index1, -1, index1, index1 + 1, index1 - 1},
		},
	}
}

var Variants = []func(string) Chord{
	Major7th,
	Major6th,
	Minor7th,
	Minor6th,
	Minor9th,
	Minor11th,
	Dominant7th,
	Dominant9th,
	Dominant13th,
	Minor7thFlat5,
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func Random(root string) Chord {
	i := rnd.Intn(len(Variants))
	return Variants[i](root)
}

func main() {
	copyNotes := make([]string, 12)
	copy(copyNotes, notes[:])

	normalizedsort.Sort(copyNotes, nil)

	fmt.Println(Random(copyNotes[0]))
	fmt.Println(Random(copyNotes[1]))
	fmt.Println(Random(copyNotes[2]))
}
