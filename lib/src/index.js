import html from './mimir.html'

let els, mimirPrefixURL
const mimirRoot = document.createElement('div')

const mimir = (draw) => {
  if (!document.body.contains(mimirRoot)) {
    document.body.appendChild(mimirRoot)
  }
  if (draw) {
    // Insert html in the end of the body and save element
    mimirRoot.innerHTML = draw
    els = {
      'send': document.getElementById('mimir_send'),
      'wrapper': document.getElementById('mimir-wrapper'),
      'comment': document.getElementById('mimir_comment'),
      'email': document.getElementById('mimir_email')
    }
  }
}

window.updateMimir = mimir

/* function */
window.mimirTextarea = (value) => {
  if (value && value.length > 0) {
    els.send.removeAttribute('disabled')
  } else {
    els.send.setAttribute('disabled', '')
  }
}

window.mimirEmail = (value) => {
  if (document.getElementById('mimir_email').checkValidity()) {
    els.send.removeAttribute('disabled')
  } else {
    els.send.setAttribute('disabled', '')
  }
}

window.mimirOpen = () => {
  els.wrapper.setAttribute('data-v-state', 'open')
  els.wrapper.setAttribute('data-state', 1)
}

window.mimirClose = () => {
  els.wrapper.setAttribute('data-v-state', 'close')

  setTimeout(() => {
    els.wrapper.removeAttribute('data-emoticon')
    els.comment.value = ''
    els.email.value = ''
    els.send.setAttribute('disabled', '')
  }, 200)
}

window.mimirSend = () => {
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

  // Url setup
  let url = '/feedback'
  if (mimirPrefixURL) {
    url = mimirPrefixURL + url
  }

  // Send request
  let xhr = new XMLHttpRequest()
  xhr.open('POST', url, true)
  xhr.setRequestHeader('Content-type', 'application/json')
  xhr.send(JSON.stringify(params))

  // Close insight
  window.mimirClose()
  setTimeout(() => {
    els.wrapper.setAttribute('data-state', 3)
  }, 150)
  setTimeout(() => {
    if (els.wrapper.getAttribute('data-state') === '3') {
      els.wrapper.setAttribute('data-state', 1)
    }
  }, 3000)
}

window.mimirEmoticon = (id) => {
  els.wrapper.setAttribute('data-emoticon', id)
}

mimir(html)
