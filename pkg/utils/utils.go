package utils

//converting the data received in JSON form to a form that we can work with i.e. encoding

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// unmarshall =convert data from byte data to origanal Ds

func ParseBody(r *http.Request, x interface{}) {
	//parse the input JSON
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
