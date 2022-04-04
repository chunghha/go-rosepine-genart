package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"sync"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

const CANVAS_WIDTH = 3456
const CANVAS_HEIGHT = 2234

var ROSE_PINE_BASE = color.RGBA{25, 23, 36, 255}
var ROSE_PINE_COLORS = []color.RGBA{
	{0xEB, 0x6F, 0x92, 0xFF},
	{0xF6, 0xC1, 0x77, 0xFF},
	{0xEB, 0xBC, 0xBA, 0xFF},
	{0x31, 0x74, 0x8F, 0xFF},
	{0x9C, 0xCF, 0xD8, 0xFF},
	{0xC4, 0xA7, 0xE7, 0xFF},
}

func PrepareCanvas(c *generativeart.Canva) {
	c.SetBackground(ROSE_PINE_BASE)
	c.FillBackground()
	c.SetColorSchema(ROSE_PINE_COLORS)
}

func GenerateCircle(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating a circle png...")

	c := generativeart.NewCanva(CANVAS_WIDTH, CANVAS_HEIGHT)
	PrepareCanvas(c)
	c.Draw(arts.NewColorCircle(1234))
	c.ToPNG("rose_pine_circle.png")
}

func GenerateCircle2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating a circle2 png...")

	c := generativeart.NewCanva(CANVAS_WIDTH, CANVAS_HEIGHT)
	PrepareCanvas(c)
	c.Draw(arts.NewColorCircle2(333))
	c.ToPNG("rose_pine_circle2.png")
}

func GenerateContourLine(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating a contourline png...")

	c := generativeart.NewCanva(CANVAS_WIDTH, CANVAS_HEIGHT)
	PrepareCanvas(c)
	c.Draw(arts.NewContourLine(999))
	c.ToPNG("rose_pine_contourline.png")
}

func GenerateMaze(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating a maze png...")

	c := generativeart.NewCanva(CANVAS_WIDTH, CANVAS_HEIGHT)
	PrepareCanvas(c)
	c.SetLineWidth(1)
	for i := range ROSE_PINE_COLORS {
		c.SetLineWidth(float64((i+1)/3 + 1))
		c.Draw(arts.NewMaze(i*10 + 100))
		c.SetLineColor(ROSE_PINE_COLORS[i])
	}
	c.ToPNG("rose_pine_maze.png")
}

func GenerateNoiseLine(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating a noiseline png...")

	c := generativeart.NewCanva(CANVAS_WIDTH, CANVAS_HEIGHT)
	PrepareCanvas(c)
	c.Draw(arts.NewNoiseLine(3456))
	c.ToPNG("rose_pine_noiseline.png")
}

func GenerateShape(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating a shape png...")

	c := generativeart.NewCanva(CANVAS_WIDTH, CANVAS_HEIGHT)
	PrepareCanvas(c)
	c.Draw(arts.NewRandomShape(111))
	c.ToPNG("rose_pine_shape.png")
}

func main() {
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	wg.Add(1)
	go GenerateCircle(&wg)

	wg.Add(1)
	go GenerateCircle2(&wg)

	wg.Add(1)
	go GenerateContourLine(&wg)

	wg.Add(1)
	go GenerateMaze(&wg)

	wg.Add(1)
	go GenerateNoiseLine(&wg)

	wg.Add(1)
	go GenerateShape(&wg)

	wg.Wait()
	fmt.Println("main finished")
}
