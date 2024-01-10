package hamming

import (
	"errors"
	"strings"
)

func Distance(a string, b string) (int, error) {
	aSplit := strings.Split(a, "")
	aSplitLen := len(aSplit)

	bSplit := strings.Split(b, "")
	bSplitLen := len(bSplit)

	if aSplitLen != bSplitLen {
		return 0, errors.New("not match")
	}

	var hamming = 0

	for i := 0; i < aSplitLen; i++ {
		if aSplit[i] != bSplit[i] {
			hamming++
		}
	}

	return hamming, nil
}
