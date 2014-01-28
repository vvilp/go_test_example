package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	"code.google.com/p/plotinum/vg"
	"image/color"
	"math/rand"
)

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

// randomTriples returns some random x, y, z triples
// with some interesting kind of trend.
func randomTriples(n int) plotter.XYZs {
	data := make(plotter.XYZs, n)
	for i := range data {
		if i == 0 {
			data[i].X = rand.Float64()
		} else {
			data[i].X = data[i-1].X + 2*rand.Float64()
		}
		data[i].Y = data[i].X + 10*rand.Float64()
		data[i].Z = data[i].X
	}
	return data
}

func test1() {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4, 4, "points.png"); err != nil {
		panic(err)
	}
}

func test2() {
	rand.Seed(int64(0))
	n := 100
	bubbleData := randomTriples(n)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Bubbles"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, err := plotter.NewBubbles(bubbleData, vg.Points(5), vg.Points(5))
	if err != nil {
		panic(err)
	}
	bs.Color = color.RGBA{R: 196, B: 128, A: 255}
	p.Add(bs)

	if err := p.Save(10, 5, "bubble.png"); err != nil {
		panic(err)
	}
}

func randomPoint(n int) plotter.XYZs {
	data := make(plotter.XYZs, n)
	for i := range data {
		data[i].X = 100*rand.Float64()
		data[i].Y = 100*rand.Float64()
		data[i].Z = 100*rand.Float64()
	}
	return data
}

func test_point() {
	rand.Seed(int64(0))
	points_data := randomPoint(200)
	points_data2 := randomPoint(50)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Points"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, _ := plotter.NewBubbles(points_data, vg.Points(5), vg.Points(5))
	bs2, _ := plotter.NewBubbles(points_data2, vg.Points(5), vg.Points(5))


	bs.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	bs2.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	p.Add(bs)
	p.Add(bs2)

	if err := p.Save(10, 10, "points.png"); err != nil {
		panic(err)
	}
}

func main() {
	test_point()
}

