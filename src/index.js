import './style.scss'
import html from './feedly.html'

const feedback = () => {
  document.write(html)

}

/* function */
window.fdlyTextarea = (value) => {
  if (value && value.length > 0) {
    document.getElementsByClassName("feedback_btn")[0].removeAttribute("disabled")
  } else {
    document.getElementsByClassName("feedback_btn")[0].setAttribute("disabled", '')
  }
}

window.fdlyOpen = () => {
  document.getElementById('feedback-wrapper').setAttribute("data-v-state", "open")
}

window.fdlyClose = () => {
  let el = document.getElementById('feedback-wrapper');
  el.setAttribute("data-v-state", "close");
  setTimeout(function () {
    el.removeAttribute("data-emoticon");
    document.getElementsByClassName("feedback_textarea")[0].value = '';
    document.getElementsByClassName("feedback_btn")[0].setAttribute("disabled", '');
  }, 200);
}

window.fdlySend = () => {
  alert("send")
}

window.fdlyEmoticon = (id) => {
  document.getElementById('feedback-wrapper').setAttribute('data-emoticon', id);
}

feedback()
