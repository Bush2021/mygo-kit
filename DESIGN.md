# Design

How the apps built on this kit look and behave. The kit holds what code can
hold; this file holds the rest. A page follows both.

Behavior follows the system; the look is our own and the same on every
platform. Windows 11 is the first platform, so its conventions decide where
a rule needs one.

## Window

- A main window has no title bar: `kit.Window` hides it, keeps the
  system's window buttons over the page and puts Mica behind the window.
- `kit.Shell` lays it out: the sidebar runs to the top edge, the page is
  beside it on a lighter layer with a hairline between them. The band under
  the window buttons and the sidebar's header move the window.
- `kit.Apply` comes first in the view. Colors come from the theme it sets,
  never from a hex value in a page.
- A main window keeps its size and place (`StateKey`), and the app runs one
  instance; a second one shows the first one's window. `desk.Main` opens
  it or brings it forward, restored if it was minimized.
- Closing the main window of an app with a tray icon leaves the app in the
  tray; Quit is in the tray's menu. An app without one quits.

## Sidebar

- `kit.SidebarHeader` names the app, with one muted line below when there
  is something worth it, as a version.
- Items are `kit.NavItem`: icon and label, the chosen one tinted with the
  accent and marked with a pill. A count goes at an item's end as a badge.
- What the app is doing (running, stopped, failed) goes at the sidebar's
  foot: a dot and a muted word.

## Pages

- Every page starts with `kit.Title` at `kit.Top` from the top and
  `kit.PageSide` from the side, in the same place on every page. A page
  opened from another uses `kit.PageTitle` with a `kit.Back`: the parent's
  title before its own leads back.
- A page that scrolls is `kit.Page`, or `kit.NarrowPage` when it holds
  only settings, whose labels and controls should not drift apart on a
  wide window. A page with a table or an editor lays out its own, with its
  header row padded `kit.Top(c)`, `kit.PageSide`.
- Settings are `kit.SettingRow`s in a `kit.Panel`: label left, control
  right, a sentence below the label only when the label alone would
  mislead. A `kit.Heading` names each panel after the first.
- Anything else raised is a `kit.Card`, or `kit.Surface` on a box of its
  own layout. No box in a box: inside a card, rows are separated by space
  or hairlines.
- Settings has the theme (System, Light, Dark, a `kit.Segmented` shown
  with `desk.ShowTheme`) and the language (System and each translation),
  first under a heading of their own. Both apply at once.
- Nothing jumps. A count, a status or a button that changes keeps its
  place and width (`kit.Figure` for numbers); what comes and goes moves
  the rest with `kit.Flow`. A table keeps its columns' order and widths.

## States

- Empty: a page or list with nothing to show has `kit.Empty` where the
  content would be: an icon, a title saying what is missing, a sentence on
  what brings it when that is not plain, and the action that does, if the
  page has one. Empty because a filter matched nothing, it says so.
- Loading: the page keeps its layout and shows what it has. What is still
  coming is a muted word or a `ui.Spinner` in its place, never one over
  the whole page.
- Failed: the error is said where the thing failed, in the theme's
  `Danger`, with what to do when there is something. A failure no page
  shows is a toast. Errors are selectable, so that they can be copied.

## Color

- Layers are told apart by fill, one step at a time: window, page layer,
  surface. Raised surfaces have no outline; in the light they have a dark ring
  and a faint shadow, in the dark a ring a little lighter than the surface.
- Shadows are for what floats: menus, popups, dialogs, toasts.
- The accent marks what is chosen or on, and the one action that matters
  on a screen (`ui.PrimaryButton`). Not headings, not icons, not borders.
- Red, amber and green mean failed, in progress and fine. Nothing else uses
  them.

## Shape

- A panel's corners are `kit.PanelRadius`; controls take the theme's
  radius. A box inside another rounds by the outer radius less the space
  between them, so that the two curves run side by side.

## Text

- The theme's size for text, 12 for `kit.Muted` details, 18 for the app's
  name, `kit.TitleSize` for a page's title. Weight 600 for titles and
  headings only.
- A number that changes in place is a `kit.Figure`, so it keeps its width.
- Code, paths and identifiers are in a monospaced face.

## Controls

- Use MyGo's controls. The kit draws only those it draws differently:
  `kit.Segmented` for two to four choices that all fit, `kit.IconButton`
  for an action an icon names clearly, with its label as the tooltip.
  Dialogs and primary actions keep text.
- Icons are Lucide's shapes through `kit.Icon`.
- A switch applies at once. A choice that costs something asks first, and
  only one that cannot be undone.

## Keyboard

- Everything a click does, the keyboard does: Tab reaches it, Enter or
  Space presses it. MyGo's controls do; a box made clickable by hand needs
  `Focusable` and the keys.
- Escape closes the topmost thing: a menu, a popup, a dialog, then the
  page opened from another (`kit.PageTitle` with a `kit.Back`), which
  Alt+Left leaves too.
- A dialog's primary button takes Enter. Shortcuts follow the system's
  names (Ctrl+S here, Cmd+S on macOS) and are in the tooltip of what they
  press.

## Motion

- `kit.Flow`, `kit.RowMotion`, `kit.Tint` and `kit.PageEnter` are the
  only transitions, and their durations (`kit.TintDuration`,
  `kit.EnterDuration`, `kit.ResizeDuration`) the only ones. An element
  that comes and goes needs a Key, as do its siblings after it. An app
  checks this in a test with `kittest.Motion`.
- A size moves on a curve, never a spring: a spring's long tail moves an
  edge one pixel at a time. A marker moving between items (the sidebar's
  pill, a segmented thumb) springs, to keep its speed when the choice
  changes on the way.
- Menus, popups, dialogs and toasts move as MyGo moves them.
- Motion says where something came from or that a state changed. A page
  comes in; the one before it goes at once.
- A tester's clock stands still: advance it before taking a picture.

## Tray

- Left click shows the main window, opening it if it is closed. Right click
  opens the menu. Nothing waits for a double click. `desk.Tray` does
  this and takes the icon away when the app quits.
- An app with no main window in that mode (a background task that asks
  something) shows a flyout on a left click instead (`mygo.NewFlyout`),
  with a button that opens the app's window.
- The menu leads with what the app does (start, stop, the current mode),
  then the choices, then **Open** and **Quit**. Items have no subtitles:
  the status is in the tooltip and the icon.
- A tray icon is a template: black shapes the system tints for the taskbar.

## Words

- Say what happens; no "please", "successfully" or "are you sure".
- One verb for one action across the app.
- An ellipsis is one character: …

## Other platforms

The look stays. What changes is what the system owns: the window's
material (`c.Vibrancy()`), where the window buttons are (`c.TitleBar()`),
shortcuts and their names, the tray (a menu bar item on macOS) and the
app menu on macOS. Code for one platform goes in files named for it.
