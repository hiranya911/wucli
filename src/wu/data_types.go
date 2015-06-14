package wu

// Auto-generated using http://json2struct.mervine.net/
type WeatherData struct {
	CurrentObservation struct {
		DisplayLocation struct {
			Full           string `json:"full"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
		} `json:"display_location"`
		FeelslikeString string   `json:"feelslike_string"`
		TemperatureString     string  `json:"temperature_string"`
		Weather               string  `json:"weather"`
		WindString            string `json:"wind_string"`
		WindchillString       string `json:"windchill_string"`
		PrecipTodayString     string `json:"precip_today_string"`
		ObservationTime       string  `json:"observation_time"`
	} `json:"current_observation"`
	Forecast struct {
		TxtForecast struct {
			Date        string `json:"date"`
			Forecastday []struct {
				Fcttext       string `json:"fcttext"`
				FcttextMetric string `json:"fcttext_metric"`
				Icon          string `json:"icon"`
				IconURL       string `json:"icon_url"`
				Period        int    `json:"period"`
				Pop           string `json:"pop"`
				Title         string `json:"title"`
			} `json:"forecastday"`
		} `json:"txt_forecast"`
	} `json:"forecast"`
	Response struct {
		Error struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"error"`
		Features struct {
			Conditions int `json:"conditions"`
			Forecast   int `json:"forecast"`
		} `json:"features"`
		Results []struct {
			City           string `json:"city"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			CountryName    string `json:"country_name"`
			L              string `json:"l"`
			Name           string `json:"name"`
			State          string `json:"state"`
			Zmw            string `json:"zmw"`
		} `json:"results"`
		TermsofService string `json:"termsofService"`
		Version        string `json:"version"`
	} `json:"response"`
}
