package controller

/*type Coord struct {
	CoordLon float64 `json:"lon"`
	CoordLat float64 `json:"lat"`
}*/

type GeoCode struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Weather struct {
	WeatherID          int    `json:"id"`
	WeatherMain        string `json:"main"`
	WeatherDescription string `json:"description"`
	WeatherIcon        string `json:"icon"`
}

type OWM struct {
	Coord struct {
		CoordLon float64 `json:"lon"`
		CoordLat float64 `json:"lat"`
	}
	Weather     []Weather
	BaseMainOWM string `json:"base"`
	Main        struct {
		MainTemp      float64 `json:"temp"`
		MainFeelsLike float64 `json:"feels_like"`
		MainPressure  int     `json:"pressure"`
		MainHumidity  int     `json:"humidity"`
		MainTempMin   float64 `json:"temp_min"`
		MainTempMax   float64 `json:"temp_max"`
		MainSeaLevel  float64 `json:"sea_level"`
		MainGrndLevel float64 `json:"grnd_level"`
	}
	VisibilityOWM int `json:"visibility"`
	Wind          struct {
		WindSpeed float64 `json:"speed"`
		WindDeg   int     `json:"deg"`
		WindGust  float64 `json:"gust"`
	}
	Clouds struct {
		CloudsALL int `json:"all"`
	}
	Rain struct {
		Rain1h float64 `json:"1h"`
		Rain3h float64 `json:"3h"`
	}
	Snow struct {
		Snow1h float64 `json:"1h"`
		Snow3h float64 `json:"3h"`
	}
	DtOWM int `json:"dt"`
	Sys   struct {
		Type    int     `json:"type"`
		SysID   int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	}
	TimeZone int    `json:"timezone"`
	IDOWM    int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
