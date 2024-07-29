package mongodb

import (
	"strconv"
	"strings"
)

func (db db) decodeStrList(input string) ([]int, error) {

	inputList := strings.Split(input, separator)
	res := make([]int, len(inputList))
	for i, str := range inputList {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		res[i] = num
	}
	return res, nil
}
func (db db) encodeIntList(input []int) string {
	if len(input) == 0 {
		return ""
	}

	var res string
	for _, num := range input {
		str := strconv.Itoa(num)
		res += str + separator
	}
	res = res[:len(res)-1]
	return res
}

func (db db) findCloserRating(rating int) (begin int, end int) {
	roundedRating := rating - rating%ratingGroup

	return roundedRating - possibleDiff, roundedRating + possibleDiff
}
