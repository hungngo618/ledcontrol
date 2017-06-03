package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/hungngo618/ledcontrol/config"
	"github.com/hungngo618/ledcontrol/db"
)

func init() {
	if err := config.LoadFromEnv(); err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	if err := db.Init(&config.DB); err != nil {
		log.Fatal("failed to connect to db: ", err)
	}
}

func LedTurnOffHanler(w http.ResponseWriter, r *http.Request) {
	err := db.TurnLedOff()
	if err != nil {
		log.Error("failed to turn off led: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var m map[string]string
	m = make(map[string]string)
	m["status"] = "success"

	jsonString, _ := json.Marshal(m)

	fmt.Fprintln(w, string(jsonString))
}
func LedTurnOnHandler(w http.ResponseWriter, r *http.Request) {
	err := db.TurnLedOn()
	if err != nil {
		log.Error("failed to turn off led: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var m map[string]string
	m = make(map[string]string)
	m["status"] = "success"

	jsonString, _ := json.Marshal(m)

	fmt.Fprintln(w, string(jsonString))

}
func LedStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := db.GetLedStatus()
	if err != nil {
		log.Error("failed to turn off led: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var m map[string]int
	m = make(map[string]int)
	m["led"] = status

	jsonString, _ := json.Marshal(m)

	fmt.Fprintln(w, string(jsonString))
}

func main() {
	http.HandleFunc("/status", LedStatusHandler)
	http.HandleFunc("/on", LedTurnOnHandler)
	http.HandleFunc("/off", LedTurnOffHanler)
	log.Infof("Server run in %s\n", config.DB.ListenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.DB.ListenPort), nil))

}