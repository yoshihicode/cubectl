package logs

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	g "cubectl/internal/graphics"
	"cubectl/internal/logger"
)

type Options struct {
	Name   *string
	Follow bool
	Tail   int32
}

func Log(ctx context.Context, ots Options) error {
	cubeName := "cube"
	if !isEmpty(ots.Name) {
		cubeName = *ots.Name
	}

	clog := logger.New()
	logs := []string{
		clog.Swarn(logger.Message{File: "loader.go", Line: 0, Text: "Warning: This output is a joke."}),
		clog.Serror(logger.Message{File: "loader.go", Line: 223, Text: "Error loading kubeconfig:\nunable to read config file \"/home/user/.kube/config\": no such file or directory"}),
		clog.Serror(logger.Message{File: "round_trippers.go", Line: 45, Text: "Failed to create Kubernetes client:\nno configuration has been provided"}),
		clog.Serror(logger.Message{File: "command.go", Line: 112, Text: "error: unknown command \"kubectl\""}),
		clog.Swarn(logger.Message{File: "command.go", Line: 112, Text: "This is not \"kubectl\" but \"cubectl\"\nDid you mean this?\n    kubectl"}),
		clog.Sinfo(logger.Message{File: "cube.go", Line: 112, Text: fmt.Sprintf("Initializing cube rendering engine for %q", cubeName)}),
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	lcnt := 0
	step := 1

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

	yaw := 0.48
	pitch := 0.24
	scale := 0.4
	twoPi := 2 * math.Pi
Loop:
	for {
		select {
		case <-ch:
			break Loop
		default:
			if lcnt < len(logs) {
				fmt.Println(logs[lcnt])
				lcnt++
			} else {
				if ots.Follow || ots.Tail > 0 {
					clog.Info(logger.Message{
						File: "cube.go",
						Line: 88,
						Text: fmt.Sprintf("Telemetry: yaw=%.2f pitch=%.2f scale=%.1f", yaw, pitch, scale),
					})
					if step >= 20 {
						step = 1
					} else {
						step++
					}
					yaw = math.Mod(yaw+0.08, twoPi)
					pitch = math.Mod(pitch+0.04, twoPi)
					drawCube(&m, yaw, pitch, scale)
					time.Sleep(300 * time.Millisecond)
				} else {
					drawCube(&m, yaw, pitch, scale)
					break Loop
				}
			}

			if ots.Tail > 0 {
				ots.Tail--
				if ots.Tail == 0 {
					break Loop
				}
			}

		}
	}

	return nil
}

func drawCube(m *g.Model, yaw, pitch, scale float64) {
	d := [40][40]int{}
	faceData := m.GetShape(yaw, pitch, scale, 20, 10)
	for _, fd := range faceData {
		for _, p := range fd.Outline {
			d[p.X][p.Y] = 1
		}
	}

	for y := 0; y < 20; y++ {
		for x := 0; x < 40; x++ {
			if d[x][y] > 0 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isEmpty(s *string) bool {
	return s == nil || *s == ""
}
