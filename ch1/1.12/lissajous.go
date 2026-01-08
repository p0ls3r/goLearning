package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
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

type lissajousParams struct {
	cycles            int
	res               float64
	size              int
	frames            int
	delay             int
	switchColorOffset int
}

func NewLissajousParams() lissajousParams {
	return lissajousParams{
		cycles:            10,
		res:               0.0001,
		size:              500,
		frames:            64,
		delay:             4,
		switchColorOffset: 10000,
	}
}

func lissajous(out io.Writer, params lissajousParams) {
	// Use a local, privately-seeded PRNG instead of the package-global one.
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	freq := r.Float64() * 2.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: params.frames}
	phase := 0.0 // phase difference
	for i := 0; i < params.frames; i++ {
		rect := image.Rect(0, 0, 2*params.size+1, 2*params.size+1)
		img := image.NewPaletted(rect, palette)
		colorOffset := 0
		currentIndex := greenIndex
		for t := 0.0; t < float64(params.cycles)*2*math.Pi; t += params.res {
			colorOffset++
			if colorOffset >= params.switchColorOffset {
				colorOffset = 0
				currentIndex++
				if currentIndex > purpleIndex {
					currentIndex = greenIndex
				}
			}
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(params.size+int(x*float64(params.size)+0.5), params.size+int(y*float64(params.size)+0.5),
				uint8(currentIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
