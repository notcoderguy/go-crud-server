package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DecodeBody(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.Unmarshal(body, &v)
}