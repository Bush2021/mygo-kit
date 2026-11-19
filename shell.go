package kit

import "github.com/egoist/mygo/ui"

// Shell lays out a main window: the sidebar and beside it the page, both
// on the window's material, below the band of the system's window buttons,
// which moves the window. The app's name and the page's title share a line.
func Shell(c *ui.Context, sidebar, page func()) {
	bar := c.TitleBar()
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		// The header's padding and the title's taller row put both on the
		// same line.
		top := Top(c) + 2
		if bar.Left > 0 {
			// macOS' window buttons are over the sidebar.
			top = bar.Bottom + 8
		}
		ui.Column(c).Width(SidebarWidth).Shrink(0).Padding(top, 8, 16).Gap(8).Children(sidebar)
		ui.Column(c).Grow(1).MinWidth(0).ClipY().Children(func() {
			page()
			ui.Box(c).Absolute().Top(0).Left(0).Right(bar.Right).Height(bar.Height).DragWindow()
		})
	})
	// Last, as SetTheme sets the root's color again.
	if c.Vibrancy() {
		c.Root().Background(ui.Transparent)
	}
}

// SidebarHeader names the app at the top of the sidebar, and moves the
// window.
func SidebarHeader(c *ui.Context, name, detail string) ui.Element {
	t := c.Theme()
	return ui.Column(c).Padding(4, 12, 8).Gap(2).DragWindow().Children(func() {
		ui.Text(c, name).FontSize(18).FontWeight(600)
		if detail != "" {
			ui.Text(c, detail).Font("monospace").FontSize(11).TextColor(t.TextMuted).SingleLine()
		}
	})
}

// NavItem is an item of a ui.Sidebar. The chosen one is tinted with the
// accent and marked with a pill; MyGo would fill it with the accent while
// the list has the focus. extra adds to the item's end, as a badge.
func NavItem(c *ui.Context, id string, icon *ui.SVG, label string, chosen bool, extra func()) ui.Element {
	t := c.Theme()
	item := *t
	if chosen {
		item.Accent, item.InactiveAccent, item.AccentText = t.Accent, t.Accent, t.Accent
	} else {
		item.Accent = t.TextMuted
	}
	c.SetTheme(&item)
	row := ui.SidebarItem(c, id, icon, label)
	c.SetTheme(t)
	if chosen {
		row.Background(t.Accent.Alpha(.12)).TextColor(t.Accent)
	}
	return row.Children(func() {
		if chosen {
			ui.Box(c).Absolute().Left(0).Top(8).Bottom(8).Width(3).Radius(2).Background(t.Accent).
				Transition(ui.ElementTransition{Animation: &slide, Match: navPill{}})
		}
		if extra != nil {
			extra()
		}
	})
}

// Top is the room above a page's title: the band of the window buttons.
func Top(c *ui.Context) float32 { return max(c.TitleBar().Height, 8) + 4 }

// Page is a page that scrolls, its content one column as wide as the
// window. Pages that hold a table or an editor lay out their own, with a
// header row padded Top(c), PageSide.
func Page(c *ui.Context, body func()) ui.Element {
	return ui.Scroll(c).Grow(1).Padding(Top(c), PageSide, PageSide).Children(func() {
		ui.Column(c).Gap(12).FillWidth().Children(body)
	})
}

// NarrowPage is a Page whose content grows no wider than ContentWidth, in
// the middle of a wider window, for pages of settings, whose rows read
// badly when far apart.
func NarrowPage(c *ui.Context, body func()) ui.Element {
	return ui.Scroll(c).Grow(1).Padding(Top(c), PageSide, PageSide).AlignItems(ui.Center).Children(func() {
		ui.Column(c).Gap(12).FillWidth().MaxWidth(ContentWidth).Children(body)
	})
}
