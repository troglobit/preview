// Any page: ? opens the keyboard shortcut help (#preview-help), ? again,
// Esc or a click on the backdrop closes it. While it is open, no other key
// reaches the page's own key handlers.
(function () {
  function isFormField(el) {
    if (!el || !el.tagName) return false;
    var tag = el.tagName.toLowerCase();
    return tag === "input" || tag === "textarea" || tag === "select" || el.isContentEditable;
  }

  document.addEventListener("DOMContentLoaded", function () {
    var dialog = document.getElementById("preview-help");
    if (!dialog) return;

    // Capture phase on window runs before the document listeners of the
    // other scripts. Esc still closes the dialog: that is a default action.
    window.addEventListener(
      "keydown",
      function (event) {
        if (dialog.open && event.key !== "?") event.stopImmediatePropagation();
      },
      true
    );

    document.addEventListener("keydown", function (event) {
      if (event.key !== "?" || event.defaultPrevented || event.isComposing) return;
      if (event.ctrlKey || event.altKey || event.metaKey) return;
      if (isFormField(event.target)) return;
      event.preventDefault();
      if (dialog.open) {
        dialog.close();
      } else {
        dialog.showModal();
      }
    });

    // The inner div fills the dialog, so a click on the dialog itself is on
    // the backdrop.
    dialog.addEventListener("click", function (event) {
      if (event.target === dialog) dialog.close();
    });
  });
})();
