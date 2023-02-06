package strategy

import (
	"fmt"
	"io"
	"os"
)

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type ConsoleSquare struct {
	PrintOutput
}

func (c *ConsoleSquare) Print() error {
	// r := bytes.NewReader([]byte("Circle"))
	// io.Copy(c.Writer, r)
	c.Writer.Write([]byte("Circle"))
	return nil
}

type ImageSquare struct {
	DestinationFIlePath string
	PrintOutput
}

func (t *ImageSquare) Print() error {
	// width := 800
	// height := 600
	// origin := image.Point{0, 0}
	// bgImage := image.NewRGBA(image.Rectangle{
	// 	Min: origin,
	// 	Max: image.Point{X: width, Y: height},
	// })

	// bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	// quality := &jpeg.Options{Quality: 75}
	// draw.Print(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	// squareWidth := 200
	// squareHeight := 200
	// squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	// square := image.Rect(0, 0, squareWidth, squareHeight)
	// square = square.Add(image.Point{
	// 	X: (width / 2) - (squareWidth / 2),
	// 	Y: (height / 2) - (squareHeight / 2),
	// })
	// squareImg := image.NewRGBA(square)
	// draw.Print(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	// w, err := os.Create(t.DestinationFilePath)
	// if err != nil {
	// 	return fmt.Errorf("Error opening image")
	// }
	// defer w.Close()
	// if err = jpeg.Encode(w, bgImage, quality); err != nil {
	// 	return fmt.Errorf("Error writing image to disk")
	// }
	fmt.Println("image square")
	return nil
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (p *PrintOutput) SetLog(w io.Writer) {
	p.LogWriter = w
}

func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (PrintStrategy, error) {
	switch s {
	case "console":
		return &ConsoleSquare{
			PrintOutput: PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case "image":
		return &ImageSquare{
			PrintOutput: PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
