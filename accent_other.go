//go:build !windows

package kit

const uiFont = ""

func desktopPalette() (palette, bool) { return palette{}, false }
