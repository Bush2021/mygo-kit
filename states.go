package kit

import "github.com/egoist/mygo/ui"

// Empty stands where a page or a list has nothing to show: an icon, a
// title saying what is missing, a sentence on what brings it when detail
// is not empty, and the actions that do, the first a ui.PrimaryButton.
func Empty(c *ui.Context, icon *ui.SVG, title, detail string, actions func()) ui.Element {
	t := c.Theme()
	return ui.Column(c).Grow(1).FillWidth().Padding(48, PageSide).Gap(8).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
		ui.Icon(c, icon).Size(32, 32).TextColor(t.TextMuted).Margin(0, 0, 4)
		ui.Text(c, title).FontWeight(600).TextAlign(ui.Center)
		if detail != "" {
			Muted(c, detail).TextAlign(ui.Center).MaxWidth(360)
		}
		if actions != nil {
			ui.Row(c).Gap(8).Margin(8, 0, 0).Children(actions)
		}
	})
}
