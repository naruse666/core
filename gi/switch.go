// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"goki.dev/colors"
	"goki.dev/cursors"
	"goki.dev/girl/abilities"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/icons"
	"goki.dev/ki/v2"
)

// Switch is a widget that can toggle between an on and off state.
// It can be displayed as a switch, checkbox, or radio button.
type Switch struct {
	WidgetBase

	// the type of switch that this is
	Type SwitchTypes `set:"-"`

	// the label text for the switch
	Text string

	// icon to use for the on, checked state of the switch
	IconOn icons.Icon `view:"show-name"`

	// icon to use for the off, unchecked state of the switch
	IconOff icons.Icon `view:"show-name"`

	// icon to use for the indeterminate (unknown) state
	IconUnk icons.Icon `view:"show-name"`
}

// SwitchTypes contains the different types of [Switch]es
type SwitchTypes int32 //enums:enum -trimprefix Switch

const (
	// SwitchSwitch indicates to display a switch as a switch (toggle slider)
	SwitchSwitch SwitchTypes = iota
	// SwitchChip indicates to display a switch as chip (like Material Design's filter chip),
	// which is typically only used in the context of [Switches].
	SwitchChip
	// SwitchCheckbox indicates to display a switch as a checkbox
	SwitchCheckbox
	// SwitchRadioButton indicates to display a switch as a radio button
	SwitchRadioButton
	// SwitchSegmentedButton indicates to display a segmented button, which is typically only used in
	// the context of [Switches].
	SwitchSegmentedButton
)

func (sw *Switch) CopyFieldsFrom(frm any) {
	fr := frm.(*Switch)
	sw.WidgetBase.CopyFieldsFrom(&fr.WidgetBase)
	sw.Type = fr.Type
	sw.Text = fr.Text
	sw.IconOn = fr.IconOn
	sw.IconOff = fr.IconOff
	sw.IconUnk = fr.IconUnk
}

func (sw *Switch) OnInit() {
	sw.WidgetBase.OnInit()
	sw.HandleEvents()
	sw.SetStyles()
}

// IsChecked tests if this switch is checked
func (sw *Switch) IsChecked() bool {
	return sw.StateIs(states.Checked)
}

// SetChecked sets the checked state and updates the icon accordingly
func (sw *Switch) SetChecked(on bool) *Switch {
	sw.SetState(on, states.Checked)
	sw.SetIconFromState()
	return sw
}

// SetIconFromState updates icon state based on checked status
func (sw *Switch) SetIconFromState() {
	if sw.Parts == nil {
		return
	}
	ist := sw.Parts.ChildByName("stack", 0)
	if ist == nil {
		return
	}
	st := ist.(*Layout)
	switch {
	case sw.StateIs(states.Indeterminate):
		st.StackTop = 2
	case sw.IsChecked():
		st.StackTop = 0
	default:
		if sw.Type == SwitchChip {
			// chips render no icon when off
			st.StackTop = -1
			return
		}
		st.StackTop = 1
	}
}

func (sw *Switch) HandleEvents() {
	sw.HandleSelectToggle() // on widgetbase
	sw.HandleClickOnEnterSpace()
	sw.OnClick(func(e events.Event) {
		e.SetHandled()
		sw.SetChecked(!sw.IsChecked())
		sw.SendChange(e)
		if sw.Type == SwitchChip {
			sw.SetNeedsLayout(true)
		}
	})
}

func (sw *Switch) SetStyles() {
	sw.Style(func(s *styles.Style) {
		s.SetAbilities(true, abilities.Activatable, abilities.Focusable, abilities.Hoverable, abilities.Checkable)
		if !sw.IsReadOnly() {
			s.Cursor = cursors.Pointer
		}
		s.Text.Align = styles.Start
		s.Text.AlignV = styles.Center
		s.Padding.Set(units.Dp(4))
		s.Border.Radius = styles.BorderRadiusSmall

		if sw.Type == SwitchChip {
			if s.Is(states.Checked) {
				s.Background = colors.C(colors.Scheme.SurfaceVariant)
				s.Color = colors.Scheme.OnSurfaceVariant
			} else {
				s.Border.Color.Set(colors.Scheme.Outline)
				s.Border.Width.Set(units.Dp(1))
				s.Padding.Left.Dp(14)
			}
		}
		if sw.Type == SwitchSegmentedButton {
			s.Border.Color.Set(colors.Scheme.Outline)
			s.Border.Width.Set(units.Dp(1))
			if s.Is(states.Checked) {
				s.Background = colors.C(colors.Scheme.SurfaceVariant)
				s.Color = colors.Scheme.OnSurfaceVariant
			}
		}

		if s.Is(states.Selected) {
			s.Background = colors.C(colors.Scheme.Select.Container)
		}
	})
	sw.OnWidgetAdded(func(w Widget) {
		switch w.PathFrom(sw) {
		case "parts":
			w.Style(func(s *styles.Style) {
				s.Gap.Zero()
				s.Align.Content = styles.Center
				s.Align.Items = styles.Center
				s.Text.AlignV = styles.Center
			})
		case "parts/stack":
			w.Style(func(s *styles.Style) {
				s.Display = styles.Stacked
				s.Grow.Set(0, 0)
				s.Gap.Zero()
			})
		case "parts/stack/icon0": // on
			w.Style(func(s *styles.Style) {
				if sw.Type == SwitchChip {
					s.Color = colors.Scheme.OnSurfaceVariant
				} else {
					s.Color = colors.Scheme.Primary.Base
				}
				// switches need to be bigger
				if sw.Type == SwitchSwitch {
					s.Min.X.Em(2)
					s.Min.Y.Em(1.5)
				} else {
					s.Min.X.Em(1.5)
					s.Min.Y.Em(1.5)
				}
			})
		case "parts/stack/icon1": // off
			w.Style(func(s *styles.Style) {
				switch sw.Type {
				case SwitchSwitch:
					// switches need to be bigger
					s.Min.X.Em(2)
					s.Min.Y.Em(1.5)
				case SwitchChip:
					// chips render no icon when off
					s.Min.X.Zero()
					s.Min.Y.Zero()
				default:
					s.Min.X.Em(1.5)
					s.Min.Y.Em(1.5)
				}
			})
		case "parts/stack/icon2": // indeterminate
			w.Style(func(s *styles.Style) {
				switch sw.Type {
				case SwitchSwitch:
					// switches need to be bigger
					s.Min.X.Em(2)
					s.Min.Y.Em(1.5)
				case SwitchChip:
					// chips render no icon when off
					s.Min.X.Zero()
					s.Min.Y.Zero()
				default:
					s.Min.X.Em(1.5)
					s.Min.Y.Em(1.5)
				}
			})
		case "parts/space":
			w.Style(func(s *styles.Style) {
				s.Min.X.Ch(0.1)
			})
		case "parts/label":
			w.Style(func(s *styles.Style) {
				s.SetNonSelectable()
				s.SetTextWrap(false)
				s.Margin.Zero()
				s.Padding.Zero()
				s.Text.AlignV = styles.Center
				s.FillMargin = false
			})
		}
	})
}

// SetType sets the styling type of the switch
func (sw *Switch) SetType(typ SwitchTypes) *Switch {
	updt := sw.UpdateStart()
	sw.Type = typ
	sw.IconUnk = icons.Blank
	switch sw.Type {
	case SwitchSwitch:
		// TODO: material has more advanced switches with a checkmark
		// if they are turned on; we could implement that at some point
		sw.IconOn = icons.ToggleOn.Fill()
		sw.IconOff = icons.ToggleOff
		sw.IconUnk = icons.ToggleMid
	case SwitchChip, SwitchSegmentedButton:
		sw.IconOn = icons.Check
		sw.IconOff = icons.None
		sw.IconUnk = icons.None
	case SwitchCheckbox:
		sw.IconOn = icons.CheckBox.Fill()
		sw.IconOff = icons.CheckBoxOutlineBlank
		sw.IconUnk = icons.IndeterminateCheckBox
	case SwitchRadioButton:
		sw.IconOn = icons.RadioButtonChecked
		sw.IconOff = icons.RadioButtonUnchecked
		sw.IconUnk = icons.RadioButtonPartial
	}
	sw.UpdateEndLayout(updt)
	return sw
}

// LabelWidget returns the label widget if present
func (sw *Switch) LabelWidget() *Label {
	lbi := sw.Parts.ChildByName("label")
	if lbi == nil {
		return nil
	}
	return lbi.(*Label)
}

// SetIcons sets the icons for the on (checked), off (unchecked)
// and indeterminate (unknown) states.  See [SetIconsUpdate] for
// a version that updates the icon rendering
func (sw *Switch) SetIcons(on, off, unk icons.Icon) *Switch {
	sw.IconOn = on
	sw.IconOff = off
	sw.IconUnk = unk
	return sw
}

// ClearIcons sets all of the switch icons to [icons.None]
func (sw *Switch) ClearIcons() *Switch {
	sw.IconOn = icons.None
	sw.IconOff = icons.None
	sw.IconUnk = icons.None
	return sw
}

func (sw *Switch) ConfigWidget() {
	config := ki.Config{}
	if sw.IconOn == "" {
		sw.IconOn = icons.ToggleOn.Fill() // fallback
	}
	if sw.IconOff == "" {
		sw.IconOff = icons.ToggleOff // fallback
	}
	ici := 0 // always there
	lbi := -1
	config.Add(LayoutType, "stack")
	if sw.Text != "" {
		config.Add(SpaceType, "space")
		lbi = len(config)
		config.Add(LabelType, "label")
	}
	sw.ConfigParts(config, func(parts *Layout) {
		ist := parts.Child(ici).(*Layout)
		ist.SetNChildren(3, IconType, "icon")
		icon := ist.Child(0).(*Icon)
		icon.SetIcon(sw.IconOn)
		icoff := ist.Child(1).(*Icon)
		icoff.SetIcon(sw.IconOff)
		icunk := ist.Child(2).(*Icon)
		icunk.SetIcon(sw.IconUnk)
		sw.SetIconFromState()
		if lbi >= 0 {
			lbl := parts.Child(lbi).(*Label)
			if lbl.Text != sw.Text {
				lbl.SetText(sw.Text)
			}
		}
	})
}

func (sw *Switch) RenderSwitch() {
	_, st := sw.RenderLock()
	sw.RenderStdBox(st)
	sw.RenderUnlock()
}

func (sw *Switch) Render() {
	sw.SetIconFromState() // make sure we're always up-to-date on render
	if sw.Parts != nil {
		ist := sw.Parts.ChildByName("stack", 0)
		if ist != nil {
			ist.(*Layout).UpdateStackedVisibility()
		}
	}
	if sw.PushBounds() {
		sw.RenderSwitch()
		sw.RenderParts()
		sw.RenderChildren()
		sw.PopBounds()
	}
}
