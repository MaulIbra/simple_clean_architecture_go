/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func GetEnv(key, defaultValue string) string {
	if envVal, exists := os.LookupEnv(key); exists {
		return envVal
	}
	return defaultValue
}

func DecodePathVariabel(val string, r *http.Request) string {
	param := mux.Vars(r)
	return param[val]
}

func JsonDecoder(val interface{}, r *http.Request) error {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&val)
	if err != nil {
		return err
	}
	return nil
}
