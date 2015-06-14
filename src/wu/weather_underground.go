package wu

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"
)

func getKey() (string,error) {
	if key := os.Getenv("WU_API_KEY"); key != "" {
		return strings.TrimSpace(key), nil
	}

	usr, err := user.Current()
	if err == nil {
		data, err := ioutil.ReadFile(path.Join(usr.HomeDir, ".wuapi"))
		if err == nil {
			return strings.TrimSpace(string(data)), nil
		}
	}
	return "", fmt.Errorf("Failed to load API key from the environment. Please set WU_API_KEY environment variable, or create file .wuapi in your home.")
}

func Query(location string) (*WeatherData,error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}

	location = strings.Replace(location, " ", "_", -1)
	url := fmt.Sprintf("http://api.wunderground.com/api/%s/conditions/forecast/q/%s.json", key, location)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var wd WeatherData
	if err = json.Unmarshal(data, &wd); err != nil {
		return nil, err
	}

	e := wd.Response.Error
	if e.Type != "" {
		return nil, fmt.Errorf("%s (%s)", e.Description, e.Type)
	}
	return &wd, nil
}
