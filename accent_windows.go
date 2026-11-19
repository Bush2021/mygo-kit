package kit

import (
	"sync"
	"syscall"
	"time"

	"github.com/egoist/mygo/ui"
)

// uiFont is Windows 11's own, which WinUI's text is in; Windows 10 lacks
// it and DirectWrite falls back to Segoe UI.
const uiFont = "Segoe UI Variable Text"

var (
	paletteMu   sync.Mutex
	paletteAt   time.Time
	paletteLast palette
	paletteOK   bool
)

// desktopPalette reads Explorer's AccentPalette, at most once a second, as
// the view asks each frame and the user may change the accent meanwhile.
func desktopPalette() (palette, bool) {
	paletteMu.Lock()
	defer paletteMu.Unlock()
	if time.Since(paletteAt) < time.Second {
		return paletteLast, paletteOK
	}
	paletteAt = time.Now()
	paletteLast, paletteOK = readPalette()
	return paletteLast, paletteOK
}

func readPalette() (palette, bool) {
	var key syscall.Handle
	path, _ := syscall.UTF16PtrFromString(`Software\Microsoft\Windows\CurrentVersion\Explorer\Accent`)
	if syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, path, 0, syscall.KEY_READ, &key) != nil {
		return palette{}, false
	}
	defer syscall.RegCloseKey(key)
	name, _ := syscall.UTF16PtrFromString("AccentPalette")
	var b [32]byte
	n := uint32(len(b))
	var kind uint32
	if syscall.RegQueryValueEx(key, name, nil, &kind, &b[0], &n) != nil || kind != syscall.REG_BINARY || n < 28 {
		return palette{}, false
	}
	c := func(i int) ui.Color { return ui.RGBA(b[i*4], b[i*4+1], b[i*4+2], 1) }
	return palette{c(0), c(1), c(2), c(3), c(4), c(5), c(6)}, true
}
