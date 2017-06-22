package main

import "fmt"

const (
	WIHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WIHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	string := []string{"WHITE", "BLACK", "BLUE", "RED", "YeLLOW"}
	return string[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WIHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("we have %d boxes in our set\n", len(boxes))
	fmt.Println("the volume of the first one is", boxes[0].Volume())
	fmt.Println("the color or the last one is", boxes[len(boxes)-1])
	fmt.Println("the biggest  one is", boxes.BiggestsColor().String())
	fmt.Println("Let's paint them all black")
	boxes.PaintBlack()
	fmt.Println("the color or the second one is ", boxes[1].color.String())
	fmt.Println("Obviously, now the bigggest one is", boxes.BiggestsColor().String())
}
