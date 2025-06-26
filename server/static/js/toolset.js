/**
 * Copy the text content of the provided element to the clipboard
 * (note this must be in response to a DOM click event)
 */
window.copyElement = function($el) {
  navigator.clipboard.writeText($el.text())
}

// Initialize all "box" elements (copyable)
$('.box').each((i, el) => {
  const $v = $('<span>').text("Click to copy")
  $(el)
    .css('position', 'relative')
    .append(
      $('<div>')
        .addClass('copy')
        .append($('<i>').addClass('fa-solid fa-copy'))
        .append($v)
    )
    .click(() => {
      copyElement($v)
      $v.text("Copied!")
    })
})
