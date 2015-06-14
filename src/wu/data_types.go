package wu

type WeatherData struct {
	CurrentObservation struct {
		UV              string `json:"UV"`
		DewpointC       float64    `json:"dewpoint_c"`
		DewpointF       float64    `json:"dewpoint_f"`
		DewpointString  string `json:"dewpoint_string"`
		DisplayLocation struct {
			City           string `json:"city"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Elevation      string `json:"elevation"`
			Full           string `json:"full"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Magic          string `json:"magic"`
			State          string `json:"state"`
			StateName      string `json:"state_name"`
			Wmo            string `json:"wmo"`
			Zip            string `json:"zip"`
		} `json:"display_location"`
		Estimated       struct{} `json:"estimated"`
		FeelslikeC      string   `json:"feelslike_c"`
		FeelslikeF      string   `json:"feelslike_f"`
		FeelslikeString string   `json:"feelslike_string"`
		ForecastURL     string   `json:"forecast_url"`
		HeatIndexC      string   `json:"heat_index_c"`
		HeatIndexF      string   `json:"heat_index_f"`
		HeatIndexString string   `json:"heat_index_string"`
		HistoryURL      string   `json:"history_url"`
		Icon            string   `json:"icon"`
		IconURL         string   `json:"icon_url"`
		Image           struct {
			Link  string `json:"link"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"image"`
		LocalEpoch          string `json:"local_epoch"`
		LocalTimeRfc822     string `json:"local_time_rfc822"`
		LocalTzLong         string `json:"local_tz_long"`
		LocalTzOffset       string `json:"local_tz_offset"`
		LocalTzShort        string `json:"local_tz_short"`
		Nowcast             string `json:"nowcast"`
		ObURL               string `json:"ob_url"`
		ObservationEpoch    string `json:"observation_epoch"`
		ObservationLocation struct {
			City           string `json:"city"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Elevation      string `json:"elevation"`
			Full           string `json:"full"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			State          string `json:"state"`
		} `json:"observation_location"`
		ObservationTime       string  `json:"observation_time"`
		ObservationTimeRfc822 string  `json:"observation_time_rfc822"`
		Precip1hrIn           string  `json:"precip_1hr_in"`
		Precip1hrMetric       string  `json:"precip_1hr_metric"`
		Precip1hrString       string  `json:"precip_1hr_string"`
		PrecipTodayIn         string  `json:"precip_today_in"`
		PrecipTodayMetric     string  `json:"precip_today_metric"`
		PrecipTodayString     string  `json:"precip_today_string"`
		PressureIn            string  `json:"pressure_in"`
		PressureMb            string  `json:"pressure_mb"`
		PressureTrend         string  `json:"pressure_trend"`
		RelativeHumidity      string  `json:"relative_humidity"`
		Solarradiation        string  `json:"solarradiation"`
		StationID             string  `json:"station_id"`
		TempC                 float64 `json:"temp_c"`
		TempF                 float64 `json:"temp_f"`
		TemperatureString     string  `json:"temperature_string"`
		VisibilityKm          string  `json:"visibility_km"`
		VisibilityMi          string  `json:"visibility_mi"`
		Weather               string  `json:"weather"`
		WindDegrees           float64     `json:"wind_degrees"`
		WindDir               string  `json:"wind_dir"`
		WindGustKph           string  `json:"wind_gust_kph"`
		WindGustMph           string  `json:"wind_gust_mph"`
		WindKph               float64 `json:"wind_kph"`
		WindMph               float64     `json:"wind_mph"`
		WindString            string  `json:"wind_string"`
		WindchillC            string  `json:"windchill_c"`
		WindchillF            string  `json:"windchill_f"`
		WindchillString       string  `json:"windchill_string"`
	} `json:"current_observation"`
	Forecast struct {
		Simpleforecast struct {
			Forecastday []struct {
				Avehumidity int `json:"avehumidity"`
				Avewind     struct {
					Degrees int    `json:"degrees"`
					Dir     string `json:"dir"`
					Kph     float64    `json:"kph"`
					Mph     float64    `json:"mph"`
				} `json:"avewind"`
				Conditions string `json:"conditions"`
				Date       struct {
					Ampm           string `json:"ampm"`
					Day            int    `json:"day"`
					Epoch          string `json:"epoch"`
					Hour           int    `json:"hour"`
					Isdst          string `json:"isdst"`
					Min            string `json:"min"`
					Month          int    `json:"month"`
					Monthname      string `json:"monthname"`
					MonthnameShort string `json:"monthname_short"`
					Pretty         string `json:"pretty"`
					Sec            int    `json:"sec"`
					TzLong         string `json:"tz_long"`
					TzShort        string `json:"tz_short"`
					Weekday        string `json:"weekday"`
					WeekdayShort   string `json:"weekday_short"`
					Yday           int    `json:"yday"`
					Year           int    `json:"year"`
				} `json:"date"`
				High struct {
					Celsius    string `json:"celsius"`
					Fahrenheit string `json:"fahrenheit"`
				} `json:"high"`
				Icon    string `json:"icon"`
				IconURL string `json:"icon_url"`
				Low     struct {
					Celsius    string `json:"celsius"`
					Fahrenheit string `json:"fahrenheit"`
				} `json:"low"`
				Maxhumidity int `json:"maxhumidity"`
				Maxwind     struct {
					Degrees int    `json:"degrees"`
					Dir     string `json:"dir"`
					Kph     float64    `json:"kph"`
					Mph     float64    `json:"mph"`
				} `json:"maxwind"`
				Minhumidity int `json:"minhumidity"`
				Period      int `json:"period"`
				Pop         int `json:"pop"`
				QpfAllday   struct {
					In float64 `json:"in"`
					Mm float64 `json:"mm"`
				} `json:"qpf_allday"`
				QpfDay struct {
					In float64 `json:"in"`
					Mm float64 `json:"mm"`
				} `json:"qpf_day"`
				QpfNight struct {
					In float64 `json:"in"`
					Mm float64 `json:"mm"`
				} `json:"qpf_night"`
				Skyicon    string `json:"skyicon"`
				SnowAllday struct {
					Cm float64 `json:"cm"`
					In float64 `json:"in"`
				} `json:"snow_allday"`
				SnowDay struct {
					Cm float64 `json:"cm"`
					In float64 `json:"in"`
				} `json:"snow_day"`
				SnowNight struct {
					Cm float64 `json:"cm"`
					In float64 `json:"in"`
				} `json:"snow_night"`
			} `json:"forecastday"`
		} `json:"simpleforecast"`
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
		TermsofService string `json:"termsofService"`
		Version        string `json:"version"`
	} `json:"response"`
}
