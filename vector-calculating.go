package main

import "math"

func CalculateFunction(t int, records []KeyboardHandwritingData) int {
	result := 0
	for i := 0; i < len(records); i++ {
		if records[i].KeyDownTime <= t && t <= records[i].KeyUpTime {
			result += Amplitude
		}
	}
	return result
}

func CalculateHaar(r, m, t int) float64 {
	rFloat, mFloat, tFloat := float64(r), float64(m), float64(t)/1000
	if tFloat >= 1.0 {
		tFloat = 0.999
	}
	result := 0.0

	if ((mFloat-1)/math.Pow(2, rFloat)) <= tFloat && tFloat < ((mFloat-0.5)/math.Pow(2, rFloat)) {
		result = math.Pow(2, rFloat/2)
	} else if ((mFloat-0.5)/math.Pow(2, rFloat)) <= tFloat && tFloat <= (mFloat/math.Pow(2, rFloat)) {
		result = -math.Pow(2, rFloat/2)
	}

	return result
}

func CalculateVector(records []KeyboardHandwritingData) []float64 {

	N := len(records)
	vector := make([]float64, N)
	functionResult := 0.0
	haarResult := 0.0
	countVariable := float64(1 / N)
	sum := 0.0

	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			functionResult = float64(CalculateFunction(records[k].KeyDownTime, records))
			haarResult = CalculateHaar(k, i, records[k].KeyUpTime-records[k].KeyDownTime)
			sum += functionResult * haarResult
		}
		vector[i] = sum * countVariable
		sum = 0.0
	}

	return vector
}

func CalculateVectorWithInformation(records []KeyboardHandwritingData) (information VectorCalculatingResultData, vector []float64) {
	N := len(records)

	vector = make([]float64, N)
	information = VectorCalculatingResultData{
		Amplitude:           Amplitude,
		PressCount:          N,
		FunctionsResult:     make([]float64, N),
		HoldTimeCalculating: make([]SingleHoldCalculatingInformation, N),
	}

	functionResult := 0.0
	haarResult := 0.0
	countVariable := 1 / float64(N)
	sum := 0.0

	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			functionResult = float64(CalculateFunction(records[k].KeyDownTime, records))
			haarResult = CalculateHaar(k, i, records[k].KeyUpTime-records[k].KeyDownTime)
			sum += functionResult * haarResult

			if i == 0 {
				information.FunctionsResult[k] = functionResult
				information.HoldTimeCalculating[k] = SingleHoldCalculatingInformation{
					HarResult: make([]float64, N),
					HoldTime:  make([]int, N),
				}
			}
			information.HoldTimeCalculating[i].HoldTime[k] = records[k].KeyUpTime - records[k].KeyDownTime
			information.HoldTimeCalculating[i].HarResult[k] = haarResult
		}
		vector[i] = sum * countVariable
		sum = 0.0
	}

	return information, vector
}
