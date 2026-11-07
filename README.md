# preview

Render Markdown files locally, with the look of GitHub.

preview is a small web server that renders Markdown the way GitHub
does, reloads the page when a file changes, and needs no network
access.

## Install

    go install github.com/troglobit/preview@latest

## Usage

Render a file, or serve a folder and browse it like on GitHub:

    preview README.md
    preview docs/
    🚀 Starting server: http://localhost:6419/
    📡 Auto-reload enabled. Files will trigger browser refresh.

Add `-b` to open it in a browser.  See `preview --help` for the other
options.

On a file page, the left and right arrow keys step to the previous and
next Markdown file in the same folder.  Esc goes up to the folder.
Press `?` on any page for the full list of keys.

## Origins

preview is a fork of [go-grip](https://github.com/chrishrb/go-grip) by
Christoph Herb, a Go reimplementation of the Python
[grip](https://github.com/joeyespo/grip) by Joe Esposito.
