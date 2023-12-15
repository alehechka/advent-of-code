package twentythree

import (
	"strconv"
	"strings"
)

func Day15Problem2(inputs []string) string {
	codes := strings.Split(inputs[0], ",")

	hashmap := Day15PrepHASHMAP(codes)

	return strconv.Itoa(hashmap.FocusPower())
}

type Day15Hashmap map[int][]Day15Lens

func (h Day15Hashmap) FocusPower() (power int) {
	for boxIndex, box := range h {
		for lensIndex, lens := range box {
			focalPower := (boxIndex + 1) * (lensIndex + 1) * lens.FocalLength
			power += focalPower
		}
	}

	return
}

func Day15PrepHASHMAP(codes []string) Day15Hashmap {
	hashmap := make(Day15Hashmap)
	for _, code := range codes {
		lens := Day15ParseLens(code)
		box := hashmap[lens.HASH]

		var found bool
		for index, boxLens := range box {
			if boxLens.Label == lens.Label {

				found = true
				tempBox := box[0:index]
				if lens.Operation == '=' {
					tempBox = append(tempBox, lens)
				}
				box = append(tempBox, box[index+1:]...)
			}
		}

		if !found && lens.Operation == '=' {
			box = append(box, lens)
		}

		hashmap[lens.HASH] = box
	}

	return hashmap
}

type Day15Lens struct {
	Label       string
	HASH        int
	Operation   rune
	FocalLength int
}

func Day15ParseLens(str string) Day15Lens {
	if strings.HasSuffix(str, "-") {
		label := str[0 : len(str)-1]
		return Day15Lens{
			Label:     label,
			HASH:      Day15HASH(label),
			Operation: '-',
		}
	}

	parts := strings.Split(str, "=")
	focalLength, _ := strconv.Atoi(parts[1])
	return Day15Lens{
		Label:       parts[0],
		HASH:        Day15HASH(parts[0]),
		Operation:   '=',
		FocalLength: focalLength,
	}
}
