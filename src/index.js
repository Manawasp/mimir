import './style.scss'
import html from './feedly.html'

const feedback = () => {
  document.write(html)
}

window.fdlyOpen = () => {
  document.getElementById('feedback-wrapper').setAttribute("data-v-state", "open");
}

window.fdlyClose = () => {
  document.getElementById('feedback-wrapper').setAttribute("data-v-state", "close");
}

window.fdlySend = () => {
  alert("send")
}

feedback()
