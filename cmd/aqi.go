package cmd

import (
	jsonSchema "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
)

var airQuality string

func GetAQI(p *jsonSchema.Data) string {
	res := p.Current.AirQuality.No2 + p.Current.AirQuality.O3 + p.Current.AirQuality.So2 + p.Current.AirQuality.Pm2_5 + p.Current.AirQuality.Pm10

	if 5 < res && res < 220 {
		airQuality = "Air pollution: Good"
	} else if 220 < res && res < 450 {
		airQuality = "Air pollution: Fair"
	} else if 450 < res && res < 675 {
		airQuality = "Air pollution: Moderate"
	} else if 6750 < res && res < 1120 {
		airQuality = "Air pollution: Poor"
	} else if 1120 < res && res < 1695 {
		airQuality = "Air pollution: Very poor"
	} else if 1695 < res && res < 5050 {
		airQuality = "Air pollution: Extremely poor"
	} else if 1695 < res && res < 5050 {
		airQuality = "Air pollution: Extremely poor"
	}

	return airQuality
}
