package compilador
import (
  "bufio"
  "bytes"
  "io"
  )
type Scanner struct{
  r *bufio.Reader
}
//NuvoEscaner regresa instancia
func NewScanner(r io.Reader) *Scanner {
  return &Scanner{r: bufio.NewReader(r)}
}
//Regresa el siguiente token y su valor literal
func (s *Scanner) Scan() (tok Token, lit string){
  ch:= s.read()
  if esEspacio(ch){
    s.unread()
    return s.scanID()
  } else if esLetra(ch){
    s.unread()
    return s.scanID()
  }
switch ch{
case eof:
  return EOF,""
case '>':
  return mayorque, string(ch)
case '<':
  return menorque, string(ch)
case '=':
  return igual, string(ch)
case '{':
  return llaveI, string(ch)
case '}':
  return llaveD,string(ch)
case '(':
  return agrupaI, string(ch)
case')':
  return agrupaD, string(ch)
case'+':
  return suma, string(ch)
case '-':
  return resta, string(ch)
case '*':
  return multi, string(ch)
case '/':
  return divi,string(ch)
case '&':
  return andand,string(ch)
case '|':
  return oror, string(ch)
case '%':
  return diventera, string(ch)
}
  return ILLEGAL, string(ch)
}
//busca espacio
func (s *Scanner) scanEspacio()(tok Token, lit string)  {
  var buf bytes.Buffer
  buf.WriteRune(s.read())
  for {
    if ch:=s.read(); ch==eof{
      break
    }else if !esEspacio(ch){
      s.unread()
      break
    }else{
      buf.WriteRune(ch)
    }
  }
  return WS, buf.String()
}
//busca letras
func (s *Scanner) scanID() (tok Token, lit string){
  var buf bytes.Buffer
  buf.WriteRune(s.read())
  for {
    if ch := s.read(); ch == eof {
      break
      } else if !esLetra(ch) && !esNumero(ch){
        s.unread()
        break
        } else {
          _,_ = buf.WriteRune(ch)
        }
      }
      switch buf.String() {
      case "entero":
        return entero, buf.String()
      case "cadena":
          return cadena, buf.String()
      case "dobles":
        return dobles, buf.String()
      case "char":
        return char, buf.String()
      case "logicos":
        return logicos, buf.String()
      case "decimal":
        return decimal, buf.String()
      case "fecha":
        return fecha, buf.String()
      case "flotante":
        return flotante, buf.String()
      case "largo" :
        return largo, buf.String()
      case "bites":
        return bites, buf.String()
      case "ushort":
        return ushort, buf.String()
      case "ulong":
        return ulong, buf.String()
      case "unet":
        return unet, buf.String()
      case "numera":
        return numera, buf.String()
      case "para":
        return para, buf.String()
      case "porcada":
        return porcada, buf.String()
      case "haz":
        return haz, buf.String()
      case "mientras":
        return mientras, buf.String()
      case "en":
        return en, buf.String()
      case "rompe":
        return rompe, buf.String()
      case "continua":
        return continua, buf.String()
      case "predef":
        return predef, buf.String()
      case "vaya":
        return vaya, buf.String()
      case "regresa":
          return regresa, buf.String()
      case "deten":
        return deten, buf.String()
      case "delega":
        return delega, buf.String()
      case "clase":
        return clase, buf.String()
      case "dinamico":
        return dinamico, buf.String()
      case "interfaz":
        return interfaz, buf.String()
      case "objeto":
        return objeto, buf.String()
      case "arroja":
        return arroja, buf.String()
      case "intenta-agarra":
        return intentaagarra, buf.String()
      case "intenta-finalmente":
        return intentafinalmente, buf.String()
      case "intenta-agarra-finalmente":
        return intentaagarrafinalmente, buf.String()
      case "revisado":
        return revisado, buf.String()
      case "norevisado":
        return norevisado, buf.String()
      case "ajustado":
        return ajustado, buf.String()
      case "cerrado":
        return cerrado, buf.String()
      case "prmt":
        return prmt, buf.String()
      case "ref":
        return ref, buf.String()
      case "fuera":
        return fuera, buf.String()
      case "NomEspacio":
        return NomEspacio, buf.String()
      case "usando":
        return usando, buf.String()
      case "como":
        return como, buf.String()
      case "espera":
        return espera, buf.String()
      case "es":
        return es, buf.String()
      case "nuevo":
        return nuevo, buf.String()
      case "tamañode":
        return tamañode, buf.String()
      case "tipode":
        return tipode, buf.String()
      case "verdadero":
        return verdadero, buf.String()
      case "falso":
        return falso, buf.String()
      case "slvPila":
        return slvPila, buf.String()
      case "nombrede":
        return nombrede, buf.String()
      case "explicito":
        return explicito, buf.String()
      case "implicito":
        return implicito, buf.String()
      case "operador":
        return operador, buf.String()
      case "base":
        return base, buf.String()
      case "esto":
        return esto, buf.String()
      case "suma":
        return suma, buf.String()
      case "resta":
        return resta, buf.String()
      case "divi":
        return divi, buf.String()
      case "TRU":
        return TRU, buf.String()
      case "FLS":
        return  FLS, buf.String()
      case "nulo":
        return nulo, buf.String()
      case "vdd":
        return vdd, buf.String()
      case "fls":
        return fls, buf.String()
      case "pred":
        return pred, buf.String()
      case "si":
        return si, buf.String()
      case "sino":
        return sino, buf.String()
      case "interruptor":
        return interruptor, buf.String()
      case "caso":
        return caso, buf.String()
      case "coment":
        return coment, buf.String()
      case "opcoment":
        return coment, buf.String()
      case "clcoment":
        return clcoment, buf.String()
      case "sen":
        return sen, buf.String()
      case "cos":
        return cos, buf.String()
      case "tan":
        return tan, buf.String()
      case "cot":
        return cot, buf.String()
      case "sec":
        return sec, buf.String()
      case "csc":
        return csc, buf.String()
      case "arcsen":
        return arcsen, buf.String()
      case "arccos":
        return arccos, buf.String()
      case "arctan":
        return arctan, buf.String()
      case "arccot":
        return arccot, buf.String()
      case "arcsec":
        return arcsec, buf.String()
      case "arccsc":
        return arccsc, buf.String()
      case "hsen":
        return hsen, buf.String()
      case "hcos":
        return hcos, buf.String()
      case "htan":
        return htan, buf.String()
      case "hcot":
        return hcot, buf.String()
      case "hsec":
        return hsec, buf.String()
      case "hcsc":
        return hcsc, buf.String()
      case "pot":
        return pot, buf.String()
      case "pub":
        return pub, buf.String()
      case "priv":
        return priv, buf.String()
      case "interno":
        return interno, buf.String()
      case "prtg":
        return prtg, buf.String()
      case "estatico":
        return estatico, buf.String()
      case "abstrct":
          return abstrct, buf.String()
        case "new":
          return new, buf.String()
        case "sobrescribe":
          return sobrescribe, buf.String()
        case "parcial":
          return parcial, buf.String()
        case "sololee":
          return sololee, buf.String()
        case "sella":
          return sella, buf.String()
        case "virtual":
          return virtual, buf.String()
        case "asncro":
          return asncro, buf.String()
        case "extrn":
          return extrn, buf.String()
        case "invar":
          return invar, buf.String()
        case "volatil":
          return volatil, buf.String()
        case "evento":
          return evento, buf.String()
        case "inseguro":
          return inseguro, buf.String()
        case "sbrescribe":
          return sbrescribe, buf.String()
        case "Limpia":
          return Limpia, buf.String()
        case "Escribe":
          return Escribe, buf.String()
        case "EscribeLinea":
          return EscribeLinea, buf.String()
        case "Bip":
          return Bip, buf.String()
        case "Jala":
          return Jala, buf.String()
        case "Empuja":
          return Empuja, buf.String()
        case "LeeLinea":
          return LeeLinea, buf.String()
        case "LeeT":
          return LeeT, buf.String()
        case "lista":
          return lista, buf.String()
        case "cola":
          return cola, buf.String()
        case "pila":
          return pila, buf.String()
        case "asigna":
          return asigna, buf.String()
        case "añade":
          return añade, buf.String()
        case "esperar":
          return esperar, buf.String()
        case "dinamicos":
          return dinamicos, buf.String()
        case "toma":
          return toma, buf.String()
        case "valor":
          return valor, buf.String()
        case "vrbl":
          return vrbl, buf.String()
        case "donde":
          return donde, buf.String()
        case "DE":
          return DE, buf.String()
        case "DONDE":
          return DONDE, buf.String()
        case "SELECCIONA":
          return SELECCIONA, buf.String()
        case "GRUPO":
          return GRUPO, buf.String()
        case "DENTRODE":
          return DENTRODE, buf.String()
        case "ORDENAPOR":
          return ORDENAPOR, buf.String()
        case "JUNTA":
          return JUNTA, buf.String()
        case "DEJA":
          return DEJA, buf.String()
        case "EN":
          return EN, buf.String()
        case "SOBRE":
          return SOBRE, buf.String()
        case "IGUAL":
          return IGUAL, buf.String()
        case "POR":
          return POR, buf.String()
        case "ASND":
          return ASND, buf.String()
        case "DSND":
          return DSND, buf.String()
      }
        return ID, buf.String()
    }
  func (s *Scanner)read() rune{
    ch,_,err :=s.r.ReadRune()
    if err!=nil{
      return eof
    }
    return ch
    }
  func (s * Scanner) unread(){_=s.r.UnreadRune()}
  func esEspacio(ch rune)bool{
    return ch==' ' || ch=='\t'||ch=='\n'
  }
  func esLetra(ch rune)bool{
    return (ch>='a' && ch<='z')||(ch>='A' && ch<='Z')
  }
  func esNumero(ch rune)bool{
    return (ch>='0'&&ch<='9')
  }
  var eof = rune(0)
