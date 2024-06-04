function checkError() {
  const errorMessage = localStorage.getItem('displayError')
  if (errorMessage) {
    handledisplayError(errorMessage)
  }
}

function handledisplayError(message) {
  const errorElement = document.getElementById('display-error')
  errorElement.innerHTML = `<h4 style="background-color: red; color: white; min-height: 50px">${message}</h4>`
  localStorage.setItem('displayError', message)
  setTimeout(() => {
    errorElement.innerHTML = ''
    localStorage.removeItem('displayError')
  }, 3000)
}

function validateEmail(mail) {
  if (/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(mail)) {
    return true
  }
  return false
}

// clear the error message and remove from local storage on user interaction
document.getElementById('display-error').addEventListener('click', () => {
  const errorElement = document.getElementById('display-error')
  errorElement.innerHTML = ''
  localStorage.removeItem('displayError')
})
