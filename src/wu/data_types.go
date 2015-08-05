package wu

// Auto-generated using http://json2struct.mervine.net/
type WeatherData struct {
	CurrentObservation struct {
		DisplayLocation struct {
			Full      string `json:"full"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"display_location"`
		FeelslikeString   string `json:"feelslike_string"`
		TemperatureString string `json:"temperature_string"`
		Weather           string `json:"weather"`
		WindString        string `json:"wind_string"`
		WindchillString   string `json:"windchill_string"`
		PrecipTodayString string `json:"precip_today_string"`
		ObservationTime   string `json:"observation_time"`
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
	Alerts []struct {
		StormBased struct {
			VertexCount int `json:"Vertex_count"`
			StormInfo   struct {
				MotionDeg   int     `json:"Motion_deg"`
				MotionSpd   int     `json:"Motion_spd"`
				PositionLat float64 `json:"position_lat"`
				PositionLon float64 `json:"position_lon"`
				TimeEpoch   int     `json:"time_epoch"`
			} `json:"stormInfo"`
			Vertices []struct {
				Lat string `json:"lat"`
				Lon string `json:"lon"`
			} `json:"vertices"`
		} `json:"StormBased"`
		ZONES []struct {
			ZONE  string `json:"ZONE"`
			State string `json:"state"`
		} `json:"ZONES"`
		Date         string `json:"date"`
		DateEpoch    string `json:"date_epoch"`
		Description  string `json:"description"`
		Expires      string `json:"expires"`
		ExpiresEpoch string `json:"expires_epoch"`
		Message      string `json:"message"`
		Phenomena    string `json:"phenomena"`
		Significance string `json:"significance"`
		Type         string `json:"type"`
		Attribution                string `json:"attribution"`
		LevelMeteoalarm            string `json:"level_meteoalarm"`
		LevelMeteoalarmDescription string `json:"level_meteoalarm_description"`
		LevelMeteoalarmName        string `json:"level_meteoalarm_name"`
		WtypeMeteoalarm            string `json:"wtype_meteoalarm"`
		WtypeMeteoalarmName        string `json:"wtype_meteoalarm_name"`
	} `json:"alerts"`
	Almanac struct {
		AirportCode string `json:"airport_code"`
		TempHigh    struct {
			Normal struct {
				C string `json:"C"`
				F string `json:"F"`
			} `json:"normal"`
			Record struct {
				C string `json:"C"`
				F string `json:"F"`
			} `json:"record"`
			Recordyear string `json:"recordyear"`
		} `json:"temp_high"`
		TempLow struct {
			Normal struct {
				C string `json:"C"`
				F string `json:"F"`
			} `json:"normal"`
			Record struct {
				C string `json:"C"`
				F string `json:"F"`
			} `json:"record"`
			Recordyear string `json:"recordyear"`
		} `json:"temp_low"`
	} `json:"almanac"`
	MoonPhase struct {
		AgeOfMoon   string `json:"ageOfMoon"`
		CurrentTime struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"current_time"`
		PercentIlluminated string `json:"percentIlluminated"`
		Sunrise            struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"moon_phase"`
	HourlyForecast []struct {
		FCTTIME struct {
			Age                    string `json:"age"`
			Ampm                   string `json:"ampm"`
			Civil                  string `json:"civil"`
			Epoch                  string `json:"epoch"`
			Hour                   string `json:"hour"`
			HourPadded             string `json:"hour_padded"`
			Isdst                  string `json:"isdst"`
			Mday                   string `json:"mday"`
			MdayPadded             string `json:"mday_padded"`
			Min                    string `json:"min"`
			Mon                    string `json:"mon"`
			MonAbbrev              string `json:"mon_abbrev"`
			MonPadded              string `json:"mon_padded"`
			MonthName              string `json:"month_name"`
			MonthNameAbbrev        string `json:"month_name_abbrev"`
			Pretty                 string `json:"pretty"`
			Sec                    string `json:"sec"`
			Tz                     string `json:"tz"`
			WeekdayName            string `json:"weekday_name"`
			WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
			WeekdayNameNight       string `json:"weekday_name_night"`
			WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
			WeekdayNameUnlang      string `json:"weekday_name_unlang"`
			Yday                   string `json:"yday"`
			Year                   string `json:"year"`
		} `json:"FCTTIME"`
		Condition string `json:"condition"`
		Dewpoint  struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"dewpoint"`
		Fctcode   string `json:"fctcode"`
		Feelslike struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"feelslike"`
		Heatindex struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"heatindex"`
		Humidity string `json:"humidity"`
		Icon     string `json:"icon"`
		IconURL  string `json:"icon_url"`
		Mslp     struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"mslp"`
		Pop string `json:"pop"`
		Qpf struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"qpf"`
		Sky  string `json:"sky"`
		Snow struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"snow"`
		Temp struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"temp"`
		Uvi  string `json:"uvi"`
		Wdir struct {
			Degrees string `json:"degrees"`
			Dir     string `json:"dir"`
		} `json:"wdir"`
		Windchill struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"windchill"`
		Wspd struct {
			English string `json:"english"`
			Metric  string `json:"metric"`
		} `json:"wspd"`
		Wx string `json:"wx"`
	} `json:"hourly_forecast"`
	Satellite struct {
		ImageURL    string `json:"image_url"`
		ImageURLIr4 string `json:"image_url_ir4"`
		ImageURLVis string `json:"image_url_vis"`
	} `json:"satellite"`
	History struct {
		Dailysummary []struct {
			Coolingdegreedays       string `json:"coolingdegreedays"`
			Coolingdegreedaysnormal string `json:"coolingdegreedaysnormal"`
			Date                    struct {
				Hour   string `json:"hour"`
				Mday   string `json:"mday"`
				Min    string `json:"min"`
				Mon    string `json:"mon"`
				Pretty string `json:"pretty"`
				Tzname string `json:"tzname"`
				Year   string `json:"year"`
			} `json:"date"`
			Fog                                string `json:"fog"`
			Gdegreedays                        string `json:"gdegreedays"`
			Hail                               string `json:"hail"`
			Heatingdegreedays                  string `json:"heatingdegreedays"`
			Heatingdegreedaysnormal            string `json:"heatingdegreedaysnormal"`
			Humidity                           string `json:"humidity"`
			Maxdewpti                          string `json:"maxdewpti"`
			Maxdewptm                          string `json:"maxdewptm"`
			Maxhumidity                        string `json:"maxhumidity"`
			Maxpressurei                       string `json:"maxpressurei"`
			Maxpressurem                       string `json:"maxpressurem"`
			Maxtempi                           string `json:"maxtempi"`
			Maxtempm                           string `json:"maxtempm"`
			Maxvisi                            string `json:"maxvisi"`
			Maxvism                            string `json:"maxvism"`
			Maxwspdi                           string `json:"maxwspdi"`
			Maxwspdm                           string `json:"maxwspdm"`
			Meandewpti                         string `json:"meandewpti"`
			Meandewptm                         string `json:"meandewptm"`
			Meanpressurei                      string `json:"meanpressurei"`
			Meanpressurem                      string `json:"meanpressurem"`
			Meantempi                          string `json:"meantempi"`
			Meantempm                          string `json:"meantempm"`
			Meanvisi                           string `json:"meanvisi"`
			Meanvism                           string `json:"meanvism"`
			Meanwdird                          string `json:"meanwdird"`
			Meanwdire                          string `json:"meanwdire"`
			Meanwindspdi                       string `json:"meanwindspdi"`
			Meanwindspdm                       string `json:"meanwindspdm"`
			Mindewpti                          string `json:"mindewpti"`
			Mindewptm                          string `json:"mindewptm"`
			Minhumidity                        string `json:"minhumidity"`
			Minpressurei                       string `json:"minpressurei"`
			Minpressurem                       string `json:"minpressurem"`
			Mintempi                           string `json:"mintempi"`
			Mintempm                           string `json:"mintempm"`
			Minvisi                            string `json:"minvisi"`
			Minvism                            string `json:"minvism"`
			Minwspdi                           string `json:"minwspdi"`
			Minwspdm                           string `json:"minwspdm"`
			Monthtodatecoolingdegreedays       string `json:"monthtodatecoolingdegreedays"`
			Monthtodatecoolingdegreedaysnormal string `json:"monthtodatecoolingdegreedaysnormal"`
			Monthtodateheatingdegreedays       string `json:"monthtodateheatingdegreedays"`
			Monthtodateheatingdegreedaysnormal string `json:"monthtodateheatingdegreedaysnormal"`
			Monthtodatesnowfalli               string `json:"monthtodatesnowfalli"`
			Monthtodatesnowfallm               string `json:"monthtodatesnowfallm"`
			Precipi                            string `json:"precipi"`
			Precipm                            string `json:"precipm"`
			Precipsource                       string `json:"precipsource"`
			Rain                               string `json:"rain"`
			Since1jancoolingdegreedays         string `json:"since1jancoolingdegreedays"`
			Since1jancoolingdegreedaysnormal   string `json:"since1jancoolingdegreedaysnormal"`
			Since1julheatingdegreedays         string `json:"since1julheatingdegreedays"`
			Since1julheatingdegreedaysnormal   string `json:"since1julheatingdegreedaysnormal"`
			Since1julsnowfalli                 string `json:"since1julsnowfalli"`
			Since1julsnowfallm                 string `json:"since1julsnowfallm"`
			Since1sepcoolingdegreedays         string `json:"since1sepcoolingdegreedays"`
			Since1sepcoolingdegreedaysnormal   string `json:"since1sepcoolingdegreedaysnormal"`
			Since1sepheatingdegreedays         string `json:"since1sepheatingdegreedays"`
			Since1sepheatingdegreedaysnormal   string `json:"since1sepheatingdegreedaysnormal"`
			Snow                               string `json:"snow"`
			Snowdepthi                         string `json:"snowdepthi"`
			Snowdepthm                         string `json:"snowdepthm"`
			Snowfalli                          string `json:"snowfalli"`
			Snowfallm                          string `json:"snowfallm"`
			Thunder                            string `json:"thunder"`
			Tornado                            string `json:"tornado"`
		} `json:"dailysummary"`
		Date struct {
			Hour   string `json:"hour"`
			Mday   string `json:"mday"`
			Min    string `json:"min"`
			Mon    string `json:"mon"`
			Pretty string `json:"pretty"`
			Tzname string `json:"tzname"`
			Year   string `json:"year"`
		} `json:"date"`
		Observations []struct {
			Conds string `json:"conds"`
			Date  struct {
				Hour   string `json:"hour"`
				Mday   string `json:"mday"`
				Min    string `json:"min"`
				Mon    string `json:"mon"`
				Pretty string `json:"pretty"`
				Tzname string `json:"tzname"`
				Year   string `json:"year"`
			} `json:"date"`
			Dewpti     string `json:"dewpti"`
			Dewptm     string `json:"dewptm"`
			Fog        string `json:"fog"`
			Hail       string `json:"hail"`
			Heatindexi string `json:"heatindexi"`
			Heatindexm string `json:"heatindexm"`
			Hum        string `json:"hum"`
			Icon       string `json:"icon"`
			Metar      string `json:"metar"`
			Precipi    string `json:"precipi"`
			Precipm    string `json:"precipm"`
			Pressurei  string `json:"pressurei"`
			Pressurem  string `json:"pressurem"`
			Rain       string `json:"rain"`
			Snow       string `json:"snow"`
			Tempi      string `json:"tempi"`
			Tempm      string `json:"tempm"`
			Thunder    string `json:"thunder"`
			Tornado    string `json:"tornado"`
			Utcdate    struct {
				Hour   string `json:"hour"`
				Mday   string `json:"mday"`
				Min    string `json:"min"`
				Mon    string `json:"mon"`
				Pretty string `json:"pretty"`
				Tzname string `json:"tzname"`
				Year   string `json:"year"`
			} `json:"utcdate"`
			Visi       string `json:"visi"`
			Vism       string `json:"vism"`
			Wdird      string `json:"wdird"`
			Wdire      string `json:"wdire"`
			Wgusti     string `json:"wgusti"`
			Wgustm     string `json:"wgustm"`
			Windchilli string `json:"windchilli"`
			Windchillm string `json:"windchillm"`
			Wspdi      string `json:"wspdi"`
			Wspdm      string `json:"wspdm"`
		} `json:"observations"`
		Utcdate struct {
			Hour   string `json:"hour"`
			Mday   string `json:"mday"`
			Min    string `json:"min"`
			Mon    string `json:"mon"`
			Pretty string `json:"pretty"`
			Tzname string `json:"tzname"`
			Year   string `json:"year"`
		} `json:"utcdate"`
	} `json:"history"`
	Tide struct {
		TideInfo []struct {
			Lat      string `json:"lat"`
			Lon      string `json:"lon"`
			TideSite string `json:"tideSite"`
			Type     string `json:"type"`
			Tzname   string `json:"tzname"`
			Units    string `json:"units"`
		} `json:"tideInfo"`
		TideSummary []struct {
			Data struct {
				Height string `json:"height"`
				Type   string `json:"type"`
			} `json:"data"`
			Date struct {
				Epoch  string `json:"epoch"`
				Hour   string `json:"hour"`
				Mday   string `json:"mday"`
				Min    string `json:"min"`
				Mon    string `json:"mon"`
				Pretty string `json:"pretty"`
				Tzname string `json:"tzname"`
				Year   string `json:"year"`
			} `json:"date"`
			Utcdate struct {
				Epoch  string `json:"epoch"`
				Hour   string `json:"hour"`
				Mday   string `json:"mday"`
				Min    string `json:"min"`
				Mon    string `json:"mon"`
				Pretty string `json:"pretty"`
				Tzname string `json:"tzname"`
				Year   string `json:"year"`
			} `json:"utcdate"`
		} `json:"tideSummary"`
		TideSummaryStats []struct {
			Maxheight float64 `json:"maxheight"`
			Minheight float64 `json:"minheight"`
		} `json:"tideSummaryStats"`
	} `json:"tide"`
	Webcams []struct {
		CAMURL                string `json:"CAMURL"`
		CURRENTIMAGEURL       string `json:"CURRENTIMAGEURL"`
		WIDGETCURRENTIMAGEURL string `json:"WIDGETCURRENTIMAGEURL"`
		AssocStationID        string `json:"assoc_station_id"`
		Cameratype            string `json:"cameratype"`
		Camid                 string `json:"camid"`
		Camindex              string `json:"camindex"`
		City                  string `json:"city"`
		Country               string `json:"country"`
		Downloaded            string `json:"downloaded"`
		Handle                string `json:"handle"`
		Isrecent              string `json:"isrecent"`
		Lat                   string `json:"lat"`
		Link                  string `json:"link"`
		Linktext              string `json:"linktext"`
		Lon                   string `json:"lon"`
		Neighborhood          string `json:"neighborhood"`
		Organization          string `json:"organization"`
		State                 string `json:"state"`
		Tzname                string `json:"tzname"`
		Updated               string `json:"updated"`
		Zip                   string `json:"zip"`
	} `json:"webcams"`
	Response struct {
		Error struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"error"`
		Features struct {
			Conditions int `json:"conditions"`
			Forecast   int `json:"forecast"`
			Alerts int `json:"alerts"`
			Almanac int `json:"almanac"`
			Astronomy int `json:"astronomy"`
			Forecast10day int `json:"forecast10day"`
			Hourly int `json:"hourly"`
			Satellite int `json:"satellite"`
			Yesterday int `json:"yesterday"`
			Tide int `json:"tide"`
			Webcams int `json:"webcams"`
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
