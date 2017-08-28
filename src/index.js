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
  if (els.wrapper.getAttribute('data-state') === '1') {
    els.wrapper.setAttribute('data-state', 2)
    els.send.setAttribute('disabled', '')
  } else {
    window.fdlyClose()
    setTimeout(() => { els.wrapper.setAttribute('data-state', 3) }, 150)
    setTimeout(() => {
      if (els.wrapper.getAttribute('data-state') === '3') {
        els.wrapper.setAttribute('data-state', 1)
      }
    }, 3000)
  }
}

window.fdlyEmoticon = (id) => {
  els.wrapper.setAttribute('data-emoticon', id)
}

feedback()
