# preview

<picture>
  <source media="(prefers-color-scheme: dark)" srcset=".github/images/logo-dark.svg">
  <img align="right" src=".github/images/logo.svg" alt="A steaming coffee cup" width="128">
</picture>

Render Markdown files locally, with the look of GitHub.

preview is a small web server that renders Markdown the way GitHub does,
reloads the page when a file changes, and needs no network access.  It
is a fork of [go-grip](https://github.com/chrishrb/go-grip) by Christoph
Herb, which in turn is a Go reimplementation of the classic Python
[grip](https://github.com/joeyespo/grip) tool by Joe Esposito.

## Install

Each [release][] has `.deb` and `.rpm` packages and archives for Linux,
macOS and Windows, all with the `preview(1)` man page.  With Go
installed:

    go install github.com/troglobit/preview@latest

or, from a checkout, to get the `preview(1)` man page as well:

    make install

[release]: https://github.com/troglobit/preview/releases

## Usage

Render a file, or serve a folder and browse it like on GitHub:

    preview README.md
    preview docs/
    🚀 Starting server: http://localhost:6419/
    📡 Auto-reload enabled. Files will trigger browser refresh.

Add `-b` to open it in a browser.  See `man preview` or `preview --help`
for the other options.

On a file page, the left and right arrow keys step to the previous and
next Markdown file in the same folder.  Esc goes up to the folder.
Press `?` on any page for the full list of keys.
