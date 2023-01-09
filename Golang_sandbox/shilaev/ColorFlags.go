package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Person struct {
	id   int
	name string
}

func (p *Person) newPerson(id int, name string) *Person {
	return &Person{id, name}
}

func (p *Person) toString() string {
	return fmt.Sprintf("id: %d\nname: %s", p.id, p.name)
}

func (p *Person) update(id int, name string) {
	p.id = id
	p.name = name
}

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	blueColor := flag.Bool("blue", false, "display colorized output")
	redColor := flag.Bool("red", false, "display colorized output")
	yellowColor := flag.Bool("yellow", false, "display colorized output")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	if *blueColor {
		colorize(ColorBlue, message)
		return
	} else if *redColor {
		colorize(ColorRed, message)
		return
	} else if *yellowColor {
		colorize(ColorYellow, message)
		return
	} else {
		colorize(ColorBlack, message)
	}
}
