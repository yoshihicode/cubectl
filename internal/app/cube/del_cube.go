package cube

import (
	"context"
	g "cubectl/internal/graphics"
	"cubectl/internal/logger"
	"cubectl/internal/terminal"
	"fmt"
	"math"
	"strings"
	"time"
)

func RenderD(ctx context.Context, opts Options) error {
	clog := logger.New()

	fmt.Printf("cube.geometry.io \"cube\" deletion triggered...\n")
	time.Sleep(500 * time.Millisecond)

	logs := []string{
		clog.Swarn(logger.Message{File: "delete.go", Line: 10, Text: "Pre-stop hook: checking geometric stability..."}),
		clog.Serror(logger.Message{File: "delete.go", Line: 20, Text: "Structural integrity failure! Cube is dispersing..."}),
	}

	logIndex := 0

	// Cube vertices
	v := g.VertexData{
		[3]int{-2, -2, -2},
		[3]int{2, -2, -2},
		[3]int{-2, 2, -2},
		[3]int{2, 2, -2},
		[3]int{-2, -2, 2},
		[3]int{2, -2, 2},
		[3]int{-2, 2, 2},
		[3]int{2, 2, 2},
	}

	f := g.FaceData{
		[]int{0, 1, 3, 2},
		[]int{5, 4, 6, 7},
		[]int{0, 1, 5, 4},
		[]int{3, 2, 6, 7},
		[]int{0, 2, 6, 4},
		[]int{3, 1, 5, 7},
	}

	m := g.NewModel(v, f, 8)
	s := terminal.New()
	if err := s.Init(); err != nil {
		return err
	}
	defer func() {
		s.Close()
		fmt.Printf("cube.geometry.io \"cube\" deleted\n")
	}()

	s.SetOutputMode()
	s.Clear()
	w, h := s.Size()

	ch := make(chan terminal.Event)
	go keyEvent(ch, s)

	yaw := 0.2
	pitch := 0.5
	scale := 0.6

	cx := w / 2
	cy := h / 2

	drawString := func(x, y int, str string) {
		for i, r := range str {
			s.SetCell(x+i, y, r, terminal.ColorDefault, terminal.ColorBlack)
		}
	}

	faceData := m.GetShape(yaw, pitch, scale, cx, cy)
	iscollapse := false

	dxs := make([][]float64, len(faceData))
	dys := make([][]float64, len(faceData))

	for i, fd := range faceData {
		dxs[i] = make([]float64, len(fd.Outline))
		dys[i] = make([]float64, len(fd.Outline))

		for j := range fd.Outline {
			dxs[i][j] = 0
			dys[i][j] = 0
		}
	}
	startTime := time.Now()

loop:
	for {

		select {
		case ev := <-ch:
			switch ev.Type {
			case terminal.EventKey:
				if ev.Key == terminal.KeyCtrlC || ev.Key == terminal.KeyEsc {
					break loop
				}
			}
		default:
			s.Clear()

			r := 0
			for l := range logIndex {
				lines := strings.Split(logs[l], "\n")
				for _, line := range lines {
					drawString(0, r, line)
					r = r + 1
				}
			}
			if logIndex < len(logs) {
				logIndex++
			}

			if !iscollapse && logIndex == len(logs) && time.Since(startTime) > time.Second {
				iscollapse = true
			}

			if iscollapse {
				elapsed := time.Since(startTime).Seconds()
				centerX, centerY := float64(cx), float64(cy)

				for i, fd := range faceData {
					for j := range fd.Outline {
						p := fd.Outline[j]

						// Calculate distance from the center
						relX := float64(p.X) - centerX
						relY := float64(p.Y) - centerY
						dist := math.Sqrt(relX*relX + relY*relY)

						// Rotation speed (faster near the center, accelerates over time)
						angleStep := (0.5 / (dist + 1)) + (elapsed * 0.1)

						// Calculate the new position of the point after rotation
						newX := relX*math.Cos(angleStep) - relY*math.Sin(angleStep)*1.2
						newY := relX*math.Sin(angleStep) + relY*math.Cos(angleStep)*1.2

						// Movement towards the outward
						dxs[i][j] = (newX*1.15 - relX)
						dys[i][j] = (newY*1.15 - relY)
					}
				}
			}

			allZero := true
			for i, fd := range faceData {
				for j := range fd.Outline {
					fd.Outline[j].X = int(float64(fd.Outline[j].X) + dxs[i][j])
					fd.Outline[j].Y = int(float64(fd.Outline[j].Y) + dys[i][j])
					p := fd.Outline[j]
					if p.X > -1 && p.Y > -1 && p.X < w && p.Y < h {
						allZero = false
						s.SetCell(p.X, p.Y, ' ', terminal.ColorDefault, terminal.ColorWhite)
					}
				}
			}

			s.Flush()
			if allZero {
				time.Sleep(500 * time.Millisecond)
				break loop
			}
			time.Sleep(50 * time.Millisecond)

		}
	}

	return nil
}
