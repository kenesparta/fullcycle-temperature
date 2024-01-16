package entity

type LocationRepo interface {
	Get(cep string) Location
}

type TemperatureRepo interface {
	Get(location string) Temperature
}
