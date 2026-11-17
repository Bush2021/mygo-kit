package kit

import "github.com/egoist/mygo/ui"

// Muted is secondary text: details, captions, values beside a label.
func Muted(c *ui.Context, s string) ui.Element {
	return ui.Text(c, s).FontSize(12).TextColor(c.Theme().TextMuted)
}

// Heading names the panel below it.
func Heading(c *ui.Context, s string) ui.Element {
	return ui.Text(c, s).FontWeight(600).Margin(12, 0, 0).Transition(Flow)
}

// Figure shows a number that changes in place with digits of one width, so
// that it holds its width.
func Figure(e ui.Element) ui.Element { return e.FontFeatures("tnum") }

// Back leads from a page to the one it opened from.
type Back struct {
	// Parent is the title of the page it leads to.
	Parent string
	// Label names the link to assistive technology, as "Back".
	Label string
	Go    func()
}

// PageTitle heads a page: in the same place on every page, as tall as a
// row of buttons beside it, after the page it opened from when back is
// not nil, as Windows' Settings show where a page is. Escape and Alt+Left
// then lead back too; a dialog or menu that is open takes Escape first.
func PageTitle(c *ui.Context, s string, back *Back) ui.Element {
	t := c.Theme()
	if back != nil && (c.Shortcut(0, ui.KeyEscape) || c.Shortcut(ui.Alt, ui.KeyLeft)) {
		back.Go()
	}
	return ui.Row(c.Key("page-title")).MinHeight(40).MinWidth(0).Gap(8).AlignItems(ui.Center).Children(func() {
		if back != nil {
			link := ui.Text(c.Key("parent"), back.Parent).FontSize(TitleSize).FontWeight(600).TextColor(t.TextMuted).SingleLine().
				Cursor(ui.CursorPointer).Label(back.Label).Transition(Tint)
			if link.Hovered() {
				link.TextColor(t.Text)
			}
			if link.Clicked() {
				back.Go()
			}
			ui.Icon(c, iconNext).TextColor(t.TextMuted)
		}
		ui.Text(c, s).FontSize(TitleSize).FontWeight(600).SingleLine().MinWidth(0)
	})
}

// Title is the PageTitle of a page that opened from none.
func Title(c *ui.Context, s string) ui.Element { return PageTitle(c, s, nil) }
