// Package kit is how the MyGo apps of one author look and move: the
// theme, the window's frame, the surfaces pages are made of and the
// controls drawn differently from MyGo's. DESIGN.md says when to use which.
package kit

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// Sizes, in DIPs.
const (
	SidebarWidth = 220
	// PageSide is the room beside a page's content.
	PageSide = 32
	// PanelRadius rounds cards and panels; controls inside them take the
	// theme's smaller one.
	PanelRadius = 12
	// TitleSize is the size of a page's title.
	TitleSize = 20
	// RowMinHeight is a setting row's height with a one-line label.
	RowMinHeight = 56
	// ContentWidth is the widest a NarrowPage's content grows, so that a
	// row's label and its control stay within one glance.
	ContentWidth = 1120
)

// Window gives a main window its frame: no title bar, the system's window
// buttons over the page, and Mica behind it.
func Window(o mygo.WindowOptions) mygo.WindowOptions {
	o.TitleBarStyle = mygo.TitleBarHidden
	o.Vibrancy = mygo.VibrancyMica
	return o
}

// Apply sets the window's theme: Windows 11's surfaces, translucent so that
// Mica shows through, and text colors that reach 4.5:1 on them. Call it
// first in the view.
func Apply(c *ui.Context) {
	t := *c.Theme()
	white := ui.Hex("#ffffff")
	if t.Dark {
		t.Background = ui.Hex("#202020")
		t.Surface = white.Alpha(.06)
		t.SurfaceHover = white.Alpha(.085)
		t.SurfacePressed = white.Alpha(.12)
		t.Border = white.Alpha(.08)
		t.TextMuted = ui.Hex("#a6a6a6")
	} else {
		black := ui.Hex("#000000")
		t.Background = ui.Hex("#f3f3f3")
		t.Surface = white.Alpha(.7)
		t.SurfaceHover = black.Alpha(.04)
		t.SurfacePressed = black.Alpha(.07)
		t.Border = black.Alpha(.09)
		t.TextMuted = ui.Hex("#616169")
		t.Danger = ui.Hex("#c42b1c")
		t.Warning = ui.Hex("#9d5d00")
		t.Success = ui.Hex("#0f7b0f")
	}
	applyAccent(&t)
	t.Font = uiFont
	c.SetTheme(&t)
}
