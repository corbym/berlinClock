package berlinclock_test

import (
	"testing"
	"github.com/corbym/berlinclock"
	. "github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gogiven"
	"os"
	"github.com/corbym/gogiven/base"
	"github.com/corbym/gogiven/testdata"
	"github.com/corbym/htmlspec"
)

func TestMain(testmain *testing.M) {
	gogiven.Generator = htmlspec.NewTestOutputGenerator()
	runOutput := testmain.Run()
	gogiven.GenerateTestOutput()
	os.Exit(runOutput)
}

type output struct {
	berlinClockOutput string
	errorOutput       error
}

func TestGivenAClockWhenTimeIsEnteredThenCorrectClock(t *testing.T) {
	var clockParams = []struct {
		time     string // input
		expected string // expected result
	}{
		{time: "00:00:00", expected: "YOOOOOOOOOOOOOOOOOOOOOOO"},
		{time: "23:59:59", expected: "ORRRRRRROYYRYYRYYRYYYYYY"},
		{time: "16:50:06", expected: "YRRROROOOYYRYYRYYRYOOOOO"},
		{time: "11:37:01", expected: "ORROOROOOYYRYYRYOOOOYYOO"},
	}
	for _, test := range clockParams {
		t.Run(test.time, func(weAreTesting *testing.T) {
			outputs := &output{}
			gogiven.Given(weAreTesting, clockParametersUnder(test)).

				When(weAskTheClockForIts(test, outputs)).

				Then(func(theTest base.TestingT, actual testdata.CapturedIO, givens testdata.InterestingGivens) {

				AssertThat(theTest, outputs.errorOutput, is.Nil())
				AssertThat(theTest, outputs.berlinClockOutput, is.EqualTo(test.expected).Reasonf("time incorrect for %s", test.time))
			})
		})
	}
}

type testData struct {
	time     string
	expected string
}

func clockParametersUnder(testData testData) (givens base.GivenData) {
	return func(givens testdata.InterestingGivens) {
		givens["time"] = testData.time
		givens["expected"] = testData.expected
	}
}
func weAskTheClockForIts(test testData, out *output) (givens func(capturedIO testdata.CapturedIO, givens testdata.InterestingGivens)) {
	givens = func(capturedIO testdata.CapturedIO, givens testdata.InterestingGivens) {
		out.berlinClockOutput, out.errorOutput = berlinclock.Clock(test.time)
		capturedIO["output"] = out.berlinClockOutput
		capturedIO["error"] = out.errorOutput
	}
	return

}

func TestGivenAClockWhenInvalidTimeIsEnteredThenClockError(t *testing.T) {
	var clockParams = []struct {
		time     string // input
		expected string // expected result
	}{
		{time: "a:b:c", expected: "invalid argument, a:b:c is invalid time format"},
		{time: "99:23:21", expected: "invalid argument"},
		{time: "16:99:06", expected: "invalid argument"},
		{time: "11:37:99", expected: "invalid argument"},
	}
	for _, test := range clockParams {
		t.Run(test.time, func(weAreTesting *testing.T) {
			out := &output{}
			gogiven.Given(weAreTesting, clockParametersUnder(test)).

				When(weAskTheClockForIts(test, out)).

				Then(func(testingT base.TestingT, actual testdata.CapturedIO, givens testdata.InterestingGivens) {
				AssertThat(weAreTesting, out.berlinClockOutput, is.Empty().Reasonf("clock was not empty for given time %s", test.time))
				AssertThat(weAreTesting, out.errorOutput.Error(), is.EqualTo(test.expected).Reasonf("incorrect error for given time %s", test.time))
			})

		})
	}
}
