package kit

import "github.com/egoist/mygo/ui"

// Surface styles a raised box: a fill one step above the page and no
// outline, but a ring. In the light the ring is dark, with a faint shadow,
// where a white surface on a near-white page would not show; in the dark
// it is a little lighter than the surface, as a shadow would not show.
func Surface(e ui.Element, t *ui.Theme) ui.Element {
	e = e.Radius(PanelRadius).Background(t.Surface)
	if t.Dark {
		return e.Shadow(0, 0, 0, 1, ui.RGBA(255, 255, 255, .06))
	}
	return e.Shadow(0, 0, 0, 1, ui.RGBA(0, 0, 0, .07)).Shadow(0, 1, 3, 0, ui.RGBA(0, 0, 0, .05))
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
