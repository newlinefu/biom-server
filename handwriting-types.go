package main

const Amplitude int = 5

type KeyboardHandwritingData struct {
	KeyUpTime   int    `json:"keyUpTime"`
	KeyDownTime int    `json:"keyDownTime"`
	Key         string `json:"key"`
	Code        string `json:"code"`
}

type SingleHoldCalculatingInformation struct {
	HarResult []float64 `json:"har_result"`
	HoldTime  []int     `json:"hold_time"`
}

type VectorCalculatingResultData struct {
	Amplitude           int                                `json:"amplitude"`
	PressCount          int                                `json:"press_count"`
	FunctionsResult     []float64                          `json:"functions_result"`
	HoldTimeCalculating []SingleHoldCalculatingInformation `json:"hold_time_calculating"`
}

type InformationVectorResponse struct {
	Vector      []float64                   `json:"vector"`
	Information VectorCalculatingResultData `json:"information"`
}
