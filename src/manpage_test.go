package preview

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// TestManPageCoversCLI keeps doc/preview.1 in step with the code: every flag,
// Markdown feature and Markdown extension must be documented, and the port
// numbers the page quotes must match.
func TestManPageCoversCLI(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("../doc/preview.1")
	if err != nil {
		t.Fatalf("read man page: %v", err)
	}
	page := string(data)
	has := func(pattern string) bool {
		return regexp.MustCompile(`(?m)` + pattern).MatchString(page)
	}

	flags := regexp.MustCompile(`--([a-z][a-z-]*)`).FindAllStringSubmatch(rootCmd.Flags().FlagUsages(), -1)
	if len(flags) == 0 {
		t.Fatal("found no flags in the usage text")
	}
	for _, f := range flags {
		if !has(`^\.It Fl .*-` + regexp.QuoteMeta(f[1]) + `\b`) {
			t.Errorf("flag --%s is not documented", f[1])
		}
	}
	for _, feature := range SupportedMarkdownFeatures() {
		if !has(`^\.Cm ` + regexp.QuoteMeta(feature) + `\b`) {
			t.Errorf("Markdown feature %s is not documented", feature)
		}
	}
	for _, ext := range markdownExts {
		if !has(`^\.Pa ` + regexp.QuoteMeta(ext) + `( |$)`) {
			t.Errorf("Markdown extension %s is not documented", ext)
		}
	}

	if port := rootCmd.Flags().Lookup("port").DefValue; !strings.Contains(page, "port "+port) {
		t.Errorf("the default port %s is not the one documented", port)
	}
	if portFallbackTries != 20 || !strings.Contains(page, "the twenty that follow") {
		t.Errorf("the page says twenty fallback ports, the code tries %d", portFallbackTries)
	}
}
