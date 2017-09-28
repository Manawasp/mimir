import './robin.scss'
import html from './robin.html'

let els

const robin = (draw) => {
  const main = document.getElementById('#robin-wrapper')
  if (main) {
    document.body.removeChild(main)
  }
  if (draw) {
    // Insert html in the end of the body and save element
    document.write(draw)
    els = {
      'send': document.getElementById('robin_send'),
      'wrapper': document.getElementById('robin-wrapper'),
      'comment': document.getElementById('robin_comment'),
      'email': document.getElementById('robin_email')
    }
  }
}

window.updateRobin = robin

/* function */
window.robinTextarea = (value) => {
  if (value && value.length > 0) {
    els.send.removeAttribute('disabled')
  } else {
    els.send.setAttribute('disabled', '')
  }
}

window.robinEmail = (value) => {
  if (document.getElementById('robin_email').checkValidity()) {
    els.send.removeAttribute('disabled')
  } else {
    els.send.setAttribute('disabled', '')
  }
}

window.robinOpen = () => {
  els.wrapper.setAttribute('data-v-state', 'open')
  els.wrapper.setAttribute('data-state', 1)
}

window.robinClose = () => {
  els.wrapper.setAttribute('data-v-state', 'close')

  setTimeout(() => {
    els.wrapper.removeAttribute('data-emoticon')
    els.comment.value = ''
    els.email.value = ''
    els.send.setAttribute('disabled', '')
  }, 200)
}

window.robinSend = () => {
  let userId
  let email
  if (els.wrapper.getAttribute('data-state') === '1') {
    userId = document.getElementsByTagName('head')[0].getAttribute('user-id')
    email = document.getElementsByTagName('head')[0].getAttribute('user-email')
    if (!userId && !email) {
      els.wrapper.setAttribute('data-state', 2)
      els.send.setAttribute('disabled', '')
      return
    }
  }

  // retrieve data
  let params = {
    'message': els.comment.value,
    'emoji': parseInt(els.wrapper.getAttribute('data-emoticon'))
  }

  if (!email) {
    email = els.email.value
  }

  if (userId) {
    params.user_id = userId
  }
  if (email) {
    params.email = email
  }

  // Send request
  var xhr = new XMLHttpRequest()
  xhr.open('POST', '/api/feedback', true)
  xhr.setRequestHeader('Content-type', 'application/json')
  xhr.send(JSON.stringify(params))

  // Close insight
  window.robinClose()
  setTimeout(() => {
    els.wrapper.setAttribute('data-state', 3)
  }, 150)
  setTimeout(() => {
    if (els.wrapper.getAttribute('data-state') === '3') {
      els.wrapper.setAttribute('data-state', 1)
    }
  }, 3000)
}

window.robinEmoticon = (id) => {
  els.wrapper.setAttribute('data-emoticon', id)
}

robin(html)
