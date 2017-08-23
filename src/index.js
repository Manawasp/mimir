import './style.scss'
import html from './feedly.html'

const feedback = () => {
  document.write(html)
}

window.fdlyOpen = () => {
  alert("open")
}

window.fdlyClose = () => {
  alert("close")
}

window.fdlySend = () => {
  alert("send")
}

feedback()
