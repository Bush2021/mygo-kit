package kit

import "github.com/egoist/mygo/ui"

// Icon makes an icon of the shapes of a 24-unit Lucide icon.
func Icon(shapes string) *ui.SVG {
	return ui.MustParseSVG([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">` + shapes + `</svg>`))
}

var iconNext = Icon(`<path d="m9 18 6-6-6-6"/>`)

// IconButton is a button that shows only glyph, with no face until the
// pointer is over it. label names it to assistive technology and to tests,
// and shows as its tooltip unless tip says more.
func IconButton(c *ui.Context, key string, glyph *ui.SVG, label, tip string) ui.Element {
	if tip == "" {
		tip = label
	}
	b := ui.Button(c.Key(key), "").Label(label).Tooltip(tip).Padding(7).Background(ui.Transparent).Border(1, ui.Transparent)
	if b.Hovered() {
		b.Background(c.Theme().SurfaceHover)
	}
	return b.Children(func() { ui.Icon(c, glyph) })
}

// Segmented is ui.Segmented with a thumb that slides to the segment chosen.
// id tells it from the other controls in the window.
func Segmented(c *ui.Context, id string, selected *int, labels ...string) ui.Element {
	t := c.Theme()
	radius := max(t.Radius-2, 0)
	thumb := ui.Hex("#ffffff")
	if t.Dark {
		thumb = t.SurfacePressed
	}
	parts := ui.SegmentedBase(c.Key(id), selected, len(labels))
	return parts.Track.AlignItems(ui.Stretch).Padding(2).Gap(2).Radius(t.Radius).Background(t.SurfaceHover).Border(1, t.Border).Children(func() {
		for i, label := range labels {
			s := parts.Segment(i).Padding(4, 12).Radius(radius)
			if i != *selected && s.Hovered() {
				s.Background(t.SurfaceHover)
			}
			s.Children(func() {
				text := ui.Text(c, label).SingleLine()
				if i != *selected {
					return
				}
				// Absolute elements paint above the flow, so only another one
				// can show above the thumb. The label in the flow keeps the
				// segment's size.
				text.Opacity(0)
				ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Radius(radius).Background(thumb).
					Shadow(0, 1, 2, 0, ui.RGBA(0, 0, 0, .12)).
					Transition(ui.ElementTransition{Animation: &slide, Match: segThumb{id}})
				ui.Text(c, label).SingleLine().Attach(ui.AnchorCenter, ui.AnchorCenter)
			})
		}
	})
}
