package converter_test

import (
	"testing"
	"github.com/corbym/berlinclock/converter"
	. "github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gogiven"
	"fmt"
	"github.com/corbym/gogiven/base"
	"github.com/corbym/gogiven/testdata"
	"reflect"
	"github.com/corbym/htmlspec"
	"os"
)

func TestMain(testmain *testing.M) {
	gogiven.Generator = htmlspec.NewTestOutputGenerator()
	runOutput := testmain.Run()
	gogiven.GenerateTestOutput()
	os.Exit(runOutput)
}

type output struct {
	text  string
	error error
}

func TestGivenInvalidMinutesThenError(t *testing.T) {
	const invalidArgumentMessage = "invalid argument"
	var functions = []struct {
		converter    func(int) (string, error)
		invalidValue int
		expected     string
	}{
		{converter: converter.ConvertSingleMinutes, invalidValue: -1, expected: invalidArgumentMessage},
		{converter: converter.ConvertSingleMinutes, invalidValue: 60, expected: invalidArgumentMessage},
		{converter: converter.ConvertSingleMinutes, invalidValue: 70, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveMinutes, invalidValue: -1, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveMinutes, invalidValue: -5, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveMinutes, invalidValue: 60, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveMinutes, invalidValue: 70, expected: invalidArgumentMessage},
		{converter: converter.ConvertSingleHours, invalidValue: -1, expected: invalidArgumentMessage},
		{converter: converter.ConvertSingleHours, invalidValue: 24, expected: invalidArgumentMessage},
		{converter: converter.ConvertSingleHours, invalidValue: 30, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveHours, invalidValue: -1, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveHours, invalidValue: 60, expected: invalidArgumentMessage},
		{converter: converter.ConvertFiveHours, invalidValue: 75, expected: invalidArgumentMessage},
		{converter: converter.ConvertSecondsRow, invalidValue: 75, expected: invalidArgumentMessage},
		{converter: converter.ConvertSecondsRow, invalidValue: 60, expected: invalidArgumentMessage},
		{converter: converter.ConvertSecondsRow, invalidValue: -1, expected: invalidArgumentMessage},
	}
	for _, with := range functions {
		t.Run(fmt.Sprintf("%d", with.invalidValue), func(weAreTesting *testing.T) {
			output := &output{}
			gogiven.Given(weAreTesting).
				When(weConvert(with.invalidValue, with.converter, to(output))).
				Then(func(theConverters base.TestingT, actual testdata.CapturedIO, givens testdata.InterestingGivens) {
				AssertThat(theConverters, output.text, is.Empty())
				AssertThat(theConverters, output.error.Error(), is.EqualTo(with.expected).Reasonf("error message for %v incorrect", with))

			})
		})
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
		AssertThat(testing, err, is.Nil())
		AssertThat(testing, minutes, is.EqualTo(test.expected).Reasonf("single minutes row %d incorrect", test.minutes))
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
		AssertThat(testing, err, is.Nil())
		AssertThat(testing, minutes, is.EqualTo(test.expected).Reasonf("five minute row %d incorrect", test.minutes))
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
		AssertThat(testing, err, is.Nil())
		AssertThat(testing, hours, is.
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
		AssertThat(testing, err, is.Nil())
		AssertThat(testing, hours, is.EqualTo(test.expected).Reasonf("five hours row %d incorrect", test.hours))
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
		AssertThat(testing, err, is.Nil())
		AssertThat(testing, seconds, is.EqualTo(test.expected).
			Reasonf("seconds row %d incorrect", test.seconds))
	}
}

func weConvert(value int, with func(int) (string, error), output *output) base.CapturedIOGivenData {
	return func(capturedIO testdata.CapturedIO, givens testdata.InterestingGivens) {
		givens["converter name"] = reflect.TypeOf(with).String()
		givens["invalid value"] = value
		output.text, output.error = with(value)
		capturedIO["converter output"] = output.text
		capturedIO["error"] = output.error.Error()
	}
}

func to(output *output) *output {
	return output
}
