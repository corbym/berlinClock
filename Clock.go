package berlinclock

import (
	"strings"
	"berlinclock/converter"
	"strconv"
	"regexp"
	"errors"
	"fmt"
)

var validTimeFormat = regexp.MustCompile(`^[0-9]*:[0-9]*:[0-9]*$`).MatchString
var lastError error

func Clock(timeInHHmmSS string) (string, error) {
	if !validTimeFormat(timeInHHmmSS) {
		return "", errors.New("invalid argument, " + timeInHHmmSS + " is invalid time format")
	}
	hhmmSS := strings.Split(timeInHHmmSS, ":")
	hours, err := strconv.Atoi(hhmmSS[0])
	minutes, err := strconv.Atoi(hhmmSS[1])
	seconds, err := strconv.Atoi(hhmmSS[2])
	if err != nil {
		return "", err
	}
	secondsRow := errorHandler(converter.ConvertSecondsRow(seconds))
	fiveHoursRow := errorHandler(converter.ConvertFiveHours(hours))
	singleHoursRow := errorHandler(converter.ConvertSingleHours(hours))
	fiveMinutesRow := errorHandler(converter.ConvertFiveMinutes(minutes))
	singleMinutesRow := errorHandler(converter.ConvertSingleMinutes(minutes))

	if lastError != nil {
		return "", lastError
	}
	return fmt.Sprintf("%s%s%s%s%s", secondsRow, fiveHoursRow, singleHoursRow, fiveMinutesRow, singleMinutesRow), nil
}

func errorHandler(result string, err error) (string) {
	if err != nil {
		lastError = err
	}
	return result
}
