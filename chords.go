package chords

import "strconv"

type Fretboard [][6]int

var Spacer = [][6]int{{-2, -2, -2, -2, -2, -2}}

func (f *Fretboard) Add(g Fretboard) {
	*f = append(*f, g...)
}

func (f Fretboard) String() string {
	s := ""

	for i := 5; i >= 0; i-- {
		s += "\n  -"
		for j, pos := range f {
			note := pos[i]
			if note == -1 {
				s += "-X--"
			} else if note == -2 {
				if j == len(f)-1 {
					s += "-  "
				} else {
					s += "-  -"
				}
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

type Chord struct {
	Name string
	Pos  [][6]int
}

func (c Chord) String() string {
	return c.Name
}

func (c Chord) Fretboard() Fretboard {
	return Fretboard(c.Pos)
}

func Major7th(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "maj7",
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

func Major6th(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "maj6",
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

func Minor7th(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "min7",
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

func Minor6th(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "min6",
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

func Minor9th(root Note) Chord {
	index1, index2 := fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "min9",
		Pos: [][6]int{
			// X R -2 0 0 X
			{-1, index1, index1 - 2, index1, index1, -1},
			// X X R -2 1 0
			{-1, -1, index2, index2 - 2, index2 + 1, index2},
		},
	}
}

func Minor11th(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "min11",
		Pos: [][6]int{
			// R X 0 0 -2 X
			{index0, -1, index0, index0, index0 - 2, -1},
			// X R X 0 1 -2
			{-1, index1, -1, index1, index1 + 1, index1 - 2},
		},
	}
}

func Dominant7th(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "7",
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

func Dominant9th(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "9",
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

func Dominant13th(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "13",
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

func Minor7thFlat5(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "min7(b5)",
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

func Diminished(root Note) Chord {
	index0, index1, index2 := fret(root, 0), fret(root, 1), fret(root, 2)

	return Chord{
		Name: root.String() + "dim7",
		Pos: [][6]int{
			// X R 1 -1 1 X
			{-1, index1, index1 + 1, index1 - 1, index1 + 1, -1},
			// X X R 1 0 1
			{-1, -1, index2, index2 + 1, index2, index2 + 1},
			// R X -1 0 -1 X
			{index0, -1, index0 - 1, index0, index0 - 1, -1},
			// X R X -1 1 -1
			{-1, index1, -1, index1 - 1, index1 + 1, index1 - 1},
		},
	}
}

func Dominant7thFlat5(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7(b5)",
		Pos: [][6]int{
			// R X 0 1 -1 X
			{index0, -1, index0, index0 + 1, index0 - 1, -1},
			// X R X 0 2 -1
			{-1, index1, -1, index1, index1 + 2, index1 - 1},
		},
	}
}

func Dominant7thSharp9(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7(#9)",
		Pos: [][6]int{
			// RX X 0 1 0 3
			{-1, -1, index0, index0 + 1, index0, index0 + 3},
			// X R -1 0 1 X
			{-1, index1, index1 - 1, index1, index1 + 1, -1},
		},
	}
}

func Dominant7thFlat13(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7(b13)",
		Pos: [][6]int{
			// R X 0 1 1 X
			{index0, -1, index0, index0 + 1, index0 + 1, -1},
			// X R X 0 2 1
			{-1, index1, -1, index1, index1 + 2, index1 + 1},
		},
	}
}

func Dominant7thFlat9Flat13(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7(b9,b13)",
		Pos: [][6]int{
			// R X 0 1 1 1
			{index0, -1, index0, index0 + 1, index0 + 1, index0 + 1},
			// X R -1 0 -1 1
			{-1, index1, index1 - 1, index1, index1 - 1, index1 + 1},
		},
	}
}

func Dominant7thFlat5Flat9(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7b5(b9)",
		Pos: [][6]int{
			// RX X 0 1 -1 1
			{-1, -1, index0, index0 + 1, index0 - 1, index0 + 1},
			// X R -1 0 -1 -1
			{-1, index1, index1 - 1, index1, index1 - 1, index1 - 1},
		},
	}
}

func Dominant7thFlat5Sharp9(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7b5(#9)",
		Pos: [][6]int{
			// RX X 0 1 -1 3
			{-1, -1, index0, index0 + 1, index0 - 1, index0 + 3},
			// X R -1 0 1 -1
			{-1, index1, index1 - 1, index1, index1 + 1, index1 - 1},
		},
	}
}

func Dominant7thSharp5Flat9(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7#5(b9)",
		Pos: [][6]int{
			// RX X 0 1 1 1
			{-1, -1, index0, index0 + 1, index0 + 1, index0 + 1},
			// X R -1 0 -1 1
			{-1, index1, index1 - 1, index1, index1 - 1, index1 + 1},
		},
	}
}

func Dominant7thSharp5Sharp9(root Note) Chord {
	index0, index1 := fret(root, 0), fret(root, 1)

	return Chord{
		Name: root.String() + "7#5(#9)",
		Pos: [][6]int{
			// RX X 0 1 1 3
			{-1, -1, index0, index0 + 1, index0 + 1, index0 + 3},
			// X R -1 0 1 1
			{-1, index1, index1 - 1, index1, index1 + 1, index1 + 1},
		},
	}
}

var Variants = []func(Note) Chord{
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
	Diminished,
	Dominant7thFlat5,
	Dominant7thSharp9,
	Dominant7thFlat13,
	Dominant7thFlat9Flat13,
	Dominant7thFlat5Flat9,
	Dominant7thFlat5Sharp9,
	Dominant7thSharp5Flat9,
	Dominant7thSharp5Sharp9,
}
