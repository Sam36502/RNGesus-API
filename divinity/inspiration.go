package divinity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	HOLY_LAND_LAT      = 31.777972
	HOLY_LAND_LON      = 35.235806
	HOLY_LAND_EXCLUDES = "minutely,hourly,daily,alerts"
	HOLY_LAND_UNITS    = "metric"
	HOLY_LAND_URL      = "https://api.openweathermap.org/data/2.5/weather"
	HAND_USES          = 12
	ENV_HOLY_LAND_KEY  = "HOLY_LAND_KEY"
)

type Trinity struct {
	Father     float64 `json:"temp"`
	Son        float64 `json:"pressure"`
	HolySpirit float64 `json:"humidity"`
}

type Path struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"deg"`
}

type HandOfGod struct {
	Trinity `json:"main"`
	Path    `json:"wind"`
	Time    time.Time
}

var numberUses int = 0
var currentHand HandOfGod

func InitialiseRNG() error {
	hog, err := getHandOfGod()
	if err != nil {
		return err
	}

	seed := RequestDivineSeed(hog)
	rand.Seed(seed)

	// Discard first few values
	rand.Float64()
	rand.Float64()
	rand.Float64()

	return nil
}

func RequestDivineSeed(hog HandOfGod) int64 {

	var seed int64

	// Consolidate the trinity
	seed = int64((hog.Trinity.Father * hog.Trinity.Son * hog.Trinity.HolySpirit) / 3)
	seed += int64(hog.Trinity.Father)
	seed += int64(hog.Trinity.Son)
	seed += int64(hog.Trinity.HolySpirit)

	// Walk the path
	path := int64(hog.Path.Direction * hog.Path.Speed)
	seed *= path

	// Add time
	seed *= (time.Now().UnixMicro() - hog.Time.UnixMicro())

	return seed
}

func getHandOfGod() (HandOfGod, error) {

	if numberUses < HAND_USES && (currentHand != HandOfGod{}) {
		numberUses++
		return currentHand, nil
	}

	// Create request
	req, err := http.NewRequest(http.MethodGet, HOLY_LAND_URL, nil)
	if err != nil {
		return HandOfGod{}, err
	}

	// Add params
	key := os.Getenv(ENV_HOLY_LAND_KEY)
	if key == "" {
		return HandOfGod{}, errors.New("invalid key configured; check your envs, or the holy land remains closed to you")
	}

	q := req.URL.Query()
	q.Add("lat", fmt.Sprint(HOLY_LAND_LAT))
	q.Add("lon", fmt.Sprint(HOLY_LAND_LON))
	q.Add("exclude", HOLY_LAND_EXCLUDES)
	q.Add("appid", key)
	q.Add("units", HOLY_LAND_UNITS)
	req.URL.RawQuery = q.Encode()

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return HandOfGod{}, err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HandOfGod{}, err
	}

	// Parse response
	var hog HandOfGod
	err = json.Unmarshal(respBody, &hog)
	if err != nil {
		return HandOfGod{}, err
	}

	// Add current time
	hog.Time = time.Now()

	currentHand = hog
	return hog, nil
}
