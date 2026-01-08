package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := NewLissajousParams()

	url, _ := url.ParseQuery(r.URL.RawQuery)

	cycles := url.Get("cycles")
	if cycles != "" {
		params.cycles, _ = strconv.Atoi(cycles)
	}

	resolution := url.Get("res")
	if resolution != "" {
		params.res, _ = strconv.ParseFloat(resolution, 64)
	}
	size := url.Get("size")
	if size != "" {
		params.size, _ = strconv.Atoi(size)
	}

	frames := url.Get("frames")
	if frames != "" {
		params.frames, _ = strconv.Atoi(frames)
	}

	delay := url.Get("delay")
	if delay != "" {
		params.delay, _ = strconv.Atoi(delay)
	}

	switchColorOffset := url.Get("switchColorOffset")
	if switchColorOffset != "" {
		params.switchColorOffset, _ = strconv.Atoi(switchColorOffset)
	}

	lissajous(w, params)
}
