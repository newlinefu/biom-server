package main

const Amplitude int = 5

type KeyboardHandwritingData struct {
	KeyUpTime   int    `json:"keyUpTime"`
	KeyDownTime int    `json:"keyDownTime"`
	Key         string `json:"key"`
	Code        string `json:"code"`
}

type SingleHoldCalculatingInformation struct {
	HarResult []float64 `json:"harResult"`
	HoldTime  []int     `json:"holdTime"`
}

type VectorCalculatingResultData struct {
	Amplitude           int                                `json:"amplitude"`
	PressCount          int                                `json:"pressCount"`
	FunctionsResult     []float64                          `json:"functionsResult"`
	HoldTimeCalculating []SingleHoldCalculatingInformation `json:"holdTimeCalculating"`
}

type InformationVectorResponse struct {
	Vector      []float64                   `json:"vector"`
	Information VectorCalculatingResultData `json:"information"`
}
