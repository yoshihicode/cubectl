package pod

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

	fmt.Printf("pod.geometry.io \"pod\" deletion triggered...\n")
	time.Sleep(500 * time.Millisecond)

	logs := []string{
		clog.Swarn(logger.Message{File: "delete.go", Line: 10, Text: "Pre-stop hook: checking geometric stability..."}),
		clog.Serror(logger.Message{File: "delete.go", Line: 20, Text: "Structural integrity failure! Pod is dispersing..."}),
	}

	logIndex := 0

	// pea pod vertices
	v := g.VertexData{
		[3]int{6, 1, 0}, // 0: Right tip

		// --- Slice 1 (Right-mid) ---
		[3]int{4, 1, -1}, // 1: Back
		[3]int{4, 1, 1},  // 2: Front
		[3]int{4, 0, 0},  // 3: Bottom curve
		[3]int{4, 2, 0},  // 4: Top curve

		// --- Slice 2 (Right-center) ---
		[3]int{2, 0, -2}, // 5: Back
		[3]int{2, 0, 2},  // 6: Front
		[3]int{2, -1, 0}, // 7: Bottom curve
		[3]int{2, 1, 0},  // 8: Top curve

		// --- Slice 3 (Left-center) ---
		[3]int{-2, 0, -2}, // 9: Back
		[3]int{-2, 0, 2},  // 10: Front
		[3]int{-2, -1, 0}, // 11: Bottom curve
		[3]int{-2, 1, 0},  // 12: Top curve

		// --- Slice 4 (Left-mid) ---
		[3]int{-4, 1, -1}, // 13: Back
		[3]int{-4, 1, 1},  // 14: Front
		[3]int{-4, 0, 0},  // 15: Bottom curve
		[3]int{-4, 2, 0},  // 16: Top curve

		[3]int{-6, 1, 0}, // 17: Left tip
	}

	// pea pod faces
	f := g.FaceData{
		// --- Top Tip Cap ---
		[]int{0, 1, 4}, // Front-left
		[]int{0, 4, 2}, // Front-right
		[]int{0, 2, 3}, // Back-right
		[]int{0, 3, 1}, // Back-left

		// --- Upper Segment ---
		[]int{1, 5, 8, 4}, // Front-left panel
		[]int{4, 8, 6, 2}, // Front-right panel
		[]int{2, 6, 7, 3}, // Back-right panel
		[]int{3, 7, 5, 1}, // Back-left panel

		// --- Middle Segment ---
		[]int{5, 9, 12, 8},
		[]int{8, 12, 10, 6},
		[]int{6, 10, 11, 7},
		[]int{7, 11, 9, 5},

		// --- Lower Segment ---
		[]int{9, 13, 16, 12},
		[]int{12, 16, 14, 10},
		[]int{10, 14, 15, 11},
		[]int{11, 15, 13, 9},

		// --- Bottom Tip Cap ---
		[]int{17, 16, 13}, // Front-left
		[]int{17, 14, 16}, // Front-right
		[]int{17, 15, 14}, // Back-right
		[]int{17, 13, 15}, // Back-left
	}

	m := g.NewModel(v, f, 8)
	s := terminal.New()
	if err := s.Init(); err != nil {
		return err
	}
	defer func() {
		s.Close()
		fmt.Printf("pod.geometry.io \"pod\" deleted\n")
	}()

	s.SetOutputMode()
	s.Clear()
	w, h := s.Size()

	ch := make(chan terminal.Event)
	go keyEvent(ch, s)

	yaw := 0.0
	pitch := 0.0
	scale := 0.3

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
						if relX == 0 || relY == 0 {
							relX += 0.5
							relY += 0.5
						}
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
						s.SetCell(p.X, p.Y, ' ', terminal.ColorDefault, terminal.ColorGreen)
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
