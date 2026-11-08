Change Log
==========

All notable changes to the project are documented in this file.

[v1.0.0][] - 2026-11-08
-----------------------

First release of preview, a fork of [go-grip][] v0.11.0 by Christoph
Herb.  The changes below are relative to that release.

### Changes

- The project is now called preview.  The command is `preview` and the
  module `github.com/troglobit/preview`, installed with `go install
  github.com/troglobit/preview@latest`
- Release packages: `.deb` and `.rpm` packages and archives for Linux,
  macOS and Windows, for amd64 and arm64
- New `preview(1)` man page, installed by the packages and by `make
  install`
- The browser is no longer opened on start, use `-b` to open it
- When port 6419 is busy and `-p` is not given, the first free port of
  the twenty that follow is used instead of failing
- Renamed options: `--bounding-box=false` is now `--no-border`, and
  `--disable-markdown-feature` is now `-d`, `--disable`
- GitHub's dark dimmed theme joins light and dark, and a new button
  switches to GitHub's colorblind palette.  All colours, code
  highlighting included, follow GitHub's current palettes
- A stored theme is applied before the page is drawn, so a dark page no
  longer flashes light while loading.  Theme and width choices stored by
  go-grip are not carried over
- On a file page, the left and right arrow keys, or `h` and `l`, step to
  the previous and next Markdown file in the folder, and Esc goes up to
  the folder
- On a folder page, Esc, `h`, `l` and the right arrow key work too
- Press `?` on any page for the list of its keys
- Live reload can be paused for one page with a new button.  By Andrew
  Chen
- New option `--mermaid-icon-pack` lets Mermaid diagrams use iconify
  icons.  The browser fetches the packs from unpkg.com, so this option
  needs network access.  By Toshihito Kon
- Files ending in `.markdown`, `.mdown`, `.mdwn`, `.mkd`, `.mkdn` and
  `.mkdown` are rendered like `.md`, also as a folder's README
- Stylesheets and scripts are revalidated instead of downloaded again on
  every page view

### Fixes

- Raw HTML blocks that end on a closing tag, such as a `<pre>` block,
  were dropped from the page

[go-grip]: https://github.com/chrishrb/go-grip


[v0.11.0][] - 2026-10-29
------------------------

Last release of go-grip, see its [release notes][v0.11.0-notes].

[v0.11.0-notes]: https://github.com/chrishrb/go-grip/releases/tag/v0.11.0

[UNRELEASED]: https://github.com/troglobit/preview/compare/v0.11.0...HEAD
[v1.0.0]:     https://github.com/troglobit/preview/compare/v0.11.0...v1.0.0
[v0.11.0]:    https://github.com/chrishrb/go-grip/releases/tag/v0.11.0
