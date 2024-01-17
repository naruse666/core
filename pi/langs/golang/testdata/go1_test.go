/* Copyright (c) 2018, The gide / Goki Authors. All rights reserved. */
/* Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file. */

package gotest

func sldff() {
	n += wasSpace & ^isSpace
}

//func sldkf() {
// putting CompositeLit at start of Expr prevents selector from operating on it.
//	return readInt(buf, unsafe.Offsetof(Dirent{}.Ino), unsafe.Sizeof(Dirent{}.Ino))
//}


func sldk() {
	spd := make(sparseDatas, 0, s.MaxEntries()) // litnum must be earlier for this to work..
}

//func lddldl() {
//	e.buf[3*i+7] = "\x22\x11\x11"[i] // lit is before slice in prim expr -- conflicts with putting later..
//	e.buf[3*i+8] = "\x00\x01\x01"[i]
//}

type Checker struct {
	*bufio.Writer
	conf *Config
	fset *token.FileSet
	pkg  *Package
	*Info
	objMap map[Object]*declInfo   // maps package-level object to declaration info
}

func dkfn() {
	se.fieldEncs[i](e, fv, opts)
}

func sldfk() {
	next := func() *Entry {
		if !e.Children {
			return nil
		}
	}
}


func fadk() {
	switch v.(type) {
	case string, []byte:
		return v, nil
	}
	return fmt.Sprintf("%v", v), nil
}

func slfk() {
	f1 := c.rune(re.Rune[j:j+1], re.Flags)
}

func tsts() {
	// calling a func that returns a func
	prfForVersion(version, suite)(masterSecret, preMasterSecret, masterSecretLabel, seed)
}

// /usr/local/Cellar/go/1.11.3/libexec/src/context/context.go:253 has this code and it fails at
// the indicated line, but is fine here.. hmmm..

func propagateCancel(parent Context, child canceler) {
	if parent.Done() == nil {
		return // parent is never canceled
	}
	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent has already been canceled
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}

func slsl() {
	p.children = make(map[canceler]struct{})
}

func skdfa() {
	switch {
	case level == NoCompression:
		d.window = make([]byte, maxStoreBlockSize)
		d.fill = (*compressor).fillStore
	
	}
}

const (
	magicGNU, versionGNU     = "ustar ", " \x00"
	magicUSTAR, versionUSTAR = "ustar\x00", "00"
	trailerSTAR              = "tar\x00"
)

func (x *stringVal) string() string {
	x.mu.Lock()
}

func dir(path string) string {
	if i := strings.LastIndexAny(path, `/\`); i > 0 {
		return path[:i]
	}
	// i <= 0
	return "."
}

func (obj *object) setOrder(order uint32)     { assert(order > 0); obj.order_ = order }
func (obj *object) setColor(color color)      { assert(color != white); obj.color_ = color }

// todo: doesn't deal with new in params list even though it should..
func (cmap CommentMap) Update(old, new Node) Node {
	if list := cmap[old]; len(list) > 0 {
		delete(cmap, old)
		cmap[new] = append(cmap[new], list...)
	}
	return new
}

func checkPkgFiles(files []*ast.File) {
	type bailout struct{}

	// if checkPkgFiles is called multiple times, set up conf only once
	conf := types.Config{
		FakeImportC: true,
		Error: func(err error) {
			if !*allErrors && errorCount >= 10 {
				panic(bailout{})
			}
			report(err)
		},
		Importer: importer.For(*compiler, nil),
		Sizes:    types.SizesFor(build.Default.Compiler, build.Default.GOARCH),
	}
}

var header = []byte(`// Copyright 2017 The Go Authors. All rights reserved.`)

var header = []byte(`// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go run make_tables.go. DO NOT EDIT.

package bits

`)

// using int as a var name here..
func modf(f float64) (int float64, frac float64) {
	if f < 1 {
		switch {
		case f < 0:
			int, frac = Modf(-f)
			return -int, -frac
		case f == 0:
			return f, f // Return -0, -0 when f == -0
		}
		return 0, f
	}

	x := Float64bits(f)
	e := uint(x>>shift)&mask - bias

	// Keep the top 12+e bits, the integer part; clear the rest.
	if e < 64-12 {
		x &^= 1<<(64-12-e) - 1
	}
	int = Float64frombits(x)
	frac = f - int
	return
}

/*
func (z *Float) sqrtDirect(x *Float) {
	u := new(Float)
	ng := func(t *Float) *Float {
		u.prec = t.prec
		u.Mul(t, t)        // u = t²
		u.Add(u, x)        //   = t² + x
		u.Mul(half, u)     //   = ½(t² + x)
		return t.Quo(u, t) //   = ½(t² + x)/t
	}
}
*/

func NewFloat(x float64) *Float {
	if math.IsNaN(x) {
		panic(ErrNaN{"NewFloat(NaN)"})
	}
	return new(Float).SetFloat64(x)
}

// todo: parser doesn't quite deal with this correctly -- multi-line backtick literals!

func stsf() {
	switch typ = implicitArrayDeref(x.typ.Underlying()); t := typ.(type) {
	case *Basic:
	}
}

func (*BadExpr) exprNode() {}

func ddd() {
	var abf = [...]string{
		full,
		full + ".gox",
		pkgdir + "lib" + pkg + ".so",
		pkgdir + "lib" + pkg + ".a",
		full + ".o"	,
	}
		
}

func tstr() {
	for i, expr := range []ast.Expr{e.Low, e.High, e.Max} { // added special case for this..
	}
}

func rsww() {
	switch e := e.(type) {
	case *ast.CompositeLit:
		var typ, base Type
		switch {
		case e.Type != nil:
		}
	}
}

func slfk() {
	x.typ = &Signature{
		// this is the []*Var ambiguity -- parsed as mult instead of ptr..
		params:   NewTuple(append([]*Var{NewVar(token.NoPos, check.pkg, "", x.typ)}, params...)...),
		results:  sig.results,
		variadic: sig.variadic,
	}
}

func ruf() {
	return &(*structTypeUncommon)(unsafe.Pointer(t)) // this is ok
	
	return &(*structTypeUncommon)(unsafe.Pointer(t)).u
	// the extra selector on top of the convert parens is not working
	// -- selector needs to be outside but is lower in order
	// moving it before convertparens breaks other things..
	// added a special case for a final selector on top of a convert.. :)
}

func sddfaf() {
	switch x.typ.Underlying().(*Basic).kind {
		default:
	}
}

func Pipe() (*PipeReader, *PipeWriter) {
	p := &pipe{
		wrCh: make(chan []byte),
		rdCh: make(chan int),
		done: make(chan struct{}),
	}
	return &PipeReader{p}, &PipeWriter{p}
}

func (eofReader) Read([]byte) (int, error) {
	return 0, EOF
}

func round(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}

func (s *ss) error(err error) {
	panic(scanError{err})
}

var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

func aaaa() {
	mvp := calcMVP(
		t.size.X, t.size.Y,
		minX, minY,
		maxX, minY,
		minX, maxY,
		)
	 // note: trick for extra comma was Expr ',' ?ArgsList -- key is ?
	 //  -- introducing any other ',' rule doesn't work since this needs to be first and will match
}

// this is Go's "most vexing parse" from a top-down perspective:
// Note: now solved by having a priority @CompositeLit case for var = and regular asgn := = cases
// gives @CompositeLit first crack for those cases and then generic expr is backup

var MultSlice = p[2]*Rule // todo: not working

var SliceAr1 = []Rule{}

var SliceAry = [25]*Rule{}

var SliceAry = []*Rule{} // todo: ? not excluding here

var RuleMap map[string]*Rule // looks like binary *

var val = a[2] * b[0]

// exclude rule -- two rules fwd and back:
// ?'key:map' '[' ? ']' '*' 'Name' ?'.' ?'Name
//  + start at ', go forward to match name, pkg.name -- exclude if no match
//  + go back.. 
// range is 
// start at *, 
// backtrack: if a given parse fails... nah, way too complicated..

var TextViewSelectors = []string{":active", ":focus", ":inactive", ":selected", ":highlight"}

func mulm() {
	return f64.Aff3{
		a[0]*b[0] + a[1]*b[3],
		a[0]*b[1] + a[1]*b[4],
		a[0]*b[2] + a[1]*b[5] + a[2],

		a[3]*b[0] + a[4]*b[3],
		a[3]*b[1] + a[4]*b[4],
		a[3]*b[2] + a[4]*b[5] + a[5],
	}
}

var ifa interface{}

func slfaa() {
	a <- b
	if err := <-errCh; err != nil {
		return nil, err
	}
}

func (tv *TreeView) FocusChanged2D(change gi.FocusChanges) {
	switch change {
	case gi.FocusInactive: // don't care..
	case gi.FocusActive:
	}
}


func baf() {
outer:
	for {
		select {
		case <-w.winClose:
		case <-w.winOpen:
		case i := <-w.winOpen:
		case j = <-w.winOpen:
		}
	}
}


func setScreen(scrIdx int, dpi, pixratio float32, widthPx, heightPx, widthMM, heightMM, depth int, sname *C.char, snlen C.int) {
	theApp.mu.Lock()
}

func saff() {
	switch apv := aps.Value.(type) {
		case ki.BlankProp:
	}
}

var TextViewSelectors = []string{":active", ":focus", ":inactive", ":selected", ":highlight",}

func ffb() {
}

func (tv *TextView) FindNextLink(pos TextPos) (TextPos, TextRegion, bool) {
	
}

func ffi() {
	a++
}

func sfa() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v %v", a, i)
	}
}

type Ityp int

type Ptyp *string

type Mtyp map[string]int

type Sltyp []float32

type Sttyp struct {
	A int
	B float32
}

type Artyp []beebs

type Rndyp peeps

func bb() {
	a := pv.FuncMonk(27)[2]
}


func adkf() {
	for range sr.Text {
		if unicode.IsSpace(sr.Text[0]) {
		}
	}
}

var StyleValueTypes = map[reflect.Type]struct{}{
	units.KiT_Value: {Key: "value"},
	KiT_Color:       {},
	KiT_ColorSpec:   {},
	KiT_Matrix2D:    {},
}

func extfun() {
	rs.Raster.SetStroke(
		Float32ToFixed(pc.StrokeWidth(rs)),
		Float32ToFixed(pc.StrokeStyle.MiterLimit),
		pc.capfunc(), nil, nil, pc.joinmode(), // todo: supports leading / trailing caps, and "gaps"
		dash, 0	)
	rs.Scanner.SetClip(rs.Bounds)
}

func (w *Window) SendKeyChordEvent(popup bool, r rune, mods ...key.Modifiers) {
	ke := key.ChordEvent{}
	ke.SetTime()
	ke.SetModifiers(mods...)
	ke.Rune = r
	ke.Action = key.Press
	w.SendEventSignal(&ke, popup)
}

func (fl *FontLib) InitFontPaths(paths ...string) {
	if len(fl.FontPaths) > 0 {
		return
	}
	fl.AddFontPaths(paths...)
}

func (ft FileTime) String(reg string, pars int) string {
	return (time.Time)(ft).Format("Mon Jan  2 15:04:05 MST 2006")
}

var _ArgDataFlags_index = [...]uint8{0, 13, 26, 39}

var FileInfoProps = ki.Props{
	"CtxtMenu": ki.PropSlice{
		{"Duplicate", ki.Props{
			"updtfunc": ActionUpdateFunc(func(fii interface{}, act *gi.Button) {
				fi := fii.(*FileInfo)
				act.SetInactiveState(fi.IsDir())
			}),
		}},
		{"Delete", ki.Props{
			"desc":    "Ok to delete this file?  This is not undoable and is not moving to trash / recycle bin",
			"confirm": true,
			"updtfunc": ActionUpdateFunc(func(fii interface{}, act *gi.Button) {
				fi := fii.(*FileInfo)
				act.SetInactiveState(fi.IsDir())
			}),
		}},
		{"Rename", ki.Props{
			"desc": "Rename file to new file name",
			"Args": ki.PropSlice{
				{"New Name", ki.Props{
					"default-field": "Name",
				}},
			},
		}},
	},
}

func aaa() {
	sf, ok := pv.(func(it interface{}, act *gi.Button) key.Chord)	
}

func ccc() {
	if sf, ok := pv.(ShortcutFunc); ok {
		ac.Shortcut = sf(md.Val, ac)
	} else if sf, ok := pv.(func(it interface{}, act *gi.Button) key.Chord); ok {
		ac.Shortcut = sf(md.Val, ac)
	} else {
		MethodViewErr(vtyp, fmt.Sprintf("ActionView for Method: %v, shortcut-func must be of type ShortcutFunc", methNm))
	}
}


func bbb() {
	a := struct{}{}
}

func (tv *TableView) RowGrabFocus(row int) *gi.WidgetBase {
	
	tv.inFocusGrab = slice{}

	defer func() { tv.inFocusGrab = false 	}
	
	defer func() { tv.inFocusGrab = false 	}()
	
	return nil
}

func sfa() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v %v", a, i)
		a++
		p := a * i
	}
	tv.inFocusGrab = true
	defer func() { tv.inFocusGrab = false }()
	tv.inFocusGrab = true
}

func tst() {
	if kit.Enums.TypeRegistered(nptyp) { // todo: bitfield
		vv := EnumValueView{}
		vv.Init(&vv)
		return &vv
	} else if _, ok := it.(fmt.Stringer); ok { // use stringer
		vv := ValueViewBase{}
		vv.Init(&vv)
		return &vv
	} else {
		vv := IntValueView{}
		vv.Init(&vv)
		return &vv
	}
}


func dkf() {
	goto pil
pil:
	return nil
}

func (tv *TreeView) FocusChanged2D(change gi.FocusChanges) {
	switch change {
	case gi.FocusInactive: // don't care..
	case gi.FocusActive:
	}
}

func adlf() {
	switch pr := bprpi.(type) {
	case map[string]interface{}:
		wb.SetIconProps(ki.Props(pr))
	case ki.Props:
		wb.SetIconProps(pr)
	}	
}

var _ArgDataFlags_index = [...]uint8{0, 13, 26, 39}

func sld() {
	<-TextViewBlinker.C
}

func main() {
	if sz > max {
		*ch = (*ch)[:max]
	}
}

func (tv *TextView) FindNextLink(pos TextPos) (TextPos, TextRegion, bool) {
	
}

func tst() {
	nwSz := gi.Vec2D{mxwd, off + extraHalf}.ToPointCeil()
}

func tst() {
	a := tv.Renders[ln].Links 

	if !tv.HasLinks && tv.Renders[ln].Links > 0 {
		tv.HasLinks = true
	}
}

func tst() {
	tvn, two := data.(ki.Ki).Embed(giv.KiT_TreeView).(*giv.TreeView)
	for a, b := range cde {
	}
}

var PiViewProps = ki.Props{
	"MainMenu": ki.PropSlice{
		"updtfunc": giv.ActionUpdateFunc(func(pvi interface{}, act *gi.Button) {
			pv := pvi.(*PiView)
			act.SetActiveState(pv.Prefs.ProjFile != "")
		}),
		"offguy": true,
	},
}

import "github.com/goki/gi/gi"

import (
	gi "github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	gogide "github.com/goki/gide/gide"
	"cogentcore.org/core/pi"
	"cogentcore.org/core/pi/piv"
)

var av1, av2 int

type Pvsi struct {
	Af int
	Bf string
}

func (ps *Pvsi) tst() {
	txt += rs[sd-1].String()
	txt += rs[i].String()
	fmt.Println(ps.Errs[len(ps.Errs)-1].Error())
}	

func tst() {
	win.OSWin.SetCloseReqFunc(func(w oswin.Window) {
		if !inClosePrompt {
			inClosePrompt = true
			if pv.Changed {
				gi.ChoiceDialog(vp, gi.DlgOpts{Title: "Close Without Saving?",
					Prompt: "Do you want to save your changes?  If so, Cancel and then Save"},
					[]string{"Close Without Saving", "Cancel"},
					win.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						switch sig {
						case 0:
							w.Close()
						case 1:
							// default is to do nothing, i.e., cancel
						}
					})
			} else {
				w.Close()
			}
		}
	})
}

func tst(txt string, amt int) (bool, *Rule) {
	txt += rs[sd-1].String()
	txt += rs[i].String()
	fmt.Println(ps.Errs[len(ps.Errs)-1].Error())
}	

func tst() {
	r := &(*rs)[i]
}	

func tst() {
	rs := &ps.Matches[abc]
}

func tst() {
	rs := Matches[scope][scope]
}	

func tst() {
	if !inClosePrompt {
		if pv.Changed {
			ChoiceDialog(func(ab int) {
				break
				return
				})
		} else {
			w.Close()
		}
	}
}


func tst() {
   SetCloseReqFunc(func(w win) {
 		if !inClosePrompt {
			if pv.Changed {
				ChoiceDialog(func(ab int) {
					break
					return
					})
			} else {
				w.Close()
			}
		}
	})
}


func (pv *PiView) ConfigSplitView() {
	Connect(func(sig int64) {
		switch sig {
		case int64(TreeViewSelected):
			break
		}
	})
}

var MakeSlice = make([]Rule, 100) // make and new require special rules b/c take type args

var MakeSlice = make([][][]*Rule, 100)

func (pv *PiView) OpenTestTextTab() {
	if ctv.Buf != &pv.TestBuf {
		ctv.SetBuf(&pv.TestBuf)
	}
}

func (ev Steps) MarshalJSON() ([]byte, error)  {
	return kit.EnumMarshalJSON(ev)
}

func (ev *Steps) UnmarshalJSON(b []byte) error { return kit.EnumUnmarshalJSON(ev, b)
}

// was not dealing with all-in-one-line case -- needs to insert EOS before } 
func (ev *Steps) UnmarshalJSON(b []byte) error { return kit.EnumUnmarshalJSON(ev, b) }

func tst() {
	tokSrc := string(ps.TokenSrc(pos))
}

func tst() {
	pv.SaveParser()
	pv.GetPrefs()
	Trace.Out(ps, pr, Run, creg.St, creg, trcAst, fmt.Sprintf("%v: optional rule: %v failed", ri, rr.Rule.Name()))
}

var unaryptr = 25 * *(ptr+2)  // directly to rhs or depth sub of it
var multexpr = 25 * (ptr + 2)
var multex = 25 * ptr + 25 * *ptr // 
var a,b,c,d = 32

var TextViewSelectors = []string{":active", ":focus", ":inactive", ":selected", ":highlight"}

func (pr *Rule) BaseIface() reflect.Type {
	return reflect.TypeOf((*Parser)(nil)).Elem()
}


func (pr *Rule) AsParseRule() *Rule {
	return pr.This().Embed(KiT_Rule).(*Rule)
}


func test() {
	RuleMap = map[string]*Rule{}
}

// interface{} here not working
func (pr *Rule) CompileAll(ps *State) bool {
	pr.SetRuleMap(ps)
	allok := true
	pr.FuncDownMeFirst(0, pr.This(), func(k ki.Ki, level int, d interface{}) bool {
		pri := k.Embed(KiT_Rule).(*Rule)
		ok := pri.Compile(ps)
		if !ok {
			allok = false
		}
		return true
	})
	return allok
}

func test() {
	if pr.Rule[0] == '-' {
		rstr = rstr[1:]
		pr.Reverse = true
	} else {
		pr.Reverse = false
	}
}

type Rule struct {
	OnePar
	ki.Node
	Off       bool     `desc:"disable this rule -- useful for testing"`
}

func tst() {
	oswin.TheApp.SetQuitCleanFunc(func() {
		fmt.Printf("Doing final Quit cleanup here..\n")
	})
}

func (pr *Parser) LexErrString() string {
	return pr.LexState.Errs.AllString()
}

func tst() {
	a = !pr.LexState.AtEol()

	pr.LexState.Filename = !pr.LexState.AtEol()

	if !pr.Sub.LexState.AtEol() && cpos == pr.LexState.Pos {
		msg := fmt.Sprintf("did not advance position -- need more rules to match current input: %v", string(pr.LexState.Src[cpos:]))
		pr.LexState.Error(cpos, msg)
		return nil
	}
}

var ext = strings.ToLower(filepath.Ext(flag.Arg(0)))

func tst() {
		if path == "" && proj == "" {
			if flag.NArg() > 0 {
	 			ext := strings.ToLower(filepath.Ext(flag.Arg(0)))
				if ext == ".gide" {
	 				proj = flag.Arg(0)
				} else {
	 				path = flag.Arg(0)
	 			}
	 		}
		}
	recv := gi.Node2DBase{}
}

func (pr *Parser) Init() {
	pr.Lexer.InitName(&pr.Lexer, "Lexer")
}

func (pr *Parser) Init2(a int, fname string, amap map[string]string) bool {
	pr.Parser.InitName(&pr.Parser, "Parser")
}

func (pr *Parser) Init3(a, b int, fname string) (bool, string) {
	pr.Ast.InitName(&pr.Ast, "Ast")
}

func (pr *Parser) Init4(a, b int, fname string) (ok bool, name string) {
	pr.LexState.Init()
}

// SetSrc sets source to be parsed, and filename it came from
func (pr *Parser) SetSrc(src [][]rune, fname string) {
}

func main() {

	if this > that {
		break
	} else {
		continue
	}

	if peas++; this > that {
		fmt.Printf("test")
		break
	}

	if this > that {
		break
	} else if something == other {
		continue
	}

	if a := b; b == a {
		fmt.Printf("test")
		break
	} else {
		continue
	}

	if a > b {
		b++
	}

	switch vvv := av; nm {
	case "baby":
		nm = "maybe"
		a++
	case "not":
		i++
		p := a * i
	default:
		non := "anon"
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%v %v", a, i)
		a++
		p := a * i
	}

	for a, i := range names {
		for i < 100 {
			for {
				fmt.Printf("%v %v", a, i)
			}
		}
	}

	fmt.Printf("starting test")
	defer my.Widget.UpdateEnd(updt)
	goto bypass

	if a == b {
		fmt.Printf("equal")
	} else if a > b {
		so++
		be--
		it := 20
	} else {
		fmt.Printf("long one")
	}

bypass:
	fmt.Printf("here")
	return
	return something
	{
		nvar := 22
		nvar += function(of + some + others)
	}
}

const neg = -1

const neg2 = -(2+2)

const Prec1 = ((2-1) * 3)

const (
	parn = 1 + (2 + 3)
	PrecedenceS2 = 25 / (3.14 + -(2 + 4)) > ((2 - 5) * 3)
)

const Precedence2 = -(3.14 + 2 * 4)

// The lexical acts
const (
	// Next means advance input position to the next character(s) after the matched characters
	Next Actions = 4.92

	// Name means read in an entire name, which is letters, _ and digits after first letter
	// position will be advanced to just after
	Name
)

type MyFloat float64

type AStruct struct {
	AField int
	TField gi.Widget `desc:"tagged"`
	AField []string
	MField map[string]int
}

var Typeo int

var ExprVar = "testofit"

var ExprTypeVar map[string]string 

var ExprInitMap = map[string]string{
	"does": {Val: "this work?", Bad: "dkfa"},
}

var ExprSlice = abc[20]

var ExprSlice2 = abc[20:30]

var ExprSlice3 = abc[20:30] + abc[:] + abc[20:] + abc[:30] + abc[20:30:2]

var ExprSelect = abc.Def

var ExprCvt = int(abc) // works with basic type names -- others are FunCall

var TypPtr *Fred

var ExprCvt2 = map[string]string(ab)

var tree = map[token.Tokens]struct{}(optMap)

var tree = (map[token.Tokens]struct{})(optMap)

var partyp = (*int)(tree)

var ExprTypeAssert = absfr.(gi.TreeView)

var ExprTypeAssertPtr = absfr.(*gi.TreeView)

var methexpr = abc.meth(a-b * 2 + bf.Meth(22 + 55) / long.meth.Call(tree))

var ExprMeth = abc.meth(c)

var ExprMethLong = long.abc.meth(c)

var ExprFunNil = fun()

var ExprFun = meth(2 + 2)

var ExprFun = meth(2 + 2, fslaf)

var ExprFunElip = meth(2 + 2, fslaf...)


func main() {
	a <- b
	c++
	c[3] = 42 * 17
 	bf := a * b + c[32]
	d += funcall(a, b, c...)
	fmt.Printf("this is ok", gi.CallMe(a.(tree) + b.(*tree) + int(22) * string(17)))
}

func mainrun() {
	oswin.TheApp.SetName("pie")
	oswin.TheApp.SetAbout(`<code>Pie</code> is the interactive parser (pi) editor written in the <b>GoGi</b> graphical interface system, within the <b>Goki</b> tree framework.  See <a href="https://cogentcore.org/core/pi">Gide on GitHub</a> and <a href="https://cogentcore.org/core/pi/wiki">Gide wiki</a> for documentation.<br>
<br>
Version: ` + pi.VersionInfo())
	if peas++; this > that {
		fmt.Printf("test")
		break
	}
	if this > that {
		fmt.Printf("test")
		break
	} else {
		continue
	}

	if this > that {
		fmt.Printf("test")
		break
	} else if something == other {
		continue
	}

	oswin.TheApp.SetQuitCleanFunc(func() {
		fmt.Printf("Doing final Quit cleanup here..\n")
	})

	pi.InitPrefs()

	var path string
	var proj string
	// process command args
	if len(os.Args) > 1 {
		flag.StringVar(&path, "path", "", "path to open -- can be to a directory or a filename within the directory")
		flag.StringVar(&proj, "proj", "", "project file to open -- typically has .gide extension")
		// todo: other args?
		flag.Parse()
		if path == "" && proj == "" {
			if flag.NArg() > 0 {
	 			ext = strings.ToLower(filepath.Ext(flag.Arg(0)))
				if ext == ".gide" {
	 				proj = flag.Arg(0)
				} else {
	 				path = flag.Arg(0)
	 			}
	 		}
		}
	}

	recv := gi.Node2DBase{}
	recv.InitName(&recv, "pie_dummy")

	inQuitPrompt := false
	oswin.TheApp.SetQuitReqFunc(func() {
		if !inQuitPrompt {
			inQuitPrompt = true
			if gide.QuitReq() {
				oswin.TheApp.Quit()
			} else {
				inQuitPrompt = false
			}
		}
	})

	if proj != "" {
		proj, _ = filepath.Abs(proj)
	 	gide.OpenGideProj(proj)
	} else {
		if path != "" {
			path, _ = filepath.Abs(path)
		}
		gide.NewGideProjPath(path)
	}

	piv.NewPiView()

	// above NewGideProj calls will have added to WinWait..
	gi.WinWait.Wait()
}

var someother int


type Lang interface {
	// Parser returns the pi.Parser for this language
	Parser() *Parser

	// ParseFile does the complete processing of a given single file, as appropriate
	// for the language -- e.g., runs the lexer followed by the parser, and
	// manages any symbol output from parsing as appropriate for the language / format.
	ParseFile(fs *FileState)
	
	// LexLine does the lexing of a given line of the file, using existing context
	// if available from prior lexing / parsing. Line is in 0-indexed "internal" line indexes.
	// The rune source information is assumed to have already been updated in FileState.
	// languages can run the parser on the line to augment the lex token output as appropriate.
	LexLine(fs *FileState, line int) lex.Line
}

var TextViewProps = ki.Props{
	"white-space":      gi.WhiteSpacePreWrap,
	"font-family":      "Go Mono",
	"border-width":     0, // don't render our own border
	"cursor-width":     units.NewValue(3, units.Px),
	"border-color":     &gi.Prefs.Colors.Border,
	"border-style":     gi.BorderSolid,
	"padding":          units.NewValue(2, units.Px),
	"margin":           units.NewValue(2, units.Px),
	"vertical-align":   gi.AlignTop,
	"text-align":       gi.AlignLeft,
	"tab-size":         4,
	"color":            &gi.Prefs.Colors.Font,
	"background-color": &gi.Prefs.Colors.Background,
	TextViewSelectors[TextViewActive]: ki.Props{
		"background-color": "highlight-10",
	},
	TextViewSelectors[TextViewFocus]: ki.Props{
		"background-color": "lighter-0",
	},
	TextViewSelectors[TextViewInactive]: ki.Props{
		"background-color": "highlight-20",
	},
	TextViewSelectors[TextViewSel]: ki.Props{
		"background-color": &gi.Prefs.Colors.Select,
	},
	TextViewSelectors[TextViewHighlight]: ki.Props{
		"background-color": &gi.Prefs.Colors.Highlight,
	},
}


