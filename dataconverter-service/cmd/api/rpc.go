package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type RPCServer struct{}

func (r *RPCServer) GetDataRPC(args string, resp *HumanReadableData) error {
	//getting data from API
	request, err := http.NewRequest(http.MethodGet, app.dataApi, nil)
	if err != nil {
		app.errorLog.Print("Could not create a request to API")
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorLog.Print("Could not do a request to API")
		return err
	}
	defer response.Body.Close()

	//Parsing data
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		app.errorLog.Print("Could not read a body of a request")
		return err
	}

	data := strings.Split(string(body), "\n")
	data = data[1 : len(data)-1]
	var parsedData [100]int

	for i, x := range data {
		data[i] = strings.Split(x, ":")[1]
		buffer, _ := strconv.Atoi(strings.Split(x, ":")[1])
		parsedData[i] = buffer
	}

	//populating struct
	*resp = HumanReadableData{
		FlowRate:                      convertReal4(parsedData[0], parsedData[1]),
		EnergyFlowRate:                convertReal4(parsedData[2], parsedData[3]),
		Velocity:                      convertReal4(parsedData[4], parsedData[5]),
		FluidSoundSpeed:               convertReal4(parsedData[6], parsedData[7]),
		PositiveAccumulator:           convertLong(parsedData[8], parsedData[9]),
		PositiveDecimalFraction:       convertReal4(parsedData[10], parsedData[11]),
		NegativeAccumulator:           convertLong(parsedData[12], parsedData[13]),
		NegativeDecimalFraction:       convertReal4(parsedData[14], parsedData[15]),
		PositiveEnergyAccumulator:     convertLong(parsedData[16], parsedData[17]),
		PositiveEnergyDecimalFraction: convertReal4(parsedData[18], parsedData[19]),
		NegativeEnergyAccumulator:     convertLong(parsedData[20], parsedData[21]),
		NegativeEnergyDecimalFraction: convertReal4(parsedData[22], parsedData[23]),
		NetAccumulator:                convertLong(parsedData[24], parsedData[25]),
		NetDecimalFraction:            convertReal4(parsedData[26], parsedData[27]),
		NetEnergyAccumulator:          convertLong(parsedData[28], parsedData[29]),
		NetEnergyDecimalFraction:      convertReal4(parsedData[30], parsedData[31]),
		Temperature1:                  convertReal4(parsedData[32], parsedData[33]),
		Temperature2:                  convertReal4(parsedData[34], parsedData[35]),
		AnalogInputA13:                convertReal4(parsedData[36], parsedData[38]),
		AnalogInputA14:                convertReal4(parsedData[38], parsedData[39]),
		AnalogInputA15:                convertReal4(parsedData[40], parsedData[41]),
		CurrentInputA13:               convertReal4(parsedData[42], parsedData[43]),
		CurrentInputA13_2:             convertReal4(parsedData[44], parsedData[45]),
		CurrentInputA13_3:             convertReal4(parsedData[46], parsedData[47]),
		SystemPassword:                convertSysPas(parsedData[48], parsedData[49]),
		PasswordForHardware:           convertHWPas(parsedData[50]),
		Calendar:                      convertSMHDMY(parsedData[52], parsedData[53], parsedData[54]),
		DayHourAutoSave:               converDTAutosave(parsedData[55]),
		KeyToInput:                    int16(parsedData[58]),
		GoToWindow:                    int16(parsedData[59]),
		LCDLights:                     int16(parsedData[60]),
		TimesForBeeper:                int16(parsedData[61]),
		PulsesLeftForOCT:              int16(parsedData[61]),
		ErrorCode:                     int8(parsedData[71]),
		PT100ResistanceInlet:          convertReal4(parsedData[76], parsedData[77]),
		PT100ResistanceOutlet:         convertReal4(parsedData[78], parsedData[79]),
		TotalTravelTime:               convertReal4(parsedData[80], parsedData[81]),
		DeltaTravelTime:               convertReal4(parsedData[82], parsedData[83]),
		UpstreamTravelTime:            convertReal4(parsedData[84], parsedData[85]),
		DownStreamTravelTime:          convertReal4(parsedData[86], parsedData[87]),
		OutputCurrent:                 convertReal4(parsedData[88], parsedData[89]),
		WorkingStep:                   int8(parsedData[91] >> 8),
		SignalQuality:                 int8(parsedData[91]),
		UpstreamStrength:              int16(parsedData[92]),
		DownStreamStrength:            int16(parsedData[93]),
		Language:                      int16(parsedData[95]),
		RateOfTravelTime:              convertReal4(parsedData[96], parsedData[97]),
		ReynoldsNumber:                convertReal4(parsedData[98], parsedData[99]),
	}

	return nil
}
