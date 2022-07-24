package routes


type request struct {
	URL
	CustomShort
	Expiry
}

type response struct {
	URL
	CustomeShort
	Expiry
	XRateRemaining
	XRateRest
}