/**
 * Copy the text content of the provided element to the clipboard
 * (note this must be in response to a DOM click event)
 */
window.copyElement = function($el) {
  navigator.clipboard.writeText($el.text())
}

// Initialize all "box" elements (copyable)
$('.box').each((i, el) => {
  const $c = $('<span>').text("Click to copy")
  const $v = $(el).find('.value')
  $(el)
    .css('position', 'relative')
    .append(
      $('<div>')
        .addClass('copy')
        .append($('<i>').addClass('fa-solid fa-copy'))
        .append($c)
    )
    .click(() => {
      copyElement($v)
      $c.text("Copied!")
    })
})
