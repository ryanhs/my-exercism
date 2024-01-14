package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

const MaxCupsPerRowPerChild = 2

// Define the Garden type here.
type Garden struct {
	cups          [][]string
	childrenNames []string
}

func CodeToPlantName(s string) string {
	switch s {
	case "G":
		return "grass"
	case "C":
		return "clover"
	case "R":
		return "radishes"
	case "V":
		return "violets"
	}

	return ""
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func IsValidFormat(diagram string) bool {
	// check prefix
	if !strings.HasPrefix(diagram, "\n") {
		return false
	}

	return true
}

func IsNameUnique(childrenNames []string) bool {
	mapNames := make(map[string]bool, len(childrenNames))

	for _, v := range childrenNames {
		_, exists := mapNames[v]

		if exists {
			return false
		}

		mapNames[v] = true
	}
	return true
}

func NewGarden(diagram string, childrenNamesUnordered []string) (*Garden, error) {
	if !IsValidFormat(diagram) {
		return nil, fmt.Errorf("diagram not valid")
	}

	if !IsNameUnique(childrenNamesUnordered) {
		return nil, fmt.Errorf("chilren names must be unique")
	}

	trimmedDiagram := strings.TrimSpace(diagram)
	rows := strings.Split(trimmedDiagram, "\n")
	cups := make([][]string, len(rows))

	cupPerRow := MaxCupsPerRowPerChild * len(childrenNamesUnordered)

	// sort alphabetically, but not modify original params
	tmp := make([]string, len(childrenNamesUnordered))
	copy(tmp, childrenNamesUnordered)
	childrenNames := sort.StringSlice(tmp)
	childrenNames.Sort()
	// fmt.Printf("Start pouring plants! cups len: %d\n", len(cups))

	i := 0
	for _, row := range rows {
		dotsPerRow := strings.Split(row, "")
		if len(dotsPerRow) != cupPerRow {
			return nil, fmt.Errorf("actual: %d, expected: %d, errors on dots count per row", len(dotsPerRow), cupPerRow)
		}

		cups[i] = dotsPerRow

		// check plants code
		for _, s := range cups[i] {
			if CodeToPlantName(s) == "" {
				return nil, fmt.Errorf("%s is not a right plan code", s)
			}
		}

		// next
		i++
	}

	return &Garden{
		cups:          cups,
		childrenNames: childrenNames,
	}, nil
}

func (g *Garden) ChildToIndexes(childName string) ([]int, error) {
	indexes := make([]int, MaxCupsPerRowPerChild)
	for i, v := range g.childrenNames {
		if v == childName {
			startIndex := i * MaxCupsPerRowPerChild
			for i := range indexes {
				indexes[i] = startIndex + i
			}
			return indexes, nil
		}
	}

	return []int{}, fmt.Errorf("%s childName not exists", childName)
}

func (g *Garden) Plants(childName string) ([]string, bool) {
	indexes, err := g.ChildToIndexes(childName)
	if err != nil {
		return []string{}, false
	}

	rowsCount := len(g.cups)
	childrenPlants := make([]string, len(indexes)*rowsCount)

	ci := 0
	for _, row := range g.cups {
		for _, i := range indexes {
			childrenPlants[ci] = CodeToPlantName(row[i])
			ci++
		}
	}

	return childrenPlants, true
}
