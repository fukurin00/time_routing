package routing

import (
	"log"
	"math"
	"time"
)

type GridMap struct {
	Resolution float64
	Origin     Point
	Width      int
	Height     int
	MapOrigin  Point

	MaxT      int    // =MaxTimeLenghth
	ObjectMap ObjMap //元からある障害物ならTrue
}

func (g GridMap) Ind2Pos(xId, yId int) (float64, float64) {
	x := g.Origin.X + float64(xId)*g.Resolution
	y := g.Origin.Y + float64(yId)*g.Resolution
	return x, y
}

func (g GridMap) Pos2Ind(x, y float64) (int, int) {
	if x < g.Origin.X || y < g.Origin.Y {
		log.Printf("position (%f,%f) is out of map", x, y)
		return 0, 0
	}
	xid := int(math.Round((x - g.Origin.X) / g.Resolution))
	yid := int(math.Round((y - g.Origin.Y) / g.Resolution))
	return xid, yid
}

// initialize gridmap by custom resolution
// using main
func NewGridMapReso(m MapMeta, robotRadius float64, resolution float64, objMap [][2]float64) *GridMap {
	start := time.Now()
	g := new(GridMap)
	g.Resolution = resolution
	g.MapOrigin = m.Origin

	var xList []float64
	var yList []float64

	for _, obj := range objMap {
		xList = append(xList, obj[0])
		yList = append(yList, obj[1])
	}

	maxX := MaxFloat(xList)
	maxY := MaxFloat(yList)
	g.Origin.X = MinFloat(xList)
	g.Origin.Y = MinFloat(yList)

	g.Width = int(math.Round((maxX - g.Origin.X) / resolution))
	g.Height = int(math.Round((maxY - g.Origin.Y) / resolution))

	g.MaxT = MaxTimeLength
	g.ObjectMap = make(map[Index]bool, g.Height)
	count := 0
	for j := 0; j < g.Height; j++ {
		y := g.Origin.Y + float64(j)*g.Resolution
		for i := 0; i < g.Width; i++ {
			if m.Data[i+j*m.W] == 50 {
				g.ObjectMap[newIndex(i, j)] = true
				continue
			}
			x := g.Origin.X + float64(i)*g.Resolution
			g.ObjectMap[newIndex(i, j)] = false
			for _, op := range objMap {
				d := math.Hypot(op[0]-x, op[1]-y)
				// if distance < robot radius, robot cannot path
				if d <= robotRadius {
					g.ObjectMap[newIndex(i, j)] = true
					count += 1
					break
				}
			}
		}
	}

	elaps := time.Since(start).Seconds()
	log.Printf("loading gridmap resolution: %f, takes: %f seconds, obj %d counts, width: %d, height: %d", resolution, elaps, count, g.Width, g.Height)
	return g
}

// initialize custom resolution
// using main
func NewGridMapResoHexa(m MapMeta, robotRadius float64, resolution float64, objMap [][2]float64) *GridMap {
	start := time.Now()
	g := new(GridMap)
	g.MapOrigin = m.Origin
	g.Resolution = resolution

	worldWidth := float64(m.W) * m.Reso
	worldHeight := float64(m.H) * m.Reso

	width := int(math.Ceil(worldWidth / (resolution * math.Sqrt(3) / 2)))
	height := int(math.Ceil(worldHeight/resolution)) + 1

	g.Origin.X = m.Origin.X
	g.Origin.Y = m.Origin.Y

	g.Width = width
	g.Height = height

	g.MaxT = MaxTimeLength
	grey := 0
	white := 0
	black := 0

	g.ObjectMap = make(map[Index]bool)
	for i := 0; i < width; i++ {
		x := m.Origin.X + float64(i)*resolution*math.Sqrt(3)/2
		mi := int(math.Round(float64(i) * resolution / m.Reso))
		if mi >= m.W {
			mi = m.W - 1
		}
		for j := 0; j < height; j++ {
			mj := int(math.Round(float64(j) * resolution / m.Reso))
			if mj >= m.H {
				mj = m.H - 1
			}
			if m.Data[mi+mj*m.W] == 50 {
				g.ObjectMap[newIndex(i, j)] = true
				grey += 1
				continue
			}
			var y float64
			if i%2 == 0 {
				y = m.Origin.Y + float64(j)*resolution
			} else {
				y = m.Origin.Y - m.Reso/2 + float64(j)*resolution
			}
			g.ObjectMap[newIndex(i, j)] = false
			white += 1
			for _, op := range objMap {
				d := math.Hypot(op[0]-x, op[1]-y)
				if d <= robotRadius {
					g.ObjectMap[newIndex(i, j)] = true
					black += 1
					white -= 1
					break
				}
			}
		}
	}

	elaps := time.Since(start).Seconds()
	log.Printf("loading gridmap resolution: %f, takes: %f seconds, white: %d, grey: %d, black: %d, counts, width: %d, height: %d, origin %f,%f", resolution, elaps, white, grey, black, g.Width, g.Height, g.Origin.X, g.Origin.Y)
	return g
}

func (g *GridMap) Pos2IndHexa(x, y float64) (int, int) {
	if x < g.MapOrigin.X || y < g.MapOrigin.Y {
		return 0, 0
	}
	aid := int(math.Round((x - g.Origin.X) / (g.Resolution * math.Sqrt(3) / 2)))
	var bid int
	if aid%2 == 0 {
		bid = int(math.Round((y - g.Origin.Y) / g.Resolution))
	} else {
		bid = int(math.Round((y - g.Origin.Y - g.Resolution/2) / g.Resolution))
	}
	return aid, bid
}

func (g *GridMap) Ind2PosHexa(a, b int) (float64, float64) {
	if a < 0 || b < 0 {
		return 0, 0
	}
	x := g.Origin.X + (g.Resolution*math.Sqrt(3)/2)*float64(a)
	var y float64
	if a%2 == 0 {
		y = g.Origin.Y + g.Resolution*float64(b)
	} else {
		y = g.Origin.Y - g.Resolution/2 + g.Resolution*float64(b)
	}
	return x, y
}
