package compilador

type Token int
  const (
    ILLEGAL Token = iota
    EOF
    WS
    ID
    igual
    mayorque //>
    menorque //<
    igualmayor //>=
    igualmenor //<=
    igualigual //==
    distinto //!=
    andand //&&
    oror //||
    agrupaD // (
    agrupaI // )
    llaveD //{
    llaveI //}
    entero //int=entero
    cadena //string=cadena
    dobles //double=doble
    char //char=char (caracter)
    logicos //bool
    decimal //decimal
    fecha //datetime
    flotante //float
    largo //long
    bites //byte
    ushort //ushort
    ulong //ulong
    unet //unet
    numera //numeric
    para //for
    porcada //foreach
    haz //do
    mientras //while
    en //in
    rompe //break
    continua //continue
    predef //default
    vaya //goto
    regresa //return
    deten //stop
    delega //delegate
    clase //class
    dinamico //dynamic
    interfaz //interface
    objeto //object
    arroja //throw
    intentaagarra //try catch
    intentafinalmente //try finally
    intentaagarrafinalmente //try catch finally
    revisado //checked
    norevisado //nonchecked
    ajustado //
    cerrado //lock
    prmt //params
    ref //reference
    fuera //out
    NomEspacio //namespace
    usando //using
    como //as
    espera //wait
    es //is
    nuevo //new
    tamañode //sizeof
    tipode //typeof
    verdadero //true
    falso //false
    slvPila //stack
    nombrede //nameof
    explicito //explicit
    implicito //implicit
    operador //operator
    base //base
    esto //this
    suma //+
    resta //-
    multi //*
    divi // /
    TRU //TRUE
    FLS //FALSE
    nulo //null
    vdd //true
    fls //false
    pred //predefinido
    si //if
    sino //else
    interruptor //switch
    caso //case
    coment // //
    opcoment // /*
    clcoment // */
    sen
    cos
    tan
    cot
    sec
    csc
    arcsen
    arccos
    arctan
    arccot
    arcsec
    arccsc
    hsen
    hcos
    htan
    hcot
    hsec
    hcsc
    diventera //%
    pot
    pub //public
    priv //private
    interno //inner
    prtg //protected
    estatico //static
    abstrct //abstract
    new //new
    sobrescribe //override
    parcial //partial
    sololee //JustRead
    sella
    virtual
    asncro
    extrn
    invar
    volatil
    evento
    inseguro
    sbrescribe
    Limpia
    Escribe
    EscribeLinea
    Bip
    Jala
    Empuja
    LeeLinea
    LeeT
    lista
    cola
    pila
    asigna
    añade
    esperar
    dinamicos
    toma
    valor
    vrbl
    donde
    DE
    DONDE
    SELECCIONA
    GRUPO
    DENTRODE
    ORDENAPOR
    JUNTA
    DEJA
    EN
    SOBRE
    IGUAL
    POR
    ASND
    DSND
)
