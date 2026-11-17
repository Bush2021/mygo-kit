package kit

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// How long the transitions take. Menus, popups, dialogs and toasts move as
// MyGo moves them.
const (
	TintDuration   = 120 * time.Millisecond
	EnterDuration  = 180 * time.Millisecond
	ResizeDuration = 200 * time.Millisecond
)

// resize is a curve, not a spring: a spring's long tail moves an edge by
// a pixel at a time, which shows as a judder.
var resize = ui.Animation{Duration: ResizeDuration, Ease: ui.EaseOut}

// The only transitions. An element that comes and goes needs a Key, as do
// its siblings after it.
var (
	// Flow moves what a row coming or going pushes around.
	Flow = ui.ElementTransition{Animation: &resize, Position: true, Size: true}
	// RowMotion is Flow for the row itself.
	RowMotion = ui.ElementTransition{Animation: &resize, Position: true, Size: true,
		Enter: &ui.Motion{Collapse: true}, Exit: &ui.Motion{Collapse: true}}
	// Tint fades a color that follows the pointer or a state. It is short,
	// as the pointer crosses many items on its way to one.
	Tint = ui.ElementTransition{Colors: true, Duration: TintDuration}
	// PageEnter brings a page up into place, as Windows' Settings do. The
	// page before it goes at once, so that two never show together.
	PageEnter = ui.ElementTransition{Colors: true, Duration: EnterDuration,
		Enter: &ui.Motion{Y: 12}}
	// slide moves a marker from one item to the next: the sidebar's pill,
	// a segmented control's thumb. It keeps its speed when the choice
	// changes on the way, which a curve would not.
	slide = ui.Snappy
)

// Identities for ui.ElementTransition's Match.
type (
	navPill  struct{}
	segThumb struct{ id string }
)
