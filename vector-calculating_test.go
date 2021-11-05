package main

import (
	"math"
	"testing"
)

func TestVectorCalculating(t *testing.T) {
	testData := []KeyboardHandwritingData{
		{KeyDownTime: 4, KeyUpTime: 15, Key: "a", Code: "a"},
		{KeyDownTime: 14, KeyUpTime: 18, Key: "a", Code: "a"},
		{KeyDownTime: 15, KeyUpTime: 25, Key: "a", Code: "a"},
	}

	zeroAmplitudeValue := CalculateFunction(1, testData)
	if zeroAmplitudeValue != 0 {
		t.Error("Expected value is 0, got", zeroAmplitudeValue)
	}
	oneAmplitudeValue := CalculateFunction(5, testData)
	if oneAmplitudeValue != 5 {
		t.Error("Expected value is 5, got", oneAmplitudeValue)
	}
	twoAmplitudeValue := CalculateFunction(14, testData)
	if twoAmplitudeValue != 10 {
		t.Error("Expected value is 10, got", twoAmplitudeValue)
	}
	threeAmplitudeValue := CalculateFunction(15, testData)
	if threeAmplitudeValue != 15 {
		t.Error("Expected value is 15, got", threeAmplitudeValue)
	}
}

func TestCalculateHaar(t *testing.T) {
	positiveValue := CalculateHaar(2, 3, 600)
	if math.Abs(positiveValue-math.Pow(2.0, 2.0/2)) > 1e-9 {
		t.Error("Positive value is not expected, got", positiveValue)
	}
	nonPositiveValue := CalculateHaar(2, 3, 700)
	if math.Abs(nonPositiveValue+math.Pow(2.0, 2.0/2)) > 1e-9 {
		t.Error("Non positive value is not expected, got", nonPositiveValue)
	}
	zeroValue := CalculateHaar(2, 3, 800)
	if math.Abs(zeroValue) > 1e-9 {
		t.Error("Zero value is not expected, got", zeroValue)
	}
}
