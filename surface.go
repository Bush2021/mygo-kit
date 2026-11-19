package kit

import "github.com/egoist/mygo/ui"

// Surface styles a raised box: a fill one step above the window and a
// faint ring, darker in the light and lighter in the dark, no shadow.
func Surface(e ui.Element, t *ui.Theme) ui.Element {
	ring := ui.RGBA(0, 0, 0, .06)
	if t.Dark {
		ring = ui.RGBA(255, 255, 255, .07)
	}
	return e.Radius(PanelRadius).Background(t.Surface).Shadow(0, 0, 0, 1, ring)
}

// Card is a surface that holds anything.
func Card(c *ui.Context, body func()) ui.Element {
	return Surface(ui.Column(c).Padding(16).Gap(8), c.Theme()).Transition(Flow).Children(body)
}

// Panel holds rows as one surface, with a hairline between them.
func Panel(c *ui.Context, rows func()) ui.Element {
	t := c.Theme()
	return Surface(ui.Column(c), t).Dividers(1, t.Border).Transition(Flow).Children(rows)
}

// SettingRow is a label, with a sentence below it when detail is not empty,
// and a control at its end. One that comes and goes needs a Key, as do
// those below it in its panel.
func SettingRow(c *ui.Context, label, detail string, control func()) ui.Element {
	return ui.Row(c).Padding(12, 16).Gap(16).AlignItems(ui.Center).MinHeight(RowMinHeight).Transition(RowMotion).Children(func() {
		ui.Column(c).Grow(1).MinWidth(0).Gap(2).Children(func() {
			ui.Text(c, label)
			if detail != "" {
				Muted(c, detail)
			}
		})
		control()
	})
}
