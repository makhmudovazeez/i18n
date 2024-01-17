package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/makhmudovazeez/i18n/helper"
)

var Location string
var Language string
var Debug bool = false

func T(word string) string {
	if len(Location) == 0 {
		if Debug {
			fmt.Println(helper.Warning("set up location"))
		}
		return word
	} else if len(Language) == 0 {
		if Debug {
			fmt.Println(helper.Warning("set up language"))
		}
		return word
	}

	file, err := ioutil.ReadFile(getLocation())
	if err != nil {
		if Debug {
			fmt.Println(helper.Warning(err.Error()))
		}
		return word
	}

	var jsonWords map[string]interface{}

	if err := json.Unmarshal(file, &jsonWords); err != nil {
		if Debug {
			fmt.Println(helper.Warning(err.Error()))
		}
		return word
	}

	wordSlice := strings.Split(word, ".")

	for _, w := range wordSlice {
		result, ok := jsonWords[w]
		if !ok {
			return word
		}

		switch result.(type) {
		case map[string]interface{}:
			jsonWords = jsonWords[w].(map[string]interface{})
			continue
		}

		return fmt.Sprintf("%v", result)
	}

	return word
}

func getLocation() (fileLocation string) {
	if Location[len(Location)-1:len(Location)] == "/" {
		fileLocation = fmt.Sprintf("%v%v.json", Location, Language)
	} else {
		fileLocation = fmt.Sprintf("%v/%v.json", Location, Language)
	}
	return fileLocation
}
