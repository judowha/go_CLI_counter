package counter

import (
	"sort"
	"strconv"
	"strings"
)

type ResultEntry struct {
	word string
	freq int
}

type ResultsVO []ResultEntry

func (resultsVO ResultsVO) String() string {
	var sb strings.Builder
	for _, entry := range resultsVO {
		sb.WriteString(entry.word + ": " + strconv.Itoa(entry.freq) + "\n")
	}
	return sb.String()
}

func CountFreq(s string, caseSens bool) (result ResultsVO) {
	resultMap := make(map[string]int)
	if !caseSens {
		s = strings.ToLower(s)
	}
	// strings.Fields: splits the string s around each instance of one or more consecutive white space characters
	strList := strings.Fields(s)
	for _, str := range strList {
		resultMap[str] += 1
	}
	for key, val := range resultMap {
		result = append(result, ResultEntry{key, val})
	}
	return result
}

func SortResult(rawResult ResultsVO, sortCritira string) {
	if sortCritira == "freq" {
		sort.Slice(rawResult, func(i, j int) bool {
			return rawResult[i].freq > rawResult[j].freq
		})
	} else {
		sort.Slice(rawResult, func(i, j int) bool {
			return rawResult[i].word < rawResult[j].word
		})
	}
}
