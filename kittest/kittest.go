// Package kittest checks in an app's tests that its pages keep to the kit.
package kittest

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// motion matches transitions and animations made outside the kit.
var motion = regexp.MustCompile(`ui\.ElementTransition\{|ui\.Animation\{|ui\.(Snappy|Smooth|Bouncy)\b|\.AnimateTo\(|\.Spring\(|\.Keyframes\(`)

// Motion fails t for each line of the Go files under dir, tests aside,
// that moves something otherwise than by the kit's transitions.
func Motion(t testing.TB, dir string) {
	t.Helper()
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for i, line := range strings.Split(string(b), "\n") {
			if motion.MatchString(line) {
				t.Errorf("%s:%d moves otherwise than by the kit's transitions: %s", path, i+1, strings.TrimSpace(line))
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
