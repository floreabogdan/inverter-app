package dummy

import (
	"github.com/floreabogdan/inverter-app/inverter/models"
)

type Dummy struct {
	config models.Config
}

func New(config models.Config) *Dummy {
	return &Dummy{
		config: config,
	}
}

func (d Dummy) SetConfig(config models.Config) {
	d.config = config
}

func (d Dummy) GetConfig() models.Config {
	return d.config
}

func (d Dummy) GetData() models.Data {
	return models.Data{
		Pv: models.PvInfo{
			Current: 0,
			Voltage: 0,
		},
		Bms: models.BmsInfo{
			ChargeCurrent:    0,
			ChargeVoltage:    0,
			DischargeCurrent: 0,
			DischargeVoltage: 0,
			CurrentVoltage:   0,
			CapacityTotal:    0,
			CapacityLeft:     0,
		},
		Grid: models.AcInfo{
			Current:   0,
			Voltage:   0,
			Frequency: 0,
		},
		Output: models.OutputInfo{
			AcInfo: models.AcInfo{
				Current:   0,
				Voltage:   0,
				Frequency: 0,
			},
			Status:  false,
			Percent: 0,
		},
		GeneralInfo: models.GeneralInfo{
			InverterMode:            "",
			OutputPriority:          "",
			ChargePriority:          "",
			SccChargeOn:             false,
			SccVoltage:              0,
			BusVoltage:              0,
			BatteryDischargeVoltage: 0,
			BatteryChargeVoltage:    0,
			BatteryFloatVoltage:     0,
			BatteryCutoffVoltage:    0,
			PvMaxChargeCurrent:      0,
			GridMaxChargeCurrent:    0,
			HeatsinkTemperature:     0,
		},
	}
}

func (d Dummy) RunCommand(command string) string {
	return "OK"
}
