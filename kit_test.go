package kit

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

var (
	iconCircle = Icon(`<circle cx="12" cy="12" r="9"/>`)
	iconSquare = Icon(`<rect x="4" y="4" width="16" height="16" rx="2"/>`)
	iconReload = Icon(`<path d="M21 12a9 9 0 1 1-9-9"/>`)
)

// gallery shows every part of the kit on the page named page.
func gallery(c *ui.Context, page *string, back func()) {
	choice, on := 0, true
	Apply(c)
	Shell(c, func() {
		SidebarHeader(c, "Kit", "0.1.0")
		ui.Sidebar(c, page, func() {
			NavItem(c, "parts", iconCircle, "Parts", *page == "parts", nil)
			NavItem(c, "empty", iconSquare, "Empty", *page == "empty", func() { ui.Badge(c, "2") })
		}).Grow(1).Background(ui.Transparent)
	}, func() {
		switch *page {
		case "empty":
			ui.Column(c).Fill().Children(func() {
				ui.Row(c).Padding(Top(c), PageSide, 0).Children(func() { Title(c, "Empty") })
				Empty(c, iconSquare, "No items yet", "Items you add show here.", func() {
					ui.PrimaryButton(c, "Add")
					ui.Button(c, "Import")
				})
			})
		default:
			Page(c, func() {
				PageTitle(c, "Parts", &Back{Parent: "Gallery", Label: "Back", Go: back})
				Card(c, func() {
					ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
						Figure(ui.Text(c, "184.0 KB/s").FontSize(22).FontWeight(600)).Grow(1)
						IconButton(c, "reload", iconReload, "Reload", "")
					})
					Muted(c, "A detail")
				})
				Heading(c, "Heading")
				Panel(c, func() {
					SettingRow(c, "Switch", "A sentence below the label.", func() { ui.Switch(c.Key("switch"), &on) })
					SettingRow(c, "Choice", "", func() { Segmented(c, "choice", &choice, "One", "Two", "Three") })
				})
			})
		}
	})
}

func tester(view func(*ui.Context), dark bool) *ui.Tester {
	tt := ui.NewTesterWith(view, 800, 560, ui.TesterOptions{Dark: dark})
	tt.SetTitleBar(ui.TitleBar{Height: 32, Right: 138, Bottom: 32})
	tt.Frame()
	tt.Advance(time.Second)
	tt.Frame()
	return tt
}

// TestGallery draws each page of the gallery, light and dark, and compares
// it with its picture in testdata. KIT_UPDATE=1 draws the pictures again;
// KIT_SHOTS=<folder> saves the pages drawn there too.
func TestGallery(t *testing.T) {
	for _, name := range []string{"parts", "empty"} {
		for _, dark := range []bool{false, true} {
			page := name
			tt := tester(func(c *ui.Context) { gallery(c, &page, func() {}) }, dark)
			shot := name + "-light.png"
			if dark {
				shot = name + "-dark.png"
			}
			got := tt.Image()
			if dir := os.Getenv("KIT_SHOTS"); dir != "" {
				save(t, filepath.Join(dir, shot), got)
			}
			golden := filepath.Join("testdata", shot)
			if os.Getenv("KIT_UPDATE") != "" {
				save(t, golden, got)
				continue
			}
			want, err := load(golden)
			if err != nil {
				t.Fatalf("%v: run with KIT_UPDATE=1 to draw it", err)
			}
			if n := differ(want, got); n > 0 {
				out := filepath.Join(t.TempDir(), shot)
				save(t, out, got)
				t.Errorf("%s: %d pixels differ from %s; the page is in %s", shot, n, golden, out)
			}
		}
	}
}

// TestBack leads back on Escape and Alt+Left, but not while a dialog
// takes Escape.
func TestBack(t *testing.T) {
	backs, open := 0, false
	page := "parts"
	tt := tester(func(c *ui.Context) {
		gallery(c, &page, func() { backs++ })
		ui.Modal(c, &open, func() { ui.Text(c, "Dialog") })
	}, false)
	tt.Key(0, ui.KeyEscape)
	tt.Frame()
	tt.Key(ui.Alt, ui.KeyLeft)
	tt.Frame()
	if backs != 2 {
		t.Fatalf("%d of two keys led back", backs)
	}
	open = true
	tt.Frame()
	tt.Key(0, ui.KeyEscape)
	tt.Frame()
	if open || backs != 2 {
		t.Fatalf("Escape over a dialog: open %v, %d times back", open, backs-2)
	}
}

// differ counts the pixels of a and b more than a little apart.
func differ(a, b image.Image) int {
	if a.Bounds() != b.Bounds() {
		return a.Bounds().Dx() * a.Bounds().Dy()
	}
	n := 0
	r := a.Bounds()
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			r1, g1, b1, a1 := a.At(x, y).RGBA()
			r2, g2, b2, a2 := b.At(x, y).RGBA()
			for _, d := range [][2]uint32{{r1, r2}, {g1, g2}, {b1, b2}, {a1, a2}} {
				if max(d[0], d[1])-min(d[0], d[1]) > 3<<8 {
					n++
					break
				}
			}
		}
	}
	return n
}

func load(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func save(t *testing.T, path string, img image.Image) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		t.Fatal(err)
	}
}
