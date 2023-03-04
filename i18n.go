package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var Location string
var Language string

func T(word string) (string, error) {
	if len(Location) == 0 {
		return "", errors.New("set up location")
	} else if len(Language) == 0 {
		return "", errors.New("set up language")
	}

	file, err := ioutil.ReadFile(getLocation())
	if err != nil {
		return "", err
	}

	var jsonWords map[string]interface{}

	if err := json.Unmarshal(file, &jsonWords); err != nil {
		return "", err
	}

	wordSlice := strings.Split(word, ".")

	for _, w := range wordSlice {
		result, ok := jsonWords[w]
		if !ok {
			return "", errors.New("no such translation")
		}

		switch result.(type) {
		case map[string]interface{}:
			jsonWords = jsonWords[w].(map[string]interface{})
			continue
		}

		return fmt.Sprintf("%v", result), nil
	}

	return "", errors.New("no such translation12312")
}

func getLocation() (fileLocation string) {
	if Location[len(Location)-1:len(Location)] == "/" {
		fileLocation = fmt.Sprintf("%v%v.json", Location, Language)
	} else {
		fileLocation = fmt.Sprintf("%v/%v.json", Location, Language)
	}
	return fileLocation
}
