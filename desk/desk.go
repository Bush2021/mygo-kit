// Package desk is what the apps do with the desktop around their windows:
// the theme they show, the main window and the tray icon. DESIGN.md says
// how these behave.
package desk

import "github.com/egoist/mygo"

// The themes a setting names.
const (
	ThemeSystem = "system"
	ThemeLight  = "light"
	ThemeDark   = "dark"
)

// ShowTheme shows the theme name names, and the system's for any other.
func ShowTheme(name string) {
	switch name {
	case ThemeLight:
		mygo.Theme.SetSource(mygo.ThemeLight)
	case ThemeDark:
		mygo.Theme.SetSource(mygo.ThemeDark)
	default:
		mygo.Theme.SetSource(mygo.ThemeSystem)
	}
}

// Main is the app's main window, which it opens when it is asked for and
// there is none. Its methods belong to the main thread.
type Main struct {
	open   func() *mygo.Window
	closed func()
	win    *mygo.Window
}

// NewMain returns a Main whose window open makes. closed, when not nil,
// runs after the window has closed.
func NewMain(open func() *mygo.Window, closed func()) *Main {
	return &Main{open: open, closed: closed}
}

// Show brings the window forward, restored if it was minimized, and opens
// it if it is closed.
func (m *Main) Show() {
	if m.win != nil {
		if m.win.IsMinimized() {
			m.win.Restore()
		}
		m.win.Show()
		m.win.Focus()
		return
	}
	m.win = m.open()
	m.win.OnClosed(func() {
		m.win = nil
		if m.closed != nil {
			m.closed()
		}
	})
}

// Window is the window while it is open, and otherwise nil.
func (m *Main) Window() *mygo.Window { return m.win }

// Tray adds the app's icon to the notification area. A left click calls
// show; a right click opens o's menu, or the one given later with SetMenu.
// The app runs on when its windows close, and the icon goes when it quits.
func Tray(o mygo.TrayOptions, show func()) (*mygo.Tray, error) {
	tray, err := mygo.NewTray(o)
	if err != nil {
		return nil, err
	}
	tray.OnClick(show)
	// A listener keeps the app running when its last window closes.
	mygo.App.OnWindowAllClosed(func() {})
	// Windows keeps showing the icon of a program that has exited until the
	// pointer passes over it.
	mygo.App.OnWillQuit(func(*mygo.QuitEvent) { tray.Destroy() })
	return tray, nil
}
