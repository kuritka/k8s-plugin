package printer

import (
	tm "github.com/buger/goterm"
	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
)

const (
	sy = 2
)

type PrettyPrinter struct {
	y int
}

//DefaultPrettyPrinter use standard output - terminal
func DefaultPrettyPrinter() *PrettyPrinter {
	return new(PrettyPrinter)
}

func (p *PrettyPrinter) Title(title string) (err error){
	p.move(1)
	_,err = tm.Printf("%v %v",emoji.GreenApple, aurora.BrightCyan(title))
	return
}

func (p *PrettyPrinter) Subtitle(subtitle string) (err error){
	p.move(5)
	_,err = tm.Printf("%v %v",emoji.GreenCircle, aurora.Cyan(subtitle))
	return
}

func (p *PrettyPrinter) Property(name, prop string) (err error){
	p.move(9)
	_,err = tm.Printf("%v : %v",aurora.BrightGreen(name), aurora.BrightYellow(prop))
	return
}

func (p *PrettyPrinter) Flush(){
	tm.Flush()
	p.y = sy
}

func (p *PrettyPrinter) Clear(){
	tm.Clear()
	p.y = sy
}

func (p *PrettyPrinter) move(x int){
	tm.MoveCursor(x,p.y)
	p.y++
}