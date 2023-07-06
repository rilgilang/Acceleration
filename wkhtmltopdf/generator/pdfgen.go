package generator

import (
	"bytes"
	"context"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	clonePath        = "./clone/"
	templatePathHtml = "./template.html"
	outputPath       = "./output"
)

type docGenerator struct {
}

func NewDocumentGenerator() *docGenerator {
	return &docGenerator{}
}

func (g *docGenerator) Parser(ctx context.Context, templateFileName string, data interface{}) error {
	time := time.Now().Unix()

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	fmt.Println("time --> ", time)

	//create the file
	err = ioutil.WriteFile(clonePath+strconv.FormatInt(time, 10)+".html", []byte(buf.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g *docGenerator) Resolve(ctx context.Context, data interface{}) error {
	t := time.Now().Unix()
	fmt.Println("t --> ", t)
	err := g.Parser(ctx, templatePathHtml, data)
	if err != nil {
		return err
	}

	//opening the file
	f, err := os.Open(clonePath + strconv.FormatInt(t, 10) + ".html")

	if f != nil {
		defer f.Close()
	}
	if err != nil {

		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {

		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(fmt.Sprintf(`%s/%v.pdf`, outputPath, t))
	if err != nil {
		return err
	}

	return nil
}
