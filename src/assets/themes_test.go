package assets

import (
	"io/fs"
	"regexp"
	"strings"
	"testing"
)

// TestThemeVariablesDefined checks that every variable a stylesheet reads is
// defined on the page (themes.css on :root, github-markdown.css on the body)
// or by the stylesheet itself. themes.css is generated from a hand-kept list
// in gen-themes.py, so a var(--x) added to a stylesheet without extending
// that list would otherwise resolve to nothing.
func TestThemeVariablesDefined(t *testing.T) {
	t.Parallel()

	defined := regexp.MustCompile(`--([A-Za-z0-9-]+)\s*:`)
	used := regexp.MustCompile(`var\(--([A-Za-z0-9-]+)`)

	global := map[string]bool{}
	for _, name := range []string{"themes.css", "github-markdown.css"} {
		css, err := fs.ReadFile(StaticFiles, "static/css/"+name)
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		for _, m := range defined.FindAllStringSubmatch(string(css), -1) {
			global[m[1]] = true
		}
	}

	entries, err := fs.ReadDir(StaticFiles, "static/css")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".css") {
			continue
		}
		css, err := fs.ReadFile(StaticFiles, "static/css/"+e.Name())
		if err != nil {
			t.Fatal(err)
		}
		local := map[string]bool{}
		for _, m := range defined.FindAllStringSubmatch(string(css), -1) {
			local[m[1]] = true
		}
		for _, m := range used.FindAllStringSubmatch(string(css), -1) {
			if !global[m[1]] && !local[m[1]] {
				t.Errorf("%s reads --%s, which themes.css does not define", e.Name(), m[1])
			}
		}
	}
}
