// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"image"
	"strings"
	"unicode"

	"github.com/chewxy/math32"
	"github.com/goki/gi/complete"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/cursor"
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/oswin/mimedata"
	"github.com/goki/gi/oswin/mouse"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
	"github.com/goki/ki/bitflag"
	"github.com/goki/ki/kit"
)

////////////////////////////////////////////////////////////////////////////////////////
// CompletionData
type CompleteData struct {
	Func        complete.Func `desc:"function to get the list of possible completions"`
	Context     interface{}   `desc:"the object that implements complete.Func"`
	Completions []string
	Seed        string `desc:"current completion seed"`
}

////////////////////////////////////////////////////////////////////////////////////////
// TextField

// TextField is a widget for editing a line of text
type TextField struct {
	WidgetBase
	Txt          string                  `json:"-" xml:"text" desc:"the last saved value of the text string being edited"`
	Placeholder  string                  `json:"-" xml:"placeholder" desc:"text that is displayed when the field is empty, in a lower-contrast manner"`
	Edited       bool                    `json:"-" xml:"-" desc:"true if the text has been edited relative to the original"`
	FocusActive  bool                    `json:"-" xml:"-" desc:"true if the keyboard focus is active or not -- when we lose active focus we apply changes"`
	EditTxt      []rune                  `json:"-" xml:"-" desc:"the live text string being edited, with latest modifications -- encoded as runes"`
	MaxWidthReq  int                     `desc:"maximum width that field will request, in characters, during Size2D process -- if 0 then is 50 -- ensures that large strings don't request super large values -- standard max-width can override"`
	StartPos     int                     `xml:"-" desc:"starting display position in the string"`
	EndPos       int                     `xml:"-" desc:"ending display position in the string"`
	CursorPos    int                     `xml:"-" desc:"current cursor position"`
	CharWidth    int                     `xml:"-" desc:"approximate number of chars that can be displayed at any time -- computed from font size etc"`
	SelectStart  int                     `xml:"-" desc:"starting position of selection in the string"`
	SelectEnd    int                     `xml:"-" desc:"ending position of selection in the string"`
	SelectMode   bool                    `xml:"-" desc:"if true, select text as cursor moves"`
	TextFieldSig ki.Signal               `json:"-" xml:"-" view:"-" desc:"signal for line edit -- see TextFieldSignals for the types"`
	RenderAll    TextRender              `json:"-" xml:"-" desc:"render version of entire text, for sizing"`
	RenderVis    TextRender              `json:"-" xml:"-" desc:"render version of just visible text"`
	StateStyles  [TextFieldStatesN]Style `json:"-" xml:"-" desc:"normal style and focus style"`
	FontHeight   float32                 `json:"-" xml:"-" desc:"font height, cached during styling"`
	Completion   CompleteData            `json:"-" xml:"-" desc:"functions and data for textfield completion"`
}

var KiT_TextField = kit.Types.AddType(&TextField{}, TextFieldProps)

var TextFieldProps = ki.Props{
	"border-width":     units.NewValue(1, units.Px),
	"border-color":     &Prefs.Colors.Border,
	"border-style":     BorderSolid,
	"padding":          units.NewValue(4, units.Px),
	"margin":           units.NewValue(1, units.Px),
	"text-align":       AlignLeft,
	"color":            &Prefs.Colors.Font,
	"background-color": &Prefs.Colors.Control,
	TextFieldSelectors[TextFieldActive]: ki.Props{
		"background-color": "lighter-0",
	},
	TextFieldSelectors[TextFieldFocus]: ki.Props{
		"border-width":     units.NewValue(2, units.Px),
		"background-color": "samelight-80",
	},
	TextFieldSelectors[TextFieldInactive]: ki.Props{
		"background-color": "highlight-10",
	},
	TextFieldSelectors[TextFieldSel]: ki.Props{
		"background-color": &Prefs.Colors.Select,
	},
}

// signals that buttons can send
type TextFieldSignals int64

const (
	// main signal -- return was pressed and an edit was completed -- data is the text
	TextFieldDone TextFieldSignals = iota

	// some text was selected (for Inactive state, selection is via WidgetSig)
	TextFieldSelected

	TextFieldSignalsN
)

//go:generate stringer -type=TextFieldSignals

// TextFieldStates are mutually-exclusive textfield states -- determines appearance
type TextFieldStates int32

const (
	// normal state -- there but not being interacted with
	TextFieldActive TextFieldStates = iota

	// textfield is the focus -- will respond to keyboard input
	TextFieldFocus

	// inactive -- not editable
	TextFieldInactive

	// selected -- for inactive state, can select entire element
	TextFieldSel

	TextFieldStatesN
)

//go:generate stringer -type=TextFieldStates

// Style selector names for the different states
var TextFieldSelectors = []string{":active", ":focus", ":inactive", ":selected"}

// Text returns the current text -- applies any unapplied changes first
func (tf *TextField) Text() string {
	tf.EditDone()
	return tf.Txt
}

// SetText sets the text to be edited and reverts any current edit to reflect this new text
func (tf *TextField) SetText(txt string) {
	if tf.Txt == txt && !tf.Edited {
		return
	}
	tf.Txt = txt
	tf.RevertEdit()
}

// Label returns the display label for this node, satisfying the Labeler interface
func (tf *TextField) Label() string {
	if tf.Txt != "" {
		return tf.Txt
	}
	return tf.Nm
}

// EditDone completes editing and copies the active edited text to the text --
// called when the return key is pressed or goes out of focus
func (tf *TextField) EditDone() {
	if tf.Edited {
		tf.Edited = false
		tf.Txt = string(tf.EditTxt)
		tf.TextFieldSig.Emit(tf.This, int64(TextFieldDone), tf.Txt)
	}
	tf.ClearSelected()
}

// RevertEdit aborts editing and reverts to last saved text
func (tf *TextField) RevertEdit() {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	tf.EditTxt = []rune(tf.Txt)
	tf.Edited = false
	tf.StartPos = 0
	tf.EndPos = tf.CharWidth
	tf.SelectReset()
}

// CursorForward moves the cursor forward
func (tf *TextField) CursorForward(steps int) {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	tf.CursorPos += steps
	if tf.CursorPos > len(tf.EditTxt) {
		tf.CursorPos = len(tf.EditTxt)
	}
	if tf.CursorPos > tf.EndPos {
		inc := tf.CursorPos - tf.EndPos
		tf.EndPos += inc
	}
	if tf.SelectMode {
		if tf.CursorPos-steps < tf.SelectStart {
			tf.SelectStart = tf.CursorPos
		} else if tf.CursorPos > tf.SelectStart {
			tf.SelectEnd = tf.CursorPos
		} else {
			tf.SelectStart = tf.CursorPos
		}
		tf.SelectUpdate()
	}
}

// CursorForward moves the cursor backward
func (tf *TextField) CursorBackward(steps int) {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	tf.CursorPos -= steps
	if tf.CursorPos < 0 {
		tf.CursorPos = 0
	}
	if tf.CursorPos <= tf.StartPos {
		dec := kit.MinInt(tf.StartPos, 8)
		tf.StartPos -= dec
	}
	if tf.SelectMode {
		if tf.CursorPos+steps < tf.SelectStart {
			tf.SelectStart = tf.CursorPos
		} else if tf.CursorPos > tf.SelectStart {
			tf.SelectEnd = tf.CursorPos
		} else {
			tf.SelectStart = tf.CursorPos
		}
		tf.SelectUpdate()
	}
}

// CursorStart moves the cursor to the start of the text, updating selection
// if select mode is active
func (tf *TextField) CursorStart() {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	tf.CursorPos = 0
	tf.StartPos = 0
	tf.EndPos = kit.MinInt(len(tf.EditTxt), tf.StartPos+tf.CharWidth)
	if tf.SelectMode {
		tf.SelectStart = 0
		tf.SelectUpdate()
	}
}

// CursorEnd moves the cursor to the end of the text
func (tf *TextField) CursorEnd() {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	ed := len(tf.EditTxt)
	tf.CursorPos = ed
	tf.EndPos = len(tf.EditTxt) // try -- display will adjust
	tf.StartPos = kit.MaxInt(0, tf.EndPos-tf.CharWidth)
	if tf.SelectMode {
		tf.SelectEnd = ed
		tf.SelectUpdate()
	}
}

// todo: ctrl+backspace = delete word
// shift+arrow = select
// uparrow = start / down = end

// CursorBackspace deletes character(s) immediately before cursor
func (tf *TextField) CursorBackspace(steps int) {
	if tf.HasSelection() {
		tf.DeleteSelection()
		return
	}
	if tf.CursorPos < steps {
		steps = tf.CursorPos
	}
	if steps <= 0 {
		return
	}
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	tf.Edited = true
	tf.EditTxt = append(tf.EditTxt[:tf.CursorPos-steps], tf.EditTxt[tf.CursorPos:]...)
	tf.CursorBackward(steps)
	if tf.CursorPos > tf.SelectStart && tf.CursorPos <= tf.SelectEnd {
		tf.SelectEnd -= steps
	} else if tf.CursorPos < tf.SelectStart {
		tf.SelectStart -= steps
		tf.SelectEnd -= steps
	}
	tf.SelectUpdate()
}

// CursorDelete deletes character(s) immediately after the cursor
func (tf *TextField) CursorDelete(steps int) {
	if tf.HasSelection() {
		tf.DeleteSelection()
	}
	if tf.CursorPos+steps > len(tf.EditTxt) {
		steps = len(tf.EditTxt) - tf.CursorPos
	}
	if steps <= 0 {
		return
	}
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	tf.Edited = true
	tf.EditTxt = append(tf.EditTxt[:tf.CursorPos], tf.EditTxt[tf.CursorPos+steps:]...)
	if tf.CursorPos > tf.SelectStart && tf.CursorPos <= tf.SelectEnd {
		tf.SelectEnd -= steps
	} else if tf.CursorPos < tf.SelectStart {
		tf.SelectStart -= steps
		tf.SelectEnd -= steps
	}
	tf.SelectUpdate()
}

// CursorKill deletes text from cursor to end of text
func (tf *TextField) CursorKill() {
	steps := len(tf.EditTxt) - tf.CursorPos
	tf.CursorDelete(steps)
}

// ClearSelected resets both the global selected flag and any current selection
func (tf *TextField) ClearSelected() {
	tf.WidgetBase.ClearSelected()
	tf.SelectReset()
}

// HasSelection returns whether there is a selected region of text
func (tf *TextField) HasSelection() bool {
	tf.SelectUpdate()
	if tf.SelectStart < tf.SelectEnd {
		return true
	}
	return false
}

// Selection returns the currently selected text
func (tf *TextField) Selection() string {
	if tf.HasSelection() {
		return string(tf.EditTxt[tf.SelectStart:tf.SelectEnd])
	}
	return ""
}

// SelectModeToggle toggles the SelectMode, updating selection with cursor movement
func (tf *TextField) SelectModeToggle() {
	if tf.SelectMode {
		tf.SelectMode = false
	} else {
		tf.SelectMode = true
		tf.SelectStart = tf.CursorPos
		tf.SelectEnd = tf.SelectStart
	}
}

// SelectAll selects all the text
func (tf *TextField) SelectAll() {
	updt := tf.UpdateStart()
	tf.SelectStart = 0
	tf.SelectEnd = len(tf.EditTxt)
	tf.UpdateEnd(updt)
}

// IsWordBreak defines what counts as a word break for the purposes of selecting words
func (tf *TextField) IsWordBreak(r rune) bool {
	if unicode.IsSpace(r) || unicode.IsSymbol(r) || unicode.IsPunct(r) {
		return true
	}
	return false
}

// SelectWord selects the word (whitespace delimited) that the cursor is on
func (tf *TextField) SelectWord() {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	sz := len(tf.EditTxt)
	if sz <= 3 {
		tf.SelectAll()
		return
	}
	tf.SelectStart = tf.CursorPos
	if tf.SelectStart >= sz {
		tf.SelectStart = sz - 2
	}
	if !tf.IsWordBreak(tf.EditTxt[tf.SelectStart]) {
		for tf.SelectStart > 0 {
			if tf.IsWordBreak(tf.EditTxt[tf.SelectStart-1]) {
				break
			}
			tf.SelectStart--
		}
		tf.SelectEnd = tf.CursorPos + 1
		for tf.SelectEnd < sz {
			if tf.IsWordBreak(tf.EditTxt[tf.SelectEnd]) {
				break
			}
			tf.SelectEnd++
		}
	} else { // keep the space start -- go to next space..
		tf.SelectEnd = tf.CursorPos + 1
		for tf.SelectEnd < sz {
			if !tf.IsWordBreak(tf.EditTxt[tf.SelectEnd]) {
				break
			}
			tf.SelectEnd++
		}
		for tf.SelectEnd < sz {
			if tf.IsWordBreak(tf.EditTxt[tf.SelectEnd]) {
				break
			}
			tf.SelectEnd++
		}
	}
}

// SelectReset resets the selection
func (tf *TextField) SelectReset() {
	tf.SelectMode = false
	if tf.SelectStart == 0 && tf.SelectEnd == 0 {
		return
	}
	updt := tf.UpdateStart()
	tf.SelectStart = 0
	tf.SelectEnd = 0
	tf.UpdateEnd(updt)
}

// SelectUpdate updates the select region after any change to the text, to keep it in range
func (tf *TextField) SelectUpdate() {
	if tf.SelectStart < tf.SelectEnd {
		ed := len(tf.EditTxt)
		if tf.SelectStart < 0 {
			tf.SelectStart = 0
		}
		if tf.SelectEnd > ed {
			tf.SelectEnd = ed
		}
	} else {
		tf.SelectReset()
	}
}

// Cut cuts any selected text and adds it to the clipboard, also returns cut text
func (tf *TextField) Cut() string {
	cut := tf.DeleteSelection()
	if cut != "" {
		oswin.TheApp.ClipBoard().Write(mimedata.NewText(cut))
	}
	return cut
}

// DeleteSelection deletes any selected text, without adding to clipboard --
// returns text deleted
func (tf *TextField) DeleteSelection() string {
	tf.SelectUpdate()
	if !tf.HasSelection() {
		return ""
	}
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	cut := tf.Selection()
	tf.Edited = true
	tf.EditTxt = append(tf.EditTxt[:tf.SelectStart], tf.EditTxt[tf.SelectEnd:]...)
	if tf.CursorPos > tf.SelectStart {
		if tf.CursorPos < tf.SelectEnd {
			tf.CursorPos = tf.SelectStart
		} else {
			tf.CursorPos -= tf.SelectEnd - tf.SelectStart
		}
	}
	tf.SelectReset()
	return cut
}

// Copy copies any selected text to the clipboard, and returns that text,
// optionaly resetting the current selection
func (tf *TextField) Copy(reset bool) string {
	tf.SelectUpdate()
	if !tf.HasSelection() {
		return ""
	}
	cpy := tf.Selection()
	oswin.TheApp.ClipBoard().Write(mimedata.NewText(cpy))
	if reset {
		tf.SelectReset()
	}
	return cpy
}

// Paste inserts text from the clipboard at current cursor position -- if
// cursor is within a current selection, that selection is
func (tf *TextField) Paste() {
	data := oswin.TheApp.ClipBoard().Read([]string{mimedata.TextPlain})
	if data != nil {
		if tf.CursorPos >= tf.SelectStart && tf.CursorPos < tf.SelectEnd {
			tf.DeleteSelection()
		}
		tf.InsertAtCursor(data.Text(mimedata.TextPlain))
	}
}

// InsertAtCursor inserts given text at current cursor position
func (tf *TextField) InsertAtCursor(str string) {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	if tf.HasSelection() {
		tf.Cut()
	}
	tf.Edited = true
	rs := []rune(str)
	rsl := len(rs)
	nt := make([]rune, 0, cap(tf.EditTxt)+cap(rs))
	nt = append(nt, tf.EditTxt[:tf.CursorPos]...)
	nt = append(nt, rs...)
	nt = append(nt, tf.EditTxt[tf.CursorPos:]...)
	tf.EditTxt = nt
	tf.EndPos += rsl
	tf.CursorForward(rsl)
}

// cpos := tf.CharStartPos(tf.CursorPos).ToPoint()

func (tf *TextField) MakeContextMenu(m *Menu) {
	cpsc := ActiveKeyMap.ChordForFun(KeyFunCopy)
	ac := m.AddMenuText("Copy", cpsc, tf.This, nil, nil, func(recv, send ki.Ki, sig int64, data interface{}) {
		tff := recv.Embed(KiT_TextField).(*TextField)
		tff.Copy(true)
	})
	ac.SetActiveState(tf.HasSelection())
	if !tf.IsInactive() {
		ctsc := ActiveKeyMap.ChordForFun(KeyFunCut)
		ptsc := ActiveKeyMap.ChordForFun(KeyFunPaste)
		ac = m.AddMenuText("Cut", ctsc, tf.This, nil, nil, func(recv, send ki.Ki, sig int64, data interface{}) {
			tff := recv.Embed(KiT_TextField).(*TextField)
			tff.Cut()
		})
		ac.SetActiveState(tf.HasSelection())
		ac = m.AddMenuText("Paste", ptsc, tf.This, nil, nil, func(recv, send ki.Ki, sig int64, data interface{}) {
			tff := recv.Embed(KiT_TextField).(*TextField)
			tff.Paste()
		})
		ac.SetInactiveState(oswin.TheApp.ClipBoard().IsEmpty())
	}
}

// OfferCompletions pops up a menu of possible completions
func (tf *TextField) OfferCompletions() {
	win := tf.ParentWindow()
	if PopupIsCompleter(win.Popup) {
		win.ClosePopup(win.Popup)
	}
	if tf.Completion.Func == nil {
		return
	}

	tf.Completion.Completions, tf.Completion.Seed = tf.Completion.Func(string(tf.EditTxt[0:tf.CursorPos]))
	count := len(tf.Completion.Completions)
	if count > 0 {
		if count == 1 && tf.Completion.Completions[0] == tf.Completion.Seed {
			return
		}
		var m Menu
		for i := 0; i < count; i++ {
			s := tf.Completion.Completions[i]
			m.AddMenuText(s, "", tf.This, nil, nil, func(recv, send ki.Ki, sig int64, data interface{}) {
				tff := recv.Embed(KiT_TextField).(*TextField)
				tff.Complete(s)
			})
		}
		cpos := tf.CharStartPos(tf.CursorPos).ToPoint()
		// todo: figure popup placement using font and line height
		vp := PopupMenu(m, cpos.X+15, cpos.Y+50, tf.Viewport, "tf-completion-menu")
		bitflag.Set(&vp.Flag, int(VpFlagCompleter))
		vp.KnownChild(0).SetProp("no-focus-name", true) // disable name focusing -- grabs key events in popup instead of in textfield!
	}
}

// Complete edits the text field using the string chosen from the completion menu
func (tf *TextField) Complete(str string) {
	s1 := string(tf.EditTxt[0:tf.CursorPos])
	s2 := string(tf.EditTxt[tf.CursorPos:len(tf.EditTxt)])
	s1 = strings.TrimSuffix(s1, tf.Completion.Seed)
	s1 += str
	txt := s1 + s2
	tf.EditTxt = []rune(txt)
	tf.CursorForward(len(str) - len(tf.Completion.Seed))
}

// PixelToCursor finds the cursor position that corresponds to the given pixel location
func (tf *TextField) PixelToCursor(pixOff float32) int {
	st := &tf.Sty

	spc := st.BoxSpace()
	px := pixOff - spc

	if px <= 0 {
		return tf.StartPos
	}

	// for selection to work correctly, we need this to be deterministic

	sz := len(tf.EditTxt)
	c := tf.StartPos + int(float64(px/st.UnContext.ToDotsFactor(units.Ch)))
	c = kit.MinInt(c, sz)

	w := tf.TextWidth(tf.StartPos, c)
	if w > px {
		for w > px {
			c--
			if c <= tf.StartPos {
				c = tf.StartPos
				break
			}
			w = tf.TextWidth(tf.StartPos, c)
		}
	} else if w < px {
		for c < tf.EndPos {
			wn := tf.TextWidth(tf.StartPos, c+1)
			if wn > px {
				break
			} else if wn == px {
				c++
				break
			}
			c++
		}
	}
	return c
}

func (tf *TextField) SetCursorFromPixel(pixOff float32, selMode mouse.SelectModes) {
	updt := tf.UpdateStart()
	defer tf.UpdateEnd(updt)
	oldPos := tf.CursorPos
	tf.CursorPos = tf.PixelToCursor(pixOff)
	if tf.SelectMode || selMode != mouse.NoSelectMode {
		if !tf.SelectMode && selMode != mouse.NoSelectMode {
			tf.SelectStart = oldPos
			tf.SelectMode = true
		}
		if !tf.IsDragging() && tf.CursorPos >= tf.SelectStart && tf.CursorPos < tf.SelectEnd {
			tf.SelectReset()
		} else if tf.CursorPos > tf.SelectStart {
			tf.SelectEnd = tf.CursorPos
		} else {
			tf.SelectStart = tf.CursorPos
		}
		tf.SelectUpdate()
	} else if tf.HasSelection() {
		tf.SelectReset()
	}
}

// KeyInput handles keyboard input into the text field and from the completion menu
func (tf *TextField) KeyInput(kt *key.ChordEvent) {
	kf := KeyFun(kt.ChordString())
	win := tf.ParentWindow()

	if PopupIsCompleter(win.Popup) {
		switch kf {
		case KeyFunFocusNext: // tab will complete if single item or try to extend if multiple items
			count := len(tf.Completion.Completions)
			if count > 0 {
				if count == 1 { // just complete
					tf.Complete(tf.Completion.Completions[0])
					win.ClosePopup(win.Popup)
				} else { // try to extend the seed
					s := complete.ExtendSeed(tf.Completion.Completions, tf.Completion.Seed)
					if s != "" {
						win.ClosePopup(win.Popup)
						tf.InsertAtCursor(s)
						tf.OfferCompletions()
					}
				}
			}
			kt.SetProcessed()
		default:
			//fmt.Printf("some char\n")
		}
	}

	if kt.IsProcessed() {
		return
	}

	// first all the keys that work for both inactive and active
	switch kf {
	case KeyFunMoveRight:
		kt.SetProcessed()
		tf.CursorForward(1)
		tf.OfferCompletions()
	case KeyFunMoveLeft:
		kt.SetProcessed()
		tf.CursorBackward(1)
		tf.OfferCompletions()
	case KeyFunHome:
		kt.SetProcessed()
		tf.CursorStart()
	case KeyFunEnd:
		kt.SetProcessed()
		tf.CursorEnd()
	case KeyFunSelectMode:
		kt.SetProcessed()
		tf.SelectModeToggle()
	case KeyFunCancelSelect:
		kt.SetProcessed()
		tf.SelectReset()
	case KeyFunSelectAll:
		kt.SetProcessed()
		tf.SelectAll()
	case KeyFunCopy:
		kt.SetProcessed()
		tf.Copy(true) // reset
	}
	if tf.IsInactive() || kt.IsProcessed() {
		return
	}
	switch kf {
	case KeyFunSelectItem: // enter
		fallthrough
	case KeyFunAccept: // ctrl+enter
		tf.EditDone()
		kt.SetProcessed()
		tf.FocusNext()
	case KeyFunAbort: // esc
		tf.RevertEdit()
		kt.SetProcessed()
		tf.FocusNext()
	case KeyFunBackspace:
		kt.SetProcessed()
		tf.CursorBackspace(1)
		tf.OfferCompletions()
	case KeyFunKill:
		kt.SetProcessed()
		tf.CursorKill()
	case KeyFunDelete:
		kt.SetProcessed()
		tf.CursorDelete(1)
	case KeyFunCut:
		kt.SetProcessed()
		tf.Cut()
	case KeyFunPaste:
		kt.SetProcessed()
		tf.Paste()
	case KeyFunComplete:
		kt.SetProcessed()
		tf.OfferCompletions()
	case KeyFunNil:
		if unicode.IsPrint(kt.Rune) {
			if !kt.HasAnyModifier(key.Control, key.Meta) {
				kt.SetProcessed()
				tf.InsertAtCursor(string(kt.Rune))
				tf.OfferCompletions()
			}
		}
	}
}

// MouseEvent handles the mouse.Event
func (tf *TextField) MouseEvent(me *mouse.Event) {
	if !tf.IsInactive() && !tf.HasFocus() {
		tf.GrabFocus()
	}
	me.SetProcessed()
	switch me.Button {
	case mouse.Left:
		if me.Action == mouse.Press {
			if tf.IsInactive() {
				tf.SetSelectedState(!tf.IsSelected())
				tf.EmitSelectedSignal()
				tf.UpdateSig()
			} else {
				pt := tf.PointToRelPos(me.Pos())
				tf.SetCursorFromPixel(float32(pt.X), me.SelectMode())
			}
		} else if me.Action == mouse.DoubleClick {
			me.SetProcessed()
			if tf.HasSelection() {
				if tf.SelectStart == 0 && tf.SelectEnd == len(tf.EditTxt) {
					tf.SelectReset()
				} else {
					tf.SelectAll()
				}
			} else {
				tf.SelectWord()
			}
		}
	case mouse.Middle:
		if !tf.IsInactive() && me.Action == mouse.Press {
			me.SetProcessed()
			pt := tf.PointToRelPos(me.Pos())
			tf.SetCursorFromPixel(float32(pt.X), me.SelectMode())
			tf.Paste()
		}
	case mouse.Right:
		if me.Action == mouse.Press {
			me.SetProcessed()
			tf.EmitContextMenuSignal()
			tf.This.(Node2D).ContextMenu()
		}
	}
}

func (tf *TextField) TextFieldEvents() {
	tf.HoverTooltipEvent()
	tf.ConnectEvent(oswin.MouseDragEvent, RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		me := d.(*mouse.DragEvent)
		me.SetProcessed()
		tf := recv.Embed(KiT_TextField).(*TextField)
		if !tf.SelectMode {
			tf.SelectModeToggle()
		}
		pt := tf.PointToRelPos(me.Pos())
		tf.SetCursorFromPixel(float32(pt.X), mouse.NoSelectMode)
	})
	tf.ConnectEvent(oswin.MouseEvent, RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		tff := recv.Embed(KiT_TextField).(*TextField)
		me := d.(*mouse.Event)
		tff.MouseEvent(me)
	})
	tf.ConnectEvent(oswin.MouseFocusEvent, RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		tff := recv.Embed(KiT_TextField).(*TextField)
		if tff.IsInactive() {
			return
		}
		me := d.(*mouse.FocusEvent)
		me.SetProcessed()
		if me.Action == mouse.Enter {
			oswin.TheApp.Cursor().PushIfNot(cursor.IBeam)
		} else {
			oswin.TheApp.Cursor().PopIf(cursor.IBeam)
		}
	})
	tf.ConnectEvent(oswin.KeyChordEvent, RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		tff := recv.Embed(KiT_TextField).(*TextField)
		kt := d.(*key.ChordEvent)
		tff.KeyInput(kt)
	})
	if dlg, ok := tf.Viewport.This.(*Dialog); ok {
		dlg.DialogSig.Connect(tf.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			tff, _ := recv.Embed(KiT_TextField).(*TextField)
			if sig == int64(DialogAccepted) {
				tff.EditDone()
			}
		})
	}
}

////////////////////////////////////////////////////
//  Node2D Interface

func (tf *TextField) Init2D() {
	tf.Init2DWidget()
	tf.EditTxt = []rune(tf.Txt)
	tf.Edited = false
}

func (tf *TextField) Style2D() {
	tf.SetCanFocusIfActive()
	tf.Style2DWidget()
	pst := &(tf.Par.(Node2D).AsWidget().Sty)
	for i := 0; i < int(TextFieldStatesN); i++ {
		tf.StateStyles[i].CopyFrom(&tf.Sty)
		tf.StateStyles[i].SetStyleProps(pst, tf.StyleProps(TextFieldSelectors[i]))
		tf.StateStyles[i].StyleCSS(tf.This.(Node2D), tf.CSSAgg, TextFieldSelectors[i])
		tf.StateStyles[i].CopyUnitContext(&tf.Sty.UnContext)
	}
}

func (tf *TextField) UpdateRenderAll() bool {
	st := &tf.Sty
	st.Font.LoadFont(&st.UnContext)
	tf.RenderAll.SetRunes(tf.EditTxt, &st.Font, &st.UnContext, &st.Text, true, 0, 0)
	return true
}

func (tf *TextField) Size2D(iter int) {
	tmptxt := tf.EditTxt
	if len(tf.Txt) == 0 && len(tf.Placeholder) > 0 {
		tf.EditTxt = []rune(tf.Placeholder)
	} else {
		tf.EditTxt = []rune(tf.Txt)
	}
	tf.Edited = false
	tf.StartPos = 0
	maxlen := tf.MaxWidthReq
	if maxlen <= 0 {
		maxlen = 50
	}
	tf.EndPos = kit.MinInt(len(tf.EditTxt), maxlen)
	tf.UpdateRenderAll()
	tf.FontHeight = tf.RenderAll.Size.Y
	w := tf.TextWidth(tf.StartPos, tf.EndPos)
	w += 2.0 // give some extra buffer
	// fmt.Printf("fontheight: %v width: %v\n", tf.FontHeight, w)
	tf.Size2DFromWH(w, tf.FontHeight)
	tf.EditTxt = tmptxt
}

func (tf *TextField) Layout2D(parBBox image.Rectangle, iter int) bool {
	tf.Layout2DBase(parBBox, true, iter) // init style
	for i := 0; i < int(TextFieldStatesN); i++ {
		tf.StateStyles[i].CopyUnitContext(&tf.Sty.UnContext)
	}
	return tf.Layout2DChildren(iter)
}

// StartCharPos returns the starting position of the given rune
func (tf *TextField) StartCharPos(idx int) float32 {
	if idx <= 0 || len(tf.RenderAll.Spans) != 1 {
		return 0.0
	}
	sr := &(tf.RenderAll.Spans[0])
	sz := len(sr.Render)
	if sz == 0 {
		return 0.0
	}
	if idx >= sz {
		return sr.LastPos.X
	}
	return sr.Render[idx].RelPos.X
}

// TextWidth returns the text width in dots between the two text string
// positions (ed is exclusive -- +1 beyond actual char)
func (tf *TextField) TextWidth(st, ed int) float32 {
	return tf.StartCharPos(ed) - tf.StartCharPos(st)
}

// CharStartPos returns the starting render coords for the given character
// position in string -- makes no attempt to rationalize that pos (i.e., if
// not in visible range, position will be out of range too)
func (tf *TextField) CharStartPos(charidx int) Vec2D {
	st := &tf.Sty
	spc := st.BoxSpace()
	pos := tf.LayData.AllocPos.AddVal(spc)
	cpos := tf.TextWidth(tf.StartPos, charidx)
	return Vec2D{pos.X + cpos, pos.Y}
}

func (tf *TextField) RenderCursor() {
	cpos := tf.CharStartPos(tf.CursorPos)
	rs := &tf.Viewport.Render
	pc := &rs.Paint
	pc.DrawLine(rs, cpos.X, cpos.Y, cpos.X, cpos.Y+tf.FontHeight)
	pc.Stroke(rs)
}

func (tf *TextField) RenderSelect() {
	if tf.SelectEnd <= tf.SelectStart {
		return
	}
	effst := kit.MaxInt(tf.StartPos, tf.SelectStart)
	if effst >= tf.EndPos {
		return
	}
	effed := kit.MinInt(tf.EndPos, tf.SelectEnd)
	if effed < tf.StartPos {
		return
	}
	if effed <= effst {
		return
	}

	spos := tf.CharStartPos(effst)

	rs := &tf.Viewport.Render
	pc := &rs.Paint
	st := &tf.StateStyles[TextFieldSel]
	tsz := tf.TextWidth(effst, effed)
	pc.FillBox(rs, spos, Vec2D{tsz, tf.FontHeight}, &st.Font.BgColor)
}

// AutoScroll scrolls the starting position to keep the cursor visible
func (tf *TextField) AutoScroll() {
	st := &tf.Sty

	tf.UpdateRenderAll()

	sz := len(tf.EditTxt)

	if sz == 0 || tf.LayData.AllocSize.X <= 0 {
		tf.CursorPos = 0
		tf.EndPos = 0
		tf.StartPos = 0
		return
	}
	spc := st.BoxSpace()
	maxw := tf.LayData.AllocSize.X - 2.0*spc
	tf.CharWidth = int(maxw / st.UnContext.ToDotsFactor(units.Ch)) // rough guess in chars

	// first rationalize all the values
	if tf.EndPos == 0 || tf.EndPos > sz { // not init
		tf.EndPos = sz
	}
	if tf.StartPos >= tf.EndPos {
		tf.StartPos = kit.MaxInt(0, tf.EndPos-tf.CharWidth)
	}
	tf.CursorPos = InRangeInt(tf.CursorPos, 0, sz)

	inc := int(math32.Ceil(.1 * float32(tf.CharWidth)))
	inc = kit.MaxInt(4, inc)

	// keep cursor in view with buffer
	startIsAnchor := true
	if tf.CursorPos < (tf.StartPos + inc) {
		tf.StartPos -= inc
		tf.StartPos = kit.MaxInt(tf.StartPos, 0)
		tf.EndPos = tf.StartPos + tf.CharWidth
		tf.EndPos = kit.MinInt(sz, tf.EndPos)
	} else if tf.CursorPos > (tf.EndPos - inc) {
		tf.EndPos += inc
		tf.EndPos = kit.MinInt(tf.EndPos, sz)
		tf.StartPos = tf.EndPos - tf.CharWidth
		tf.StartPos = kit.MaxInt(0, tf.StartPos)
		startIsAnchor = false
	}

	if startIsAnchor {
		gotWidth := false
		spos := tf.StartCharPos(tf.StartPos)
		for {
			w := tf.StartCharPos(tf.EndPos) - spos
			if w < maxw {
				if tf.EndPos == sz {
					break
				}
				nw := tf.StartCharPos(tf.EndPos+1) - spos
				if nw >= maxw {
					gotWidth = true
					break
				}
				tf.EndPos++
			} else {
				tf.EndPos--
			}
		}
		if gotWidth || tf.StartPos == 0 {
			return
		}
		// otherwise, try getting some more chars by moving up start..
	}

	// end is now anchor
	epos := tf.StartCharPos(tf.EndPos)
	for {
		w := epos - tf.StartCharPos(tf.StartPos)
		if w < maxw {
			if tf.StartPos == 0 {
				break
			}
			nw := epos - tf.StartCharPos(tf.StartPos-1)
			if nw >= maxw {
				break
			}
			tf.StartPos--
		} else {
			tf.StartPos++
		}
	}
}

func (tf *TextField) Render2D() {
	if tf.FullReRenderIfNeeded() {
		return
	}
	if tf.PushBounds() {
		tf.TextFieldEvents()
		tf.AutoScroll() // inits paint with our style
		if tf.IsInactive() {
			if tf.IsSelected() {
				tf.Sty = tf.StateStyles[TextFieldSel]
			} else {
				tf.Sty = tf.StateStyles[TextFieldInactive]
			}
		} else if tf.HasFocus() {
			if tf.FocusActive {
				tf.Sty = tf.StateStyles[TextFieldFocus]
			} else {
				tf.Sty = tf.StateStyles[TextFieldActive]
			}
		} else if tf.IsSelected() {
			tf.Sty = tf.StateStyles[TextFieldSel]
		} else {
			tf.Sty = tf.StateStyles[TextFieldActive]
		}
		rs := &tf.Viewport.Render
		st := &tf.Sty
		st.Font.LoadFont(&st.UnContext)
		tf.RenderStdBox(st)
		cur := tf.EditTxt[tf.StartPos:tf.EndPos]
		tf.RenderSelect()
		pos := tf.LayData.AllocPos.AddVal(st.BoxSpace())
		if len(tf.EditTxt) == 0 && len(tf.Placeholder) > 0 {
			st.Font.Color = st.Font.Color.Highlight(50)
			tf.RenderVis.SetString(tf.Placeholder, &st.Font, &st.UnContext, &st.Text, true, 0, 0)
			tf.RenderVis.RenderTopPos(rs, pos)

		} else {
			tf.RenderVis.SetRunes(cur, &st.Font, &st.UnContext, &st.Text, true, 0, 0)
			tf.RenderVis.RenderTopPos(rs, pos)
		}
		if tf.HasFocus() {
			tf.RenderCursor()
		}
		tf.Render2DChildren()
		tf.PopBounds()
	} else {
		tf.DisconnectAllEvents(RegPri)
	}
}

func (tf *TextField) FocusChanged2D(change FocusChanges) {
	switch change {
	case FocusLost:
		tf.FocusActive = false
		tf.EditDone()
		tf.UpdateSig()
	case FocusGot:
		tf.FocusActive = true
		tf.ScrollToMe()
		tf.CursorEnd()
		tf.EmitFocusedSignal()
		tf.UpdateSig()
	case FocusInactive:
		tf.FocusActive = false
		tf.EditDone()
		tf.UpdateSig()
	case FocusActive:
		tf.FocusActive = true
		tf.ScrollToMe()
		// tf.UpdateSig()
		// todo: see about cursor
	}
}

func (tf *TextField) SetCompleter(data interface{}, fun complete.Func) {
	if fun == nil {
		return
	}
	tf.Completion.Context = data
	tf.Completion.Func = fun
}
