import './style.scss'
import html from './feedly.html'

let els

const feedback = () => {
  document.write(html)
  els = {
    'send': document.getElementById('feedback_send'),
    'wrapper': document.getElementById('feedback-wrapper'),
    'comment': document.getElementById('feedback_comment'),
    'email': document.getElementById('feedback_email')
  }
}

/* function */
window.fdlyTextarea = (value) => {
  if (value && value.length > 0) {
    els.send.removeAttribute('disabled')
  } else {
    els.send.setAttribute('disabled', '')
  }
}

window.fdlyEmail = (value) => {
  if (document.getElementById('feedback_email').checkValidity()) {
    els.send.removeAttribute('disabled')
  } else {
    els.send.setAttribute('disabled', '')
  }
}

window.fdlyOpen = () => {
  els.wrapper.setAttribute('data-v-state', 'open')
  els.wrapper.setAttribute('data-state', 1)
}

window.fdlyClose = () => {
  els.wrapper.setAttribute('data-v-state', 'close')

  setTimeout(() => {
    els.wrapper.removeAttribute('data-emoticon')
    els.comment.value = ''
    els.email.value = ''
    els.send.setAttribute('disabled', '')
  }, 200)
}

window.fdlySend = () => {
  let user
  if (els.wrapper.getAttribute('data-state') === '1') {
    user = document.getElementsByTagName('head')[0].getAttribute('user-id') ||
      document.getElementsByTagName('head')[0].getAttribute('user-email')
    if (!user) {
      els.wrapper.setAttribute('data-state', 2)
      els.send.setAttribute('disabled', '')
      return
    }
  }

  // retrieve data
  let params = 'comment=' + encodeURI(els.comment.value) + '&reaction=' +
    els.wrapper.getAttribute('data-emoticon')

  if (!user) {
    user = els.email.value
  }

  if (user) {
    params += '&user=' + user
  }

  // Send request
  var xhr = new XMLHttpRequest()
  xhr.open('POST', '/insight', true)
  xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded')
  xhr.send(params)

  // Close insight
  window.fdlyClose()
  setTimeout(() => { els.wrapper.setAttribute('data-state', 3) }, 150)
  setTimeout(() => {
    if (els.wrapper.getAttribute('data-state') === '3') {
      els.wrapper.setAttribute('data-state', 1)
    }
  }, 3000)
}

window.fdlyEmoticon = (id) => {
  els.wrapper.setAttribute('data-emoticon', id)
}

feedback()
