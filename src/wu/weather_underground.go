// Package wu provides the utility methods for querying the Weather Underground
// REST API.
package wu

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
)

func getKey() (string, error) {
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

func cacheKey(url string) string {
	h := fnv.New32a()
	h.Write([]byte(url))
	return path.Join(os.TempDir(), fmt.Sprintf("wucli_%d", h.Sum32()))
}

func lookup(url string) []byte {
	key := cacheKey(url)
	info, err := os.Stat(key)
	if err != nil {
		return nil
	}

	duration := time.Now().Sub(info.ModTime())
	if duration < time.Minute*5 {
		data, err := ioutil.ReadFile(key)
		if err != nil {
			return nil
		}
		return data
	}
	return nil
}

func save(url string, data []byte) {
	ioutil.WriteFile(cacheKey(url), data, 0644)
}

// Query the Weather Undergraound REST API for weather data regarding some 
// location. This function caches the responses returned by the REST API 
// for future reference.
func Query(location string) (*WeatherData, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}

	location = strings.Replace(location, " ", "_", -1)
	url := fmt.Sprintf("http://api.wunderground.com/api/%s/conditions/forecast/q/%s.json", key, location)
	var data []byte
	data = lookup(url)
	if data == nil {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		data, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		save(url, data)
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
