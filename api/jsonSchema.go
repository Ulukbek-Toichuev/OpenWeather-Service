package api

type Data struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		LocalTime string `json:"localtime"`
	}
	Current struct {
		Last_Update string  `json:"last_updated"`
		CurrentTemp float32 `json:"temp_c"`
		Condition   struct {
			Text string `json:"text"`
		}
		UV         float64 `json:"uv"`
		AirQuality struct {
			Co             float32 `json:"co"`
			No2            float32 `json:"no2"`
			O3             float32 `json:"o3"`
			So2            float32 `json:"so2"`
			Pm2_5          float32 `json:"pm2_5"`
			Pm10           float32 `json:"pm10"`
			US_epa_index   int     `json:"us-epa-index"`
			GB_defra_index int     `json:"gb-defra-index"`
		} `json:"air_quality"`
	}
}
