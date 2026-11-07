// Theme toggle, cycling light, dark and dimmed, and a colorblind toggle that
// swaps in GitHub's colorblind hues. The head script in layout.html owns the
// storage keys and the theme list (window.previewTheme) and applies a stored
// choice before the first paint. Without a stored choice the page follows
// the OS, so nothing is written here until a button is clicked.
(function () {
  var config = window.previewTheme;
  var root = document.documentElement;
  var ICONS = { light: "☀", dark: "☾", dimmed: "◐" };

  function currentTheme() {
    var set = root.getAttribute("data-theme");
    if (config.themes.indexOf(set) !== -1) return set;
    return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
  }

  function renderButtons() {
    var theme = currentTheme();
    var btn = document.getElementById("theme-toggle");
    if (btn) {
      btn.querySelector(".theme-toggle-icon").textContent = ICONS[theme];
      btn.title = "Theme: " + theme.charAt(0).toUpperCase() + theme.slice(1);
    }

    var on = root.hasAttribute("data-colorblind");
    var cb = document.getElementById("colorblind-toggle");
    if (cb) {
      cb.setAttribute("aria-pressed", on ? "true" : "false");
      cb.title = on ? "Colorblind palette: on" : "Colorblind palette: off";
    }
  }

  function store(key, value) {
    try {
      localStorage.setItem(key, value);
    } catch (e) {}
  }

  document.addEventListener("DOMContentLoaded", function () {
    renderButtons();
    window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", renderButtons);

    var theme = document.getElementById("theme-toggle");
    if (theme) {
      theme.addEventListener("click", function () {
        var themes = config.themes;
        var next = themes[(themes.indexOf(currentTheme()) + 1) % themes.length];
        root.setAttribute("data-theme", next);
        store(config.themeKey, next);
        renderButtons();
        document.body.dispatchEvent(new CustomEvent("themechange", { detail: { mode: next } }));
      });
    }

    var colorblind = document.getElementById("colorblind-toggle");
    if (colorblind) {
      colorblind.addEventListener("click", function () {
        var on = !root.hasAttribute("data-colorblind");
        if (on) {
          root.setAttribute("data-colorblind", "");
        } else {
          root.removeAttribute("data-colorblind");
        }
        store(config.colorblindKey, on ? "1" : "0");
        renderButtons();
      });
    }
  });
})();
