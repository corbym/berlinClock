package converter_test

import (
	"testing"
	"berlinclock/converter"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
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
		actual, err := test.function(test.invalidValue)
		then.AssertThat(testing, actual, is.Empty())
		then.AssertThat(testing, err.Error(), is.EqualTo(test.expected).Reasonf("error message for %v incorrect", test))
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
		then.AssertThat(testing, err, is.Nil())
		then.AssertThat(testing, minutes, is.EqualTo(test.expected).Reasonf("single minutes row %d incorrect", test.minutes))
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
		then.AssertThat(testing, err, is.Nil())
		then.AssertThat(testing, minutes, is.EqualTo(test.expected).Reasonf("five minute row %d incorrect", test.minutes))
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
		hours, err := converter.ConvertSingleHours(test.hours)
		then.AssertThat(testing, err, is.Nil())
		then.AssertThat(testing, hours, is.
			EqualTo(test.expected).
			Reasonf("single hours row %d incorrect", test.hours),
		)
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
		hours, err := converter.ConvertFiveHours(test.hours)
		then.AssertThat(testing, err, is.Nil())
		then.AssertThat(testing, hours, is.EqualTo(test.expected).Reasonf("five hours row %d incorrect", test.hours))
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
		seconds, err := converter.ConvertSecondsRow(test.seconds)
		then.AssertThat(testing, err, is.Nil())
		then.AssertThat(testing, seconds, is.EqualTo(test.expected).
			Reasonf("seconds row %d incorrect", test.seconds))
	}
}
