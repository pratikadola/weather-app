package weather

type WeatherResponse struct {
	Weather []WeatherDetails `json:"weather"`
	Main    MainParameters   `json:"main"`
	Wind    Wind             `json:"wind"`
	Rain    *Rain            `json:"rain,omitempty"`
	Snow    *Snow            `json:"snow,omitempty"`
}

type WeatherDetails struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type MainParameters struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json"feels_like"`
	Humidity  float32 `json:"humidity"`
}

type Wind struct {
	Speed float32 `json:"speed"`
}

type Rain struct {
	H1 string `json:"1h"`
}

type Snow struct {
	H1 string `json:"1h"`
}
