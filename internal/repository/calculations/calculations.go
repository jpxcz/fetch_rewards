package calculations

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func AlphanumericCharacters(s string, point int) int {
	cleaned := clearString(s)
	trimmed := strings.ReplaceAll(cleaned, " ", "")
	l := len(trimmed)
	total := l * point

	return total
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func IsZeroCents(v string, point int) (int, error) {
	t, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}

	if math.Trunc(t) == t {
		return point, nil
	}

	return 0, nil
}

func multipleOf(v int, m int) bool {
	is := v % m
	return is == 0
}

func TwoItemsPoints(l int, point int) int {
	v := l / 2
	t := v * point
	return t
}

func IsMultipleOf25(s string, point int) (int, error) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	m := t * 100
	j := int(m)

	if multipleOf(j, 25) {
		return point, nil
	}

	return 0, nil

}

func TrimmedDescriptionPoints(s string, price string) (int, error) {
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, err
	}
	trimmed := strings.TrimSpace(s)
	if !multipleOf(len(trimmed), 3) {
		return 0, nil
	}

	points := int(math.Ceil(p * 0.2))
	return points, nil
}

func DateIsOddPoints(s string, point int) (int, error) {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return 0, err
	}

	if multipleOf(date.Day(), 2) {
		return 0, nil
	}

	return point, nil
}

func BetweenTimePoints(s string, initalTime int, finalTime int, points int) (int, error) {
	t, err := time.Parse("15:04:05", s+":00")
	if err != nil {
		return 0, err
	}

	if t.Hour() >= initalTime && t.Hour() <= finalTime {
		return points, nil
	}

	return 0, nil
}
