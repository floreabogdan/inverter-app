package models

type Config struct {
	Connection string
}

type PvInfo struct {
	Current float32
	Voltage float32
}

type BmsInfo struct {
	ChargeCurrent    float32
	ChargeVoltage    float32
	DischargeCurrent float32
	DischargeVoltage float32

	CurrentVoltage float32

	CapacityTotal float32
	CapacityLeft  float32
}

type AcInfo struct {
	Current   float32
	Voltage   float32
	Frequency float32
}

type OutputInfo struct {
	AcInfo
	Status  bool
	Percent float32
}

type GeneralInfo struct {
	InverterMode string

	OutputPriority string
	ChargePriority string

	SccChargeOn bool
	SccVoltage  float32

	BusVoltage float32

	BatteryDischargeVoltage float32
	BatteryChargeVoltage    float32
	BatteryFloatVoltage     float32
	BatteryCutoffVoltage    float32

	PvMaxChargeCurrent   float32
	GridMaxChargeCurrent float32

	HeatsinkTemperature float32
}

type Data struct {
	Pv   PvInfo
	Bms  BmsInfo
	Grid AcInfo

	Output OutputInfo

	GeneralInfo
}
