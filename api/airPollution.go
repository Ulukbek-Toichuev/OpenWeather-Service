package api

// Структура для данных о загрязнении воздуха в JSON формате
// Structure for air pollution data in JSON format
type AirPollution struct {
	Coord struct {
		CoordLon float64 `json:"lon"`
		CoordLat float64 `json:"lat"`
	}
	List []struct {
		Main struct {
			Aqi int `json:"aqi"`
		}
		Components struct {
			Co    float64 `json:"co"`
			No    float64 `json:"no"`
			No2   float64 `json:"no2"`
			O3    float64 `json:"o3"`
			So2   float64 `json:"so2"`
			Pm2_5 float64 `json:"pm2_5"`
			Pm10  float64 `json:"pm10"`
			Nh3   float64 `json:"nh3"`
		}
		DT int `json:"dt"`
	}
}
