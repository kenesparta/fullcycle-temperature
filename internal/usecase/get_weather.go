package usecase

import (
	"context"
	"github.com/kenesparta/fullcycle-temperature/internal/dto"
)

type GetWeather struct {
}

func NewGetWeather() *GetWeather {
	return &GetWeather{}
}

func (gw *GetWeather) Execute(
	ctx context.Context,
	input dto.LocationInput,
) (dto.TemperatureOutput, error) {
	// "localidade": "Manaus",
	// get localidade from CEP API
	// put "localidade" in the q:
	// https://api.weatherapi.com/v1/current.json?key=&q=Manaus&aqi=yes
	return dto.TemperatureOutput{}, nil
}
