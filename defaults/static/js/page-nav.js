// Markdown page keys: Esc follows the link back to the containing folder
// (#preview-back), ArrowLeft and ArrowRight follow the page's rel=prev and
// rel=next links to the neighbouring Markdown files. A key is left alone
// when it already means something else here.
(function () {
  function isFormField(el) {
    if (!el || !el.tagName) return false;
    var tag = el.tagName.toLowerCase();
    return tag === "input" || tag === "textarea" || tag === "select" || el.isContentEditable;
  }

  // MathJax closes its context menu and its dialogs on Esc. Leave the keys to
  // it while focus is inside math, a menu or a dialog, or while one is open.
  var mathJaxSelector =
    'mjx-container, mjx-dialog, .mjx-dialog, [class^="CtxtMenu_"], [class*=" CtxtMenu_"]';

  function inMathJax(el) {
    return !!(el && el.closest && el.closest(mathJaxSelector));
  }

  function mathJaxOverlayOpen() {
    return !!document.querySelector(
      '.CtxtMenu_MenuFrame, .CtxtMenu_Menu, mjx-dialog, .mjx-dialog, dialog[open]'
    );
  }

  function linkFor(key) {
    switch (key) {
      case "Escape":
        return document.getElementById("preview-back");
      case "ArrowLeft":
        return document.querySelector('link[rel="prev"]');
      case "ArrowRight":
        return document.querySelector('link[rel="next"]');
      default:
        return null;
    }
  }

  function onKeyDown(event) {
    if (event.defaultPrevented || event.isComposing) return;
    if (event.ctrlKey || event.altKey || event.metaKey || event.shiftKey) return;
    var link = linkFor(event.key);
    if (!link) return;
    if (isFormField(event.target) || inMathJax(event.target)) return;
    if (mathJaxOverlayOpen()) return;

    event.preventDefault();
    window.location.href = link.href;
  }

  document.addEventListener("keydown", onKeyDown);
})();
