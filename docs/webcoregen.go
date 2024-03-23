// Code generated by "core generate -webcore content"; DO NOT EDIT.

package main

import (
	"errors"
	"fmt"
	"maps"
	"strings"
	"time"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/colors/gradient"
	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keyfun"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
	"cogentcore.org/core/webcore"
)

func init() {
	maps.Copy(webcore.Examples, WebcoreExamples)
}

// WebcoreExamples are the compiled webcore examples for this app.
var WebcoreExamples = map[string]func(parent gi.Widget){
	"getting-started/hello-world-0": func(parent gi.Widget) {
		b := parent
		gi.NewButton(b).SetText("Hello, World!")
	},
	"basics/widgets-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").SetIcon(icons.Add)
	},
	"basics/events-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Button clicked")
		})
	},
	"basics/events-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprint("Button clicked at ", e.Pos()))
		})
	},
	"basics/styling-0": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Bold text").Style(func(s *styles.Style) {
			s.Font.Weight = styles.WeightBold
		})
	},
	"basics/styling-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Success button").Style(func(s *styles.Style) {
			s.Background = colors.C(colors.Scheme.Success.Base)
			s.Color = colors.C(colors.Scheme.Success.On)
		})
	},
	"basics/styling-2": func(parent gi.Widget) {
		gi.NewBox(parent).Style(func(s *styles.Style) {
			s.Min.Set(units.Dp(50))
			s.Background = colors.C(colors.Scheme.Primary.Base)
		})
	},
	"widgets/buttons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Download")
	},
	"widgets/buttons-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetIcon(icons.Download)
	},
	"widgets/buttons-2": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Download").SetIcon(icons.Download)
	},
	"widgets/buttons-3": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-4": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Share").SetIcon(icons.Share).SetMenu(func(m *gi.Scene) {
			gi.NewButton(m).SetText("Copy link")
			gi.NewButton(m).SetText("Send message")
		})
	},
	"widgets/buttons-5": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Save").SetShortcut("Command+S").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "File saved")
		})
	},
	"widgets/buttons-6": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Open").SetKey(keyfun.Open).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "File opened")
		})
	},
	"widgets/buttons-7": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonFilled).SetText("Filled")
	},
	"widgets/buttons-8": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonTonal).SetText("Tonal")
	},
	"widgets/buttons-9": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonElevated).SetText("Elevated")
	},
	"widgets/buttons-10": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonOutlined).SetText("Outlined")
	},
	"widgets/buttons-11": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonText).SetText("Text")
	},
	"widgets/buttons-12": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonAction).SetText("Action")
	},
	"widgets/choosers-0": func(parent gi.Widget) {
		gi.NewChooser(parent).SetStrings("macOS", "Windows", "Linux")
	},
	"widgets/choosers-1": func(parent gi.Widget) {
		gi.NewChooser(parent).SetItems(
			gi.ChooserItem{Value: "Computer", Icon: icons.Computer, Tooltip: "Use a computer"},
			gi.ChooserItem{Value: "Phone", Icon: icons.Smartphone, Tooltip: "Use a phone"},
		)
	},
	"widgets/choosers-2": func(parent gi.Widget) {
		gi.NewChooser(parent).SetPlaceholder("Choose a platform").SetStrings("macOS", "Windows", "Linux")
	},
	"widgets/choosers-3": func(parent gi.Widget) {
		gi.NewChooser(parent).SetStrings("Apple", "Orange", "Strawberry").SetCurrentValue("Orange")
	},
	"widgets/choosers-4": func(parent gi.Widget) {
		gi.NewChooser(parent).SetType(gi.ChooserOutlined).SetStrings("Apple", "Orange", "Strawberry")
	},
	"widgets/choosers-5": func(parent gi.Widget) {
		gi.NewChooser(parent).SetIcon(icons.Sort).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-6": func(parent gi.Widget) {
		gi.NewChooser(parent).SetEditable(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-7": func(parent gi.Widget) {
		gi.NewChooser(parent).SetAllowNew(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-8": func(parent gi.Widget) {
		gi.NewChooser(parent).SetEditable(true).SetAllowNew(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-9": func(parent gi.Widget) {
		ch := gi.NewChooser(parent).SetStrings("Newest", "Oldest", "Popular")
		ch.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Sorting by %v", ch.CurrentItem.Value))
		})
	},
	"widgets/dialogs-0": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Message")
		bt.OnClick(func(e events.Event) {
			gi.MessageDialog(bt, "Something happened", "Message")
		})
	},
	"widgets/dialogs-1": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Error")
		bt.OnClick(func(e events.Event) {
			gi.ErrorDialog(bt, errors.New("invalid encoding format"), "Error loading file")
		})
	},
	"widgets/dialogs-2": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Confirm")
		bt.OnClick(func(e events.Event) {
			d := gi.NewBody().AddTitle("Confirm").AddText("Send message?")
			d.AddBottomBar(func(pw gi.Widget) {
				d.AddCancel(pw).OnClick(func(e events.Event) {
					gi.MessageSnackbar(bt, "Dialog canceled")
				})
				d.AddOk(pw).OnClick(func(e events.Event) {
					gi.MessageSnackbar(bt, "Dialog accepted")
				})
			})
			d.NewDialog(bt).Run()
		})
	},
	"widgets/frames-0": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-1": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Background = colors.C(colors.Scheme.Warn.Container)
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-2": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Background = gradient.NewLinear().AddStop(colors.Yellow, 0).AddStop(colors.Orange, 0.5).AddStop(colors.Red, 1)
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-3": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.C(colors.Scheme.Outline))
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-4": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Border.Radius = styles.BorderRadiusLarge
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.C(colors.Scheme.Outline))
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-5": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Grow.Set(0, 0)
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.C(colors.Scheme.Outline))
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/icons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetIcon(icons.Send)
	},
	"widgets/icons-1": func(parent gi.Widget) {
		gi.NewIcon(parent).SetIcon(icons.Home)
	},
	"widgets/icons-2": func(parent gi.Widget) {
		gi.NewButton(parent).SetIcon(icons.Home.Fill())
	},
	"widgets/labels-0": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Hello, world!")
	},
	"widgets/labels-1": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("This is a very long sentence that demonstrates how label content will overflow onto multiple lines when the size of the label text exceeds the size of its surrounding container; labels are a customizable widget that Cogent Core provides, allowing you to display many kinds of text")
	},
	"widgets/labels-2": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText(`<b>You</b> can use <i>HTML</i> <u>formatting</u> inside of <b><i><u>Cogent Core</u></i></b> labels, including <span style="color:red;background-color:yellow">custom styling</span> and <a href="https://example.com">links</a>`)
	},
	"widgets/labels-3": func(parent gi.Widget) {
		gi.NewLabel(parent).SetType(gi.LabelHeadlineMedium).SetText("Hello, world!")
	},
	"widgets/labels-4": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Hello,\n\tworld!").Style(func(s *styles.Style) {
			s.Font.Size.Dp(21)
			s.Font.Style = styles.Italic
			s.Text.WhiteSpace = styles.WhiteSpacePre
			s.Color = colors.C(colors.Scheme.Success.Base)
			s.Font.Family = string(gi.AppearanceSettings.MonoFont)
		})
	},
	"widgets/layouts-0": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-1": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-2": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Gap.Set(units.Em(2))
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-3": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Max.X.Em(10)
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-4": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Overflow.X = styles.OverflowAuto
			s.Max.X.Em(10)
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-5": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Wrap = true
			s.Max.X.Em(10)
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-6": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Display = styles.Grid
			s.Columns = 2
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
		gi.NewButton(ly).SetText("Fourth")
	},
	"widgets/meters-0": func(parent gi.Widget) {
		gi.NewMeter(parent)
	},
	"widgets/meters-1": func(parent gi.Widget) {
		gi.NewMeter(parent).SetValue(0.7)
	},
	"widgets/meters-2": func(parent gi.Widget) {
		gi.NewMeter(parent).SetMin(5.7).SetMax(18).SetValue(10.2)
	},
	"widgets/meters-3": func(parent gi.Widget) {
		gi.NewMeter(parent).Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
	},
	"widgets/meters-4": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterCircle)
	},
	"widgets/meters-5": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterSemicircle)
	},
	"widgets/meters-6": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterCircle).SetText("50%")
	},
	"widgets/meters-7": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterSemicircle).SetText("50%")
	},
	"widgets/sliders-0": func(parent gi.Widget) {
		gi.NewSlider(parent)
	},
	"widgets/sliders-1": func(parent gi.Widget) {
		gi.NewSlider(parent).SetValue(0.7)
	},
	"widgets/sliders-2": func(parent gi.Widget) {
		gi.NewSlider(parent).SetMin(5.7).SetMax(18).SetValue(10.2)
	},
	"widgets/sliders-3": func(parent gi.Widget) {
		gi.NewSlider(parent).SetStep(0.2)
	},
	"widgets/sliders-4": func(parent gi.Widget) {
		gi.NewSlider(parent).SetStep(0.2).SetEnforceStep(true)
	},
	"widgets/sliders-5": func(parent gi.Widget) {
		gi.NewSlider(parent).SetIcon(icons.DeployedCode.Fill())
	},
	"widgets/sliders-6": func(parent gi.Widget) {
		sr := gi.NewSlider(parent)
		sr.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("OnChange: %v", sr.Value))
		})
	},
	"widgets/sliders-7": func(parent gi.Widget) {
		sr := gi.NewSlider(parent)
		sr.OnInput(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("OnInput: %v", sr.Value))
		})
	},
	"widgets/snackbars-0": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Message")
		bt.OnClick(func(e events.Event) {
			gi.MessageSnackbar(bt, "New messages loaded")
		})
	},
	"widgets/snackbars-1": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Error")
		bt.OnClick(func(e events.Event) {
			gi.ErrorSnackbar(bt, errors.New("file not found"), "Error loading page")
		})
	},
	"widgets/snackbars-2": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Custom")
		bt.OnClick(func(e events.Event) {
			gi.NewBody().AddSnackbarText("Files updated").
				AddSnackbarButton("Refresh", func(e events.Event) {
					gi.MessageSnackbar(bt, "Refreshed files")
				}).AddSnackbarIcon(icons.Close).NewSnackbar(bt).Run()
		})
	},
	"widgets/spinners-0": func(parent gi.Widget) {
		gi.NewSpinner(parent)
	},
	"widgets/spinners-1": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetValue(12.7)
	},
	"widgets/spinners-2": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetMin(-0.5).SetMax(2.7)
	},
	"widgets/spinners-3": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetStep(6)
	},
	"widgets/spinners-4": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetStep(4).SetEnforceStep(true)
	},
	"widgets/spinners-5": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/spinners-6": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetFormat("%X").SetStep(1).SetValue(44)
	},
	"widgets/spinners-7": func(parent gi.Widget) {
		sp := gi.NewSpinner(parent)
		sp.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Value changed to %g", sp.Value))
		})
	},
	"widgets/splits-0": func(parent gi.Widget) {
		sp := gi.NewSplits(parent)
		gi.NewLabel(sp).SetText("First")
		gi.NewLabel(sp).SetText("Second")
	},
	"widgets/splits-1": func(parent gi.Widget) {
		sp := gi.NewSplits(parent)
		gi.NewLabel(sp).SetText("First")
		gi.NewLabel(sp).SetText("Second")
		gi.NewLabel(sp).SetText("Third")
		gi.NewLabel(sp).SetText("Fourth")
	},
	"widgets/switches-0": func(parent gi.Widget) {
		gi.NewSwitch(parent)
	},
	"widgets/switches-1": func(parent gi.Widget) {
		gi.NewSwitch(parent).SetText("Remember me")
	},
	"widgets/switches-2": func(parent gi.Widget) {
		gi.NewSwitch(parent).SetType(gi.SwitchCheckbox).SetText("Remember me")
	},
	"widgets/switches-3": func(parent gi.Widget) {
		gi.NewSwitch(parent).SetType(gi.SwitchRadioButton).SetText("Remember me")
	},
	"widgets/switches-4": func(parent gi.Widget) {
		sw := gi.NewSwitch(parent).SetText("Remember me")
		sw.OnChange(func(e events.Event) {
			gi.MessageSnackbar(sw, fmt.Sprintf("Switch is %v", sw.IsChecked()))
		})
	},
	"widgets/switches-5": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-6": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetItems(
			gi.SwitchItem{Label: "Go", Tooltip: "Elegant, fast, and easy-to-use"},
			gi.SwitchItem{Label: "Python", Tooltip: "Slow and duck-typed"},
			gi.SwitchItem{Label: "C++", Tooltip: "Hard to use and slow to compile"},
		)
	},
	"widgets/switches-7": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetMutex(true).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-8": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchChip).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-9": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchCheckbox).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-10": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchRadioButton).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-11": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchSegmentedButton).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-12": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetStrings("Go", "Python", "C++").Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
	},
	"widgets/switches-13": func(parent gi.Widget) {
		sw := gi.NewSwitches(parent).SetStrings("Go", "Python", "C++")
		sw.OnChange(func(e events.Event) {
			gi.MessageSnackbar(sw, fmt.Sprintf("Currently selected: %v", sw.SelectedItems()))
		})
	},
	"widgets/tabs-0": func(parent gi.Widget) {
		ts := gi.NewTabs(parent)
		ts.NewTab("First")
		ts.NewTab("Second")
	},
	"widgets/tabs-1": func(parent gi.Widget) {
		ts := gi.NewTabs(parent)
		first := ts.NewTab("First")
		gi.NewLabel(first).SetText("I am first!")
		second := ts.NewTab("Second")
		gi.NewLabel(second).SetText("I am second!")
	},
	"widgets/tabs-2": func(parent gi.Widget) {
		ts := gi.NewTabs(parent)
		ts.NewTab("First")
		ts.NewTab("Second")
		ts.NewTab("Third")
		ts.NewTab("Fourth")
	},
	"widgets/tabs-3": func(parent gi.Widget) {
		ts := gi.NewTabs(parent)
		ts.NewTab("First", icons.Home)
		ts.NewTab("Second", icons.Explore)
	},
	"widgets/tabs-4": func(parent gi.Widget) {
		ts := gi.NewTabs(parent).SetType(gi.FunctionalTabs)
		ts.NewTab("First")
		ts.NewTab("Second")
		ts.NewTab("Third")
	},
	"widgets/tabs-5": func(parent gi.Widget) {
		ts := gi.NewTabs(parent).SetType(gi.NavigationAuto)
		ts.NewTab("First", icons.Home)
		ts.NewTab("Second", icons.Explore)
		ts.NewTab("Third", icons.History)
	},
	"widgets/tabs-6": func(parent gi.Widget) {
		ts := gi.NewTabs(parent).SetNewTabButton(true)
		ts.NewTab("First")
		ts.NewTab("Second")
	},
	"widgets/text-fields-0": func(parent gi.Widget) {
		gi.NewTextField(parent)
	},
	"widgets/text-fields-1": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Name:")
		gi.NewTextField(parent).SetPlaceholder("Jane Doe")
	},
	"widgets/text-fields-2": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("Hello, world!")
	},
	"widgets/text-fields-3": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("This is a long sentence that demonstrates how text field content can overflow onto multiple lines")
	},
	"widgets/text-fields-4": func(parent gi.Widget) {
		gi.NewTextField(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/text-fields-5": func(parent gi.Widget) {
		gi.NewTextField(parent).SetTypePassword()
	},
	"widgets/text-fields-6": func(parent gi.Widget) {
		gi.NewTextField(parent).AddClearButton()
	},
	"widgets/text-fields-7": func(parent gi.Widget) {
		gi.NewTextField(parent).SetLeadingIcon(icons.Euro).SetTrailingIcon(icons.OpenInNew, func(e events.Event) {
			gi.MessageSnackbar(parent, "Opening shopping cart")
		})
	},
	"widgets/text-fields-8": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.SetValidator(func() error {
			if !strings.Contains(tf.Text(), "Go") {
				return errors.New("Must contain Go")
			}
			return nil
		})
	},
	"widgets/text-fields-9": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, "OnChange: "+tf.Text())
		})
	},
	"widgets/text-fields-10": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.OnInput(func(e events.Event) {
			gi.MessageSnackbar(parent, "OnInput: "+tf.Text())
		})
	},
	"widgets/tooltips-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetIcon(icons.Add).SetTooltip("Add a new item to the list")
	},
	"widgets/tooltips-1": func(parent gi.Widget) {
		gi.NewSlider(parent)
	},
	"views/values-0": func(parent gi.Widget) {
		giv.NewValue(parent, colors.Orange)
	},
	"views/values-1": func(parent gi.Widget) {
		t := time.Now()
		giv.NewValue(parent, &t).OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, "The time is "+t.Format(time.DateTime))
		})
	},
	"views/values-2": func(parent gi.Widget) {
		giv.NewValue(parent, 70, `view:"slider"`)
	},
	"views/map-views-0": func(parent gi.Widget) {
		giv.NewMapView(parent).SetMap(&map[string]int{"Go": 1, "C++": 3, "Python": 5})
	},
	"views/map-views-1": func(parent gi.Widget) {
		m := map[string]int{"Go": 1, "C++": 3, "Python": 5}
		giv.NewMapView(parent).SetMap(&m).OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Map: %v", m))
		})
	},
	"views/map-views-2": func(parent gi.Widget) {
		giv.NewMapView(parent).SetMap(&map[string]int{"Go": 1, "C++": 3, "Python": 5}).SetReadOnly(true)
	},
	"views/map-views-3": func(parent gi.Widget) {
		giv.NewMapViewInline(parent).SetMap(&map[string]int{"Go": 1, "C++": 3})
	},
	"views/map-views-4": func(parent gi.Widget) {
		giv.NewValue(parent, &map[string]int{"Go": 1, "C++": 3})
	},
	"views/map-views-5": func(parent gi.Widget) {
		giv.NewValue(parent, &map[string]int{"Go": 1, "C++": 3, "Python": 5})
	},
	"views/slice-views-0": func(parent gi.Widget) {
		giv.NewSliceView(parent).SetSlice(&[]int{1, 3, 5})
	},
	"views/slice-views-1": func(parent gi.Widget) {
		sl := []int{1, 3, 5}
		giv.NewSliceView(parent).SetSlice(&sl).OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Slice: %v", sl))
		})
	},
	"views/slice-views-2": func(parent gi.Widget) {
		giv.NewSliceView(parent).SetSlice(&[]int{1, 3, 5}).SetReadOnly(true)
	},
	"views/slice-views-3": func(parent gi.Widget) {
		giv.NewSliceViewInline(parent).SetSlice(&[]int{1, 3, 5})
	},
	"views/slice-views-4": func(parent gi.Widget) {
		giv.NewValue(parent, &[]int{1, 3, 5})
	},
	"views/slice-views-5": func(parent gi.Widget) {
		giv.NewValue(parent, &[]int{1, 3, 5, 7, 9})
	},
	"views/struct-views-0": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int
		}
		giv.NewStructView(parent).SetStruct(&person{Name: "Go", Age: 35})
	},
	"views/struct-views-1": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int
		}
		p := person{Name: "Go", Age: 35}
		giv.NewStructView(parent).SetStruct(&p).OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("You are %v", p))
		})
	},
	"views/struct-views-2": func(parent gi.Widget) {
		type person struct {
			Name string `immediate:"+"`
			Age  int
		}
		p := person{Name: "Go", Age: 35}
		giv.NewStructView(parent).SetStruct(&p).OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("You are %v", p))
		})
	},
	"views/struct-views-3": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int `view:"-"`
		}
		giv.NewStructView(parent).SetStruct(&person{Name: "Go", Age: 35})
	},
	"views/struct-views-4": func(parent gi.Widget) {
		type person struct {
			Name string `edit:"-"`
			Age  int
		}
		giv.NewStructView(parent).SetStruct(&person{Name: "Go", Age: 35})
	},
	"views/struct-views-5": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int
		}
		giv.NewStructView(parent).SetStruct(&person{Name: "Go", Age: 35}).SetReadOnly(true)
	},
	"views/struct-views-6": func(parent gi.Widget) {
		type Person struct {
			Name string
			Age  int
		}
		type employee struct {
			Person
			Role string
		}
		giv.NewStructView(parent).SetStruct(&employee{Person{Name: "Go", Age: 35}, "Programmer"})
	},
	"views/struct-views-7": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int
		}
		type employee struct {
			Role    string
			Manager person `view:"add-fields"`
		}
		giv.NewStructView(parent).SetStruct(&employee{"Programmer", person{Name: "Go", Age: 35}})
	},
	"views/struct-views-8": func(parent gi.Widget) {
		type person struct {
			Name      string `default:"Gopher"`
			Age       int    `default:"20:30"`
			Precision int    `default:"64,32"`
		}
		giv.NewStructView(parent).SetStruct(&person{Name: "Go", Age: 35, Precision: 50})
	},
	"views/struct-views-9": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int
		}
		giv.NewStructViewInline(parent).SetStruct(&person{Name: "Go", Age: 35})
	},
	"views/struct-views-10": func(parent gi.Widget) {
		type person struct {
			Name string
			Age  int
		}
		giv.NewValue(parent, &person{Name: "Go", Age: 35})
	},
	"views/struct-views-11": func(parent gi.Widget) {
		type person struct {
			Name        string
			Age         int
			Job         string
			LikesGo     bool
			LikesPython bool
		}
		giv.NewValue(parent, &person{Name: "Go", Age: 35, Job: "Programmer", LikesGo: true})
	},
	"views/table-views-0": func(parent gi.Widget) {
		type language struct {
			Name   string
			Rating int
		}
		giv.NewTableView(parent).SetSlice(&[]language{{"Go", 10}, {"Python", 5}})
	},
	"advanced/styling-0": func(parent gi.Widget) {
		parent.OnWidgetAdded(func(w gi.Widget) {
			w.Style(func(s *styles.Style) {
				s.Color = colors.C(colors.Scheme.Error.Base)
			})
		})
		gi.NewLabel(parent).SetText("Label")
		gi.NewSwitch(parent).SetText("Switch")
		gi.NewTextField(parent).SetText("Text field")
	},
	"advanced/styling-1": func(parent gi.Widget) {
		parent.OnWidgetAdded(func(w gi.Widget) {
			switch w := w.(type) {
			case *gi.Button:
				w.Style(func(s *styles.Style) {
					s.Border.Radius = styles.BorderRadiusSmall
				})
			}
		})
		gi.NewButton(parent).SetText("First")
		gi.NewButton(parent).SetText("Second")
		gi.NewButton(parent).SetText("Third")
	},
}
