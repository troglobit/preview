package preview

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "preview [file | folder]",
	Short:   "Preview Markdown files in the browser, rendered like GitHub",
	Version: buildVersion(),
	Args:    cobra.MatchAll(cobra.OnlyValidArgs),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, err := logLevel(cmd)
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// avoid printing the usage text on a bad host or a busy port
		cmd.SilenceUsage = true

		browser, _ := cmd.Flags().GetBool("browser")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		noBorder, _ := cmd.Flags().GetBool("no-border")
		disabledFeatures, _ := cmd.Flags().GetStringSlice("disable")
		iconPacks, _ := cmd.Flags().GetStringSlice("mermaid-icon-pack")
		noReload, _ := cmd.Flags().GetBool("no-reload")
		logger, err := newLogger(cmd, os.Stderr)
		if err != nil {
			return err
		}

		var file string
		if len(args) == 1 {
			file = args[0]
		}

		if err := ValidateMarkdownFeatures(disabledFeatures); err != nil {
			return err
		}

		parser := NewParser(disabledFeatures)
		server := NewServer(host, port, !noBorder, browser, !noReload, parser, logger)
		server.UsePortFallback(!cmd.Flags().Changed("port"))
		if err := server.SetMermaidIconPacks(iconPacks); err != nil {
			return fmt.Errorf("--mermaid-icon-pack: %w", err)
		}
		return server.Serve(file)
	},
}

// Execute runs the command line and returns its error, for main to exit on.
func Execute() error {
	return rootCmd.Execute()
}

// logLevel returns the validated --log-level value.
func logLevel(cmd *cobra.Command) (slog.Level, error) {
	value, _ := cmd.Flags().GetString("log-level")
	level, err := ParseLogLevel(value)
	if err != nil {
		return 0, fmt.Errorf("--log-level: %w", err)
	}
	return level, nil
}

// newLogger builds the one logger the server and its watcher share.
func newLogger(cmd *cobra.Command, w io.Writer) (*slog.Logger, error) {
	level, err := logLevel(cmd)
	if err != nil {
		return nil, err
	}
	return NewLogger(w, level), nil
}

func init() {
	rootCmd.SetVersionTemplate("{{.Name}} {{.Version}}\n")
	rootCmd.Flags().BoolP("browser", "b", false, "Open a browser tab")
	rootCmd.Flags().StringP("host", "H", "localhost", "Host or IP to listen on")
	rootCmd.Flags().IntP("port", "p", 6419, "Port to use, the next free one is tried when the default is busy")
	rootCmd.Flags().Bool("no-border", false, "Render without the border around the page content")
	rootCmd.Flags().StringSliceP(
		"disable", "d",
		nil,
		"Disable Markdown feature(s): "+strings.Join(SupportedMarkdownFeatures(), ", "),
	)
	rootCmd.Flags().StringSlice(
		"mermaid-icon-pack",
		nil,
		"Fetch iconify icon pack(s) for Mermaid diagrams from unpkg.com, e.g. logos (needs network)",
	)
	rootCmd.Flags().Bool("no-reload", false, "Disable automatic browser reload on file changes")
	rootCmd.Flags().String("log-level", "info", "Log level on stderr: error, warn, info or debug")
}

// These are set at build time via -ldflags, e.g.
//
//	go build -ldflags "-X github.com/troglobit/preview/src.version=v1.2.3"
//
// When they are left empty the values are derived from the embedded build
// info, which covers `go install` and builds from a checkout.
var (
	version = ""
	commit  = ""
	date    = ""
)

// buildVersion assembles the string reported by `preview --version`.
func buildVersion() string {
	v, c, d := version, commit, date
	modified := false

	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				if c == "" {
					c = s.Value
				}
			case "vcs.time":
				if d == "" {
					d = s.Value
				}
			case "vcs.modified":
				modified = s.Value == "true"
			}
		}
	}

	// A version stamped by the go tool is either a release tag or a
	// pseudo-version that already contains the commit and date, so it is
	// reported as-is.
	if v == "" && ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}

	if v == "" {
		v = "dev"
	}

	var details []string
	if c != "" {
		if len(c) > 12 {
			c = c[:12]
		}
		if modified {
			c += "-dirty"
		}
		details = append(details, "commit "+c)
	}
	if d != "" {
		details = append(details, "built "+d)
	}

	if len(details) == 0 {
		return v
	}
	return v + " (" + strings.Join(details, ", ") + ")"
}
