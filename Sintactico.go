package compilador
import (
  "fmt"
  "io"
)
type siStatement struct{}

type Parser struct{
  s *Scanner
  buf struct {
    tok Token
    lit string
    n int
  }
}
func NewParser(r io.Reader) *Parser{
  return &Parser{s: NewScanner(r)}
}
//si "if"
func (p *Parser)Parse()(*siStatement, error){
  stmt := &siStatement{}

if tok,lit := p.scanIgnoreEspacio(); tok != si{
  return nil, fmt.Errorf("Se encontro %q, esperaba: si", lit)
  }
if tok, lit := p.scanIgnoreEspacio(); tok != agrupaD{
  return nil, fmt.Errorf("Se encontro %q, esperaba: (",lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != ID{
  return nil, fmt.Errorf("Se encontro %q, esperaba: un ID",lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != mayorque|menorque|igual+igual {
  return nil, fmt.Errorf("Se encontro %q, espraba: un operador (>,<,==)", lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != ID{
  return nil, fmt.Errorf("Se encontro %q, esperaba: un ID", lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != agrupaI{
  return nil, fmt.Errorf("Se encontro %q, esperaba: un )",lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != llaveD{
  return nil, fmt.Errorf("Se encontro %q, esperaba: una {",lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != ID{
  return nil, fmt.Errorf("Se encontro %q, esperaba: un ID",lit)
}
if tok, lit := p.scanIgnoreEspacio(); tok != llaveI{
  return nil, fmt.Errorf("Se encontro %q, esperaba: una }",lit)
}
return stmt, nil
}

//SCAN
func (p *Parser) scan()(tok Token, lit string){
  if p.buf.n!=0{
    p.buf.n=0
  }
  return p.buf.tok, p.buf.lit
}

//IGNORA ISPACIO
func (p *Parser) scanIgnoreEspacio()(tok Token, lit string){
    tok, lit = p.scan()
    if tok == WS {
      tok,lit = p.scan()
    }
    return
}
//unscan
func (p *Parser) unscan() { p.buf.n = 1 }
