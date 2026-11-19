package kit

import "github.com/egoist/mygo/ui"

// palette is the shades the desktop makes of its accent, lightest first,
// as Windows 11 keeps them for its own controls.
type palette struct {
	light3, light2, light1, base, dark1, dark2, dark3 ui.Color
}

// lookupPalette returns the desktop's palette, false where there is none.
// Tests replace it, so that their pictures do not follow the desktop.
var lookupPalette = desktopPalette

// applyAccent gives t the accent Windows 11's controls take: the second
// darker shade in the light, the second lighter in the dark, with black
// text on it there. The accent the desktop reports is the base shade,
// lighter than what Windows fills a control with.
func applyAccent(t *ui.Theme) {
	p, ok := lookupPalette()
	if !ok {
		return
	}
	a, text := p.dark1, ui.Hex("#ffffff")
	if t.Dark {
		a, text = p.light2, ui.Hex("#000000")
	}
	t.Accent, t.AccentHover, t.AccentPressed, t.AccentText = a, a.Alpha(.9), a.Alpha(.8), text
}
