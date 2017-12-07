package converter_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"berlinclock/converter"
)

func TestGivenInvalidMinutesThenError(testing *testing.T) {
	const invalidArgumentMessage = "invalid argument"
	var functions = []struct {
		function     func(int) (string, error)
		invalidValue int
		expected     string
	}{
		{function: converter.ConvertSingleMinutes, invalidValue: -1, expected: invalidArgumentMessage},
		{function: converter.ConvertSingleMinutes, invalidValue: 60, expected: invalidArgumentMessage},
		{function: converter.ConvertSingleMinutes, invalidValue: 70, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveMinutes, invalidValue: -1, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveMinutes, invalidValue: -5, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveMinutes, invalidValue: 60, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveMinutes, invalidValue: 70, expected: invalidArgumentMessage},
		{function: converter.ConvertSingleHours, invalidValue: -1, expected: invalidArgumentMessage},
		{function: converter.ConvertSingleHours, invalidValue: 24, expected: invalidArgumentMessage},
		{function: converter.ConvertSingleHours, invalidValue: 30, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveHours, invalidValue: -1, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveHours, invalidValue: 60, expected: invalidArgumentMessage},
		{function: converter.ConvertFiveHours, invalidValue: 75, expected: invalidArgumentMessage},
		{function: converter.ConvertSecondsRow, invalidValue: 75, expected: invalidArgumentMessage},
		{function: converter.ConvertSecondsRow, invalidValue: 60, expected: invalidArgumentMessage},
		{function: converter.ConvertSecondsRow, invalidValue: -1, expected: invalidArgumentMessage},
	}
	for _, test := range functions {
		minutes, err := test.function(test.invalidValue)
		assert.Empty(testing, minutes)
		assert.EqualError	(testing, err, test.expected)
	}
}

func TestGivenAConverterWhenTimeIsEnteredThenCorrectSingleMinutesRow(testing *testing.T) {
	var singleMinuteTests = []struct {
		minutes  int    // input
		expected string // expected result
	}{
		{0, "OOOO"},
		{59, "YYYY"},
		{32, "YYOO"},
		{34, "YYYY"},
		{35, "OOOO"},
	}

	for _, test := range singleMinuteTests {
		minutes, err := converter.ConvertSingleMinutes(test.minutes)
		assert.Nil(testing, err)
		assert.Equal(testing, test.expected, minutes, fmt.Sprintf("single minutes for %d row should be %s", test.minutes, test.expected))
	}
}

func TestGivenAConverterWhenTimeIsEnteredThenCorrectFiveMinutesRow(testing *testing.T) {
	var fiveMinuteTests = []struct {
		minutes  int    // input
		expected string // expected result
	}{
		{0, "OOOOOOOOOOO"},
		{04, "OOOOOOOOOOO"},
		{05, "YOOOOOOOOOO"},
		{59, "YYRYYRYYRYY"},
		{23, "YYRYOOOOOOO"},
		{35, "YYRYYRYOOOO"},
	}

	for _, test := range fiveMinuteTests {
		minutes, err := converter.ConvertFiveMinutes(test.minutes)
		assert.Nil(testing, err)
		assert.Equal(testing, test.expected, minutes, fmt.Sprintf("five minutes value %d row should be %s", test.minutes, test.expected))
	}
}
func TestGivenAConverterWhenTimeIsEnteredThenCorrectSingleHourRow(testing *testing.T) {
	var singleHoursParams = []struct {
		hours    int    // input
		expected string // expected result
	}{
		{0, "OOOO"},
		{23, "RRRO"},
		{02, "RROO"},
		{8, "RRRO"},
		{14, "RRRR"},
	}

	for _, test := range singleHoursParams {
		minutes, err := converter.ConvertSingleHours(test.hours)
		assert.Nil(testing, err)
		assert.Equal(testing, test.expected, minutes, fmt.Sprintf("single seconds value %d row should be %s", test.hours, test.expected))
	}
}

func TestGivenAConverterWhenTimeIsEnteredThenCorrectFiveHourRow(testing *testing.T) {
	var fiveHoursParams = []struct {
		hours    int    // input
		expected string // expected result
	}{
		{0, "OOOO"},
		{23, "RRRR"},
		{02, "OOOO"},
		{8, "ROOO"},
		{16, "RRRO"},
	}

	for _, test := range fiveHoursParams {
		minutes, err := converter.ConvertFiveHours(test.hours)
		assert.Nil(testing, err)
		assert.Equal(testing, test.expected, minutes, fmt.Sprintf("five seconds value %d row should be %s", test.hours, test.expected))
	}
}

func TestGivenAConverterWhenTimeIsEnteredThenCorrectSecondsRow(testing *testing.T) {
	var secondsParams = []struct {
		seconds  int    // input
		expected string // expected result
	}{
		{0, "Y"},
		{1, "O"},
		{2, "Y"},
		{3, "O"},
		{59, "O"},
	}

	for _, test := range secondsParams {
		minutes, err := converter.ConvertSecondsRow(test.seconds)
		assert.Nil(testing, err)
		assert.Equal(testing, test.expected, minutes, fmt.Sprintf("seconds value %d row should be %s", test.seconds, test.expected))
	}
}