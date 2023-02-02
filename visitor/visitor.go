package visitor

import (
	"fmt"
	"io"
	"os"
)

type MessageA struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type Visitable interface {
	Accept(Visitor)
}

type MessageVisitor struct{}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

type TestHelper struct {
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor2 interface {
	Visit(ProductInfoRetriever)
}

type Visitable2 interface {
	Accept(Visitor2)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) Accept(v Visitor2) {
	v.Visit(p)
}

func (p *Product) GetName() string {
	return p.Name
}

type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
	Names       string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.Names = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v Visitor2) {
	v.Visit(f)
}
