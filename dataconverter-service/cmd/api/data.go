package main

import (
	"math"
	"strconv"
)

type HumanReadableData struct {
	FlowRate                      float32 `json:"flowRate"`
	EnergyFlowRate                float32 `json:"energyFlowRate"`
	Velocity                      float32 `json:"Velocity"`
	FluidSoundSpeed               float32 `json:"fluidSoundSpeed"`
	PositiveAccumulator           int32   `json:"positiveAccumulator"`
	PositiveDecimalFraction       float32 `json:"positiveDecimalFraction"`
	NegativeAccumulator           int32   `json:"negativeAccumulator"`
	NegativeDecimalFraction       float32 `json:"negativeDecimalFraction"`
	PositiveEnergyAccumulator     int32   `json:"positiveEnergyAccumulator"`
	PositiveEnergyDecimalFraction float32 `json:"positiveEnergyDecimalFraction"`
	NegativeEnergyAccumulator     int32   `json:"negativeEnergyAccumulator"`
	NegativeEnergyDecimalFraction float32 `json:"negativeEnergyDecimalFraction"`
	NetAccumulator                int32   `json:"netAccumulator"`
	NetDecimalFraction            float32 `json:"netDecimalFraction"`
	NetEnergyAccumulator          int32   `json:"netEnergyAccumulator"`
	NetEnergyDecimalFraction      float32 `json:"netEnergyDecimalFraction"`
	Temperature1                  float32 `json:"Temperature1"`
	Temperature2                  float32 `json:"Temperature2"`
	AnalogInputA13                float32 `json:"analogInputA13"`
	AnalogInputA14                float32 `json:"analogInputA14"`
	AnalogInputA15                float32 `json:"analogInputA15"`
	CurrentInputA13               float32 `json:"currentInputA13"`
	CurrentInputA13_2             float32 `json:"currentInputA13_2"`
	CurrentInputA13_3             float32 `json:"currentInputA13_3"`
	SystemPassword                string  `json:"systemPassword"`
	PasswordForHardware           string  `json:"passwordForHardware"`
	Calendar                      string  `json:"calendar"`
	DayHourAutoSave               string  `json:"dayHourAutoSave"`
	KeyToInput                    int16   `json:"keyToInput"`
	GoToWindow                    int16   `json:"goToWindow"`
	LCDLights                     int16   `json:"lCDLights"`
	TimesForBeeper                int16   `json:"timesForBeeper"`
	PulsesLeftForOCT              int16   `json:"pulsesLeftForOCT"`
	ErrorCode                     int8    `json:"errorCode"`
	PT100ResistanceInlet          float32 `json:"pT100ResistanceInlet"`
	PT100ResistanceOutlet         float32 `json:"pT100ResistanceOutlet"`
	TotalTravelTime               float32 `json:"totalTravelTime"`
	DeltaTravelTime               float32 `json:"deltaTravelTime"`
	UpstreamTravelTime            float32 `json:"upstreamTravelTime"`
	DownStreamTravelTime          float32 `json:"downStreamTravelTime"`
	OutputCurrent                 float32 `json:"outputCurrent"`
	WorkingStep                   int8    `json:"workingStep"`
	SignalQuality                 int8    `json:"signalQuality"`
	UpstreamStrength              int16   `json:"upstreamStrength"`
	DownStreamStrength            int16   `json:"downStreamStrength"`
	Language                      int16   `json:"language"`
	RateOfTravelTime              float32 `json:"rateOfTravelTime"`
	ReynoldsNumber                float32 `json:"reynoldsNumber"`
}

func convertSMHDMY(regA, regB, regC int) string {
	var date string = ""
	var time string = ""

	time += strconv.Itoa((regB >> 4) & 0xf)   //HOUR fst BCD
	time += strconv.Itoa(regB&0xf) + ":"      //HOUR snd BCD
	time += strconv.Itoa((regA >> 12) & 0xf)  //MINUTE fst BCD
	time += strconv.Itoa((regA>>8)&0xf) + ":" //MINUTE snd BCD
	time += strconv.Itoa((regA >> 4) & 0xf)   //SECOND fst BCD
	time += strconv.Itoa(regA & 0xf)          //SECOND snd BCD

	date += strconv.Itoa((regB >> 12) & 0xf)  //DAY fst BCD
	date += strconv.Itoa((regB>>8)&0xf) + "." //DAY snd BCD
	date += strconv.Itoa((regC >> 4) & 0xf)   //MONTH fst BCD
	date += strconv.Itoa(regC&0xf) + "."      // MONTH snd BCD
	date += strconv.Itoa((regC >> 12) & 0xf)  // YEAR fst BCD
	date += strconv.Itoa((regC >> 8) & 0xf)   // YEAR snd BCD
	return time + " " + date
}

func convertLong(regA, regB int) int32 {
	return int32(regB<<16 + regA)
}

func convertReal4(regA, regB int) float32 {
	return math.Float32frombits(uint32(regB<<16 + regA))
}

func convertSysPas(regA, regB int) string {
	var res string = ""
	res += strconv.Itoa((regA >> 12) & 0xf)
	res += strconv.Itoa((regA >> 8) & 0xf)
	res += strconv.Itoa((regA >> 4) & 0xf)
	res += strconv.Itoa(regA & 0xf)

	res += strconv.Itoa((regB >> 12) & 0xf)
	res += strconv.Itoa((regB >> 8) & 0xf)
	res += strconv.Itoa((regB >> 4) & 0xf)
	res += strconv.Itoa(regB & 0xf)

	return res
}

func convertHWPas(reg int) string {
	var res string = ""
	res += strconv.Itoa((reg >> 12) & 0xf)
	res += strconv.Itoa((reg >> 8) & 0xf)
	res += strconv.Itoa((reg >> 4) & 0xf)
	res += strconv.Itoa(reg & 0xf)
	return res
}

func converDTAutosave(reg int) string {
	var res string = ""
	res += strconv.Itoa((reg >> 12) & 0xf)
	res += strconv.Itoa((reg >> 8) & 0xf)
	res += strconv.Itoa((reg >> 4) & 0xf)
	res += strconv.Itoa(reg & 0xf)
	return res
}
