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
	w,h,y int
}

//DefaultPrettyPrinter use standard output - terminal
func DefaultPrettyPrinter() *PrettyPrinter {
	return &PrettyPrinter{
		tm.Width(),
		tm.Height(),
		sy,
	}
}

func (p *PrettyPrinter) Title(title string) (err error){
	err = p.print(1,"%v %v",emoji.GreenApple, aurora.BrightCyan(title))
	return
}

func (p *PrettyPrinter) Subtitle(subtitle string) (err error){
	err = p.print(5,"%v %-30v",emoji.GreenCircle, aurora.Cyan(subtitle))
	return
}

func (p *PrettyPrinter) Property(name, prop string) (err error){
	err = p.print(9,"%v : %v",aurora.BrightGreen(name), aurora.BrightYellow(prop))
	return
}

func (p *PrettyPrinter) PropertySlice(name string, slice []string) (err error){
	err = p.print(9,"%v",aurora.BrightGreen(name))
	for i,s := range slice {
		_ = p.print(11,"%v: %-15v",i, aurora.BrightYellow(s))
	}
	return
}

func (p *PrettyPrinter) PropertyMap(name string,mp map[string]string) (err error){
	err = p.print(9,"%v",aurora.BrightGreen(name))
	for k,s := range mp {
		_ = p.print(11,"%v: %-15v",aurora.BrightBlue(k), aurora.BrightYellow(s))
	}
	return
}

func (p *PrettyPrinter) NewLine(){
	p.y++
}


func (p *PrettyPrinter) Flush(){
	if p.w != tm.Width() || p.h != tm.Height() {
		p.w, p.h = tm.Width(), tm.Height()
		p.Clear()
	}
	tm.Flush()
	p.y = sy
}

func (p *PrettyPrinter) Clear(){
	tm.MoveCursor(1,1)
	tm.Clear()
}

func (p *PrettyPrinter) print(x int,format string, a ...interface{}) (err error){
	p.move(x)
	_,err = tm.Printf(format,a...)
	return
}

func (p *PrettyPrinter) move(x int){
	tm.MoveCursor(x,p.y)
	p.y++
}