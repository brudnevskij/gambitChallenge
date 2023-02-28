package main

import "math"

type HumanReadableData struct {
	FlowRate                      float32
	EnergyFlowRate                float32
	Velocity                      float32
	FluidSoundSpeed               float32
	PositiveAccumulator           int32
	PositiveDecimalFraction       float32
	NegativeAccumulator           int32
	NegativeDecimalFraction       float32
	PositiveEnergyAccumulator     int32
	PositiveEnergyDecimalFraction float32
	NegativeEnergyAccumulator     int32
	NegativeEnergyDecimalFraction float32
	NetAccumulator                int32
	NetDecimalFraction            float32
	NetEnergyAccumulator          int32
	NetEnergyDecimalFraction      float32
	Temperature1                  float32
	Temperature2                  float32
	AnalogInputA13                float32
	AnalogInputA14                float32
	AnalogInputA15                float32
	CurrentInputA13               float32
	CurrentInputA13_2             float32
	CurrentInputA13_3             float32
	SystemPassword                string
	PasswordForHardware           string
	Calendar                      string
	DayHourAutoSave               string
	KeyToInput                    int16
	GoToWindow                    int16
	LCDLights                     int16
	TimesForBeeper                int16
	PulsesLeftForOCT              int16
	ErrorCode                     int8
	PT100ResistanceInlet          float32
	PT100ResistanceOutlet         float32
	TotalTravelTime               float32
	DeltaTravelTime               float32
	UpstreamTravelTime            float32
	DownStreamTravelTime          float32
	OutputCurrent                 float32
	WorkingStep                   int8
	SignalQuality                 int8
	UpstreamStrength              int16
	DownStreamStrength            int16
	Language                      int16
	RateOfTravelTime              float32
	ReynoldsNumber                float32
}

func convertLong(regA, regB int) int32 {
	// var res uint32
	// res = uint32(regB << 16)
	// res += int64(regA)

	return int32(regB<<16 + regA)
}

func convertReal4(regA, regB int) float32 {
	return math.Float32frombits(uint32(regB<<16 + regA))
}
