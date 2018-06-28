package clusterer

import (
	"math"
	"image/color"
	"image"
	"image/draw"
	"sync"
	"fmt"
	"os"
	"image/png"
	"time"
)

func TimeDecorator(f func() []interface{}) {
	start := time.Now()
	fmt.Println(start)
	r := f()
	fmt.Println(time.Now().Sub(start))
	fmt.Println(r...)
}

func initRegularPoints(maxx, maxy int) []IPoint {
	points := []IPoint{}
	for x := 0; x < maxx; x++ {
		for y := 0; y < maxy; y++ {
			points = append(points, NewPoint(float64(x), float64(y)))
		}
	}
	return points
}

func initPointsOnCircle(r int) []IPoint {
	points := []IPoint{}
	R := float64(r)
	step := math.Asin(1.0/R)
	x, y := math.MaxFloat64, math.MaxFloat64
	angle := 0.0
	for angle < 2*math.Pi {
		y1, x1 := math.Round(math.Sin(angle)*R), math.Round(math.Cos(angle)*R)
		if x1 != x || y1 != y {
			x, y = x1, y1
			points = append(points, NewPoint(x+R, y+R))
		}
		angle += step
	}
	return points
}

func HSVtoRGB(h, s, v float64) (r, g, b uint8) {
	hi := int(h/60) % 6
	vmin := (100.0 - s) * v / 100.0
	a := (v - vmin) * float64(int(h)%60) / 60.0
	vinc := vmin + a
	vdec := v - a
	//v, vinc, vmin, vdec
	V, Vinc, Vmin, Vdec := uint8(255*v/100), uint8(255*vinc/100), uint8(255*vmin/100), uint8(255*vdec/100)
	switch hi {
	case 0:
		return V, Vinc, Vmin
	case 1:
		return Vdec, V, Vmin
	case 2:
		return Vmin, V, Vinc
	case 3:
		return Vmin, Vdec, V
	case 4:
		return Vinc, Vmin, V
	case 5:
		return V, Vmin, Vdec
	}
	return 0, 0, 0
}

// vis writes a .png for extra credit 2.
func vis(clusters Clusters, centroids Centroids, points []IPoint, fn string) {
	l := len(clusters)
	colors := make([]color.NRGBA, l)
	h, s, v := 0.0, 100.0, 100.0
	hStep := 360.0 / float64(l)
	for i := range colors {
		r, g, b := HSVtoRGB(h, s, v)
		colors[i] = color.NRGBA{r, g, b, 255}
		h = h + hStep
	}
	maxx, maxy := 0.0, 0.0
	for _, p := range points {
		maxx = math.Max(maxx, p.GetCoord(0))
		maxy = math.Max(maxy, p.GetCoord(0))
	}
	bounds := image.Rect(0, 0, int(maxx)+1, int(maxy)+1)
	im := image.NewNRGBA(bounds)
	draw.Draw(im, bounds, image.NewUniform(color.Gray16{0x8888}), image.ZP, draw.Src)

	//for i, cluster := range clusters {
	//	for _, p := range cluster {
	//		im.SetNRGBA(int(points[p].GetCoord(0)), int(points[p].GetCoord(1)), colors[i])
	//	}
	//}
	var wg sync.WaitGroup
	for i, p := range points {
		wg.Add(1)
		go func(i int, p IPoint) {
			defer wg.Done()
			found := false
			for j, c := range clusters {
				if indexOf(i, c) != -1 {
					im.SetNRGBA(int(p.GetCoord(0)), int(p.GetCoord(1)), colors[j])
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Point", i, int(p.GetCoord(0)), int(p.GetCoord(1)), "not found in any cluster")
				im.SetNRGBA(int(p.GetCoord(0)), int(p.GetCoord(1)), color.NRGBA{0,0,0,1})
			}
		}(i, p)
	}
	wg.Wait()

	if centroids != nil {
		for i, c := range centroids {
			im.SetNRGBA(int(points[c].GetCoord(0)), int(points[c].GetCoord(1)), color.NRGBA{colors[i].R, colors[i].G, colors[i].B, 128})

		}
	}

	f, err := os.Create(fn + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(f, im)
	if err != nil {
		fmt.Println(err)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}
