package main

import (
	"gosl/plt"
	"gosl/utl"
	"math"
)

func main() {

	// grid size
	xmin, xmax, N := -math.Pi/2.0+0.1, math.Pi/2.0-0.1, 21

	// mesh grid
	X, Y, F := utl.MeshGrid2dF(xmin, xmax, xmin, xmax, N, N, func(x, y float64) (z float64) {
		m := math.Pow(math.Cos(x), 2.0) + math.Pow(math.Cos(y), 2.0)
		z = -math.Pow(m, 2.0)
		return
	})

	// configuration
	a := &plt.A{
		NumFmt:  "%.1f",
		Lw:      0.8,
		CbarLbl: "$f(x,y)$",
		SelectC: "yellow",
		SelectV: -2.5,
		Nlevels: 10,
	}

	plt.Reset(true, nil)
	plt.Equal()
	plt.ContourF(X, Y, F, a)
	plt.SetLabels("$x$", "$y$", nil)

	plt.Save("output", "myplot")
}
