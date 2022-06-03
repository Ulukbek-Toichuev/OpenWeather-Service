package api

type List struct {
	Main struct {
		Aqi int `json:"aqi"`
	}
	Component struct {
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

type Coor struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type AirPollution struct {
	Coor Coor
	List []List
}
