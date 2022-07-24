package routes

import "time"

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL				string	`json:"url"`
	CustomShort	string
	Expiry			time.Duration
	XRateRemaining	int
	XRateLimitRest	time.Duration
}