package routes

import "time"

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL				string	`json:"url"`
	CustomShort		string	`json:"custom_short"`
	Expiry			time.Duration	`json:"expiry"`
	XRateRemaining	int				`json:"xRateRemaining"`
	XRateLimitRest	time.Duration	`json:"XRateLimitRest"`
}