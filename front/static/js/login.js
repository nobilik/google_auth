document.addEventListener('DOMContentLoaded', () => {
  document.getElementById('login-form').addEventListener('submit', (event) => {
    event.preventDefault()
    loginUser()
  })
  checkError()
})

function loginUser() {
  const form = document.getElementById('login-form')
  const email = form.elements.email.value
  const password = form.elements.password.value
  if (email === '' || password === '') {
    handledisplayError("email and password can't be empty")
    return
  }
  if (validateEmail(email) === false) {
    handledisplayError('email is not valid')
    return
  }
  const formData = new URLSearchParams()
  formData.append('email', email)
  formData.append('password', password)

  fetch('/login', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      if (response.status === 200) {
        window.location.replace(`/user/show`)
      } else {
        response.text().then((data) => {
          handledisplayError(data)
        })
      }
    })
    .catch((error) => {
      console.error('Error during login:', error)
    })
}
