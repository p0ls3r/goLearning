// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//http://localhost:8000/params?cells=50
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//surface(os.Stdout)
}

func surface(out io.Writer, r *http.Request) {
	cells := cells

	if param := r.URL.Query().Get("cells"); param != "" {
		cells, _ = strconv.Atoi(param)
	}

	head := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	out.Write([]byte(head))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			av := (az + bz + cz + dz) / 4

			if av > 0 {
				polygons := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='blue' fill='lightblue'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				out.Write([]byte(polygons))
			} else {
				polygons := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='red' fill='lightred'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				out.Write([]byte(polygons))
			}
		}
	}
	out.Write([]byte("</svg>"))
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	//r := math.Hypot(x, y) // distance from (0,0)
	//return roundTo(math.Sin(r)/r, 10)
	r := math.Copysign(x, y)
	//return roundTo(r, 10)
	return roundTo(math.Sin(r)/r, 10)
}

func roundTo(n float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(n*multiplier) / multiplier
}
