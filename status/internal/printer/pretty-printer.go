package printer

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
)

type PrettyPrinter struct {
	out io.Writer
}


func DefaultPrettyPrinter() *PrettyPrinter {
	return &PrettyPrinter{
		out: os.Stdout,
	}
}

func (p *PrettyPrinter) Title(title string) (err error) {
	err = p.print( "%s %s", emoji.GreenApple, aurora.BrightCyan(title))
	p.NewLine()
	return
}

func (p *PrettyPrinter) Subtitle(subtitle string) (err error){
	err = p.print("%4s %s",emoji.GreenCircle, aurora.Cyan(subtitle))
	p.NewLine()
	return
}

func (p *PrettyPrinter) Property(name, prop string) (err error){
	err = p.print("%8s%-10s : %v"," ",aurora.BrightGreen(name), aurora.BrightYellow(prop))
	p.NewLine()
	return
}

func (p *PrettyPrinter) PropertySlice(name string, slice []string) (err error){
	err = p.print("%12s%-10s","",aurora.BrightGreen(name))
	p.NewLine()
	for i,s := range slice {
		_ = p.print("%32s%-2v : %s"," ",i, aurora.BrightYellow(s))
		p.NewLine()
	}
	p.NewLine()
	return
}

func (p *PrettyPrinter) PropertyMap(name string,mp map[string]string) (err error){
	err = p.print("%8s%-10s","",aurora.BrightGreen(name))
	p.NewLine()
	for k,s := range mp {
		_ = p.print("%12s%-20v : %s"," ",aurora.BrightBlue(k), aurora.BrightYellow(s))
		p.NewLine()
	}
	p.NewLine()
	return
}

func (p *PrettyPrinter) NewLine(){
	_ = p.print("\n")
}

func (p *PrettyPrinter) print(format string, a ...interface{}) (err error){
	_,err = fmt.Fprintf(p.out, format,a...)
	return
}