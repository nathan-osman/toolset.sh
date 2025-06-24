/**
 * Copy the text content of the provided element to the clipboard
 * (note this must be in response to a DOM click event)
 */
window.copyElement = function($el) {
  navigator.clipboard.writeText($el.text())
}

/**
 * Show a popup attached to the provided element
 */
window.showPopup = function($el, text) {
  $el.find('.popup').remove()
  $el.append(
    $('<div>').addClass('popup').text(text),
  )
}

// Initialize all 'single' components
$('.single').each((i, el) => {
  const $el = $(el)
  const $v = $el.find('.value')
  $el.click(() => {
    copyElement($v)
    showPopup($el, "Copied!")
  })
})
