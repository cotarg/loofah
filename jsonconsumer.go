package main

import "encoding/json"

type JSONLogObject struct {
	date       string            `json:"date"`
	date_day   string            `json:"date_day"`
	date_month string            `json:"date_month"`
	date_time  string            `json:"date_time"`
	hostname   string            `json:"hostname"`
	message    string            `json:"message"`
	pid        string            `json:"pid"`
	rig        map[string]string `json:"rig"`
	syslog     map[string]string `json:"syslog"`
	version    string            `json:"version"`
}

func (j *JSONObject) ingestLogline(loglineString []byte) {
	var formattedLog JSONLogObject
	err := json.Unmarshal(loglineString, &formattedLog)
	return formattedLog
}
