package converter

import (
	"bytes"
	"errors"
)

const alternateRedLampModulo = 3
const fiveMinuteMarker = 5

const maxHours = 23
const maxMins = 59
const maxSeconds = 59

const fiveHoursTotalLamps = 4
const singleRowTotalLamps = 4
const fiveMinuteTotalLamps = 11

const yellowLamp = "Y"
const redLamp = "R"
const offLamp = "O"
const emptyLamps = ""

const invalidArgumentMessage = "invalid argument"

func ConvertSingleMinutes(minutes int) (string, error) {
	if invalidValue(minutes, maxMins) {
		return emptyLamps, errors.New(invalidArgumentMessage)
	}
	yellowLamps := lampsForColour(minutes, yellowLamp)
	return yellowLamps + offLamps(yellowLamps, singleRowTotalLamps), nil
}


func ConvertFiveMinutes(minutes int) (string, error) {
	if invalidValue(minutes, maxMins) {
		return emptyLamps, errors.New(invalidArgumentMessage)
	}
	lamps := alternateColourLamps(minutes, yellowLamp)
	return lamps + offLamps(lamps, fiveMinuteTotalLamps), nil
}

func ConvertSingleHours(hour int) (string, error) {
	if invalidValue(hour, maxHours) {
		return emptyLamps, errors.New(invalidArgumentMessage)
	}
	redLamps := lampsForColour(hour, redLamp)
	return redLamps + offLamps(redLamps, singleRowTotalLamps), nil
}

func ConvertFiveHours(hours int) (string, error) {
	if invalidValue(hours, maxHours) {
		return emptyLamps, errors.New(invalidArgumentMessage)
	}
	lamps := alternateColourLamps(hours, redLamp)
	return lamps + offLamps(lamps, fiveHoursTotalLamps), nil
}

func ConvertSecondsRow(seconds int) (string, error) {
	if invalidValue(seconds, maxSeconds) {
		return emptyLamps, errors.New(invalidArgumentMessage)
	}
	var lamp = offLamp
	if seconds%2 == 0 {
		lamp = yellowLamp
	}
	return lamp, nil
}

func lampsForColour(value int, lampColour string) string {
	numberOfOnLamps := value % fiveMinuteMarker
	var yellowLampBuffer bytes.Buffer
	for i := 0; i < numberOfOnLamps; i++ {
		yellowLampBuffer.WriteString(lampColour)
	}
	return yellowLampBuffer.String()
}

func alternateColourLamps(value int, lampColour string) string {
	numberOfOnLamps := value / fiveMinuteMarker
	var yellowLampBuffer bytes.Buffer
	for i := 0; i < numberOfOnLamps; i++ {
		if isRedLamp(i) {
			yellowLampBuffer.WriteString(redLamp)
		} else {
			yellowLampBuffer.WriteString(lampColour)
		}
	}
	return yellowLampBuffer.String()
}
func isRedLamp(lampNumber int) bool {
	return (lampNumber+1)%alternateRedLampModulo == 0
}

func offLamps(currentLamps string, totalLamps int) string {
	numberOfLampsOn := len(currentLamps)
	var offLampBuffer bytes.Buffer
	if numberOfLampsOn < (totalLamps + 1) {
		for i := numberOfLampsOn; i < totalLamps; i++ {
			offLampBuffer.WriteString(offLamp)
		}
	}
	return offLampBuffer.String()
}

func invalidValue(minutes int, maxValue int) bool {
	return minutes < 0 || minutes > maxValue
}
