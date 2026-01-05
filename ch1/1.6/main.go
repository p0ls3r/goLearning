package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{41, 196, 41, 0xFF}, color.RGBA{9, 237, 226, 0xFF}, color.RGBA{9, 66, 237, 0xFF}, color.RGBA{188, 9, 237, 0xFF}}

const (
	blackIndex  = 0
	greenIndex  = 1
	cyanIndex   = 2
	blueIndex   = 3
	purpleIndex = 4
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles            = 10     // number of complete x oscillator revolutions
		res               = 0.0001 // angular resolution
		size              = 500    // image canvas covers [-size..+size]
		nframes           = 64     // number of animation frames
		delay             = 4      // delay between frames in 10ms units
		switchColorOffset = 10000  // number of points before switching color
	)
	// Use a local, privately-seeded PRNG instead of the package-global one.
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	freq := r.Float64() * 2.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorOffset := 0
		currentIndex := greenIndex
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			colorOffset++
			if colorOffset >= switchColorOffset {
				colorOffset = 0
				currentIndex++
				if currentIndex > purpleIndex {
					currentIndex = greenIndex
				}
			}
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(currentIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
