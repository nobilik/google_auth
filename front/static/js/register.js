document.addEventListener('DOMContentLoaded', () => {
  document
    .getElementById('register-form')
    .addEventListener('submit', (event) => {
      event.preventDefault()
      registerUser()
    })
  checkError()
})

function registerUser() {
  const form = document.getElementById('register-form')
  const email = form.elements.email.value
  const password = form.elements.password.value
  if (email === '' || password === '') {
    handledisplayError("email and password can't be empty")
    return
  }
  console.log(validateEmail(email))
  if (validateEmail(email) === false) {
    handledisplayError('email is not valid')
    return
  }
  const formData = new URLSearchParams()
  formData.append('email', email)
  formData.append('password', password)

  fetch('/register', {
    method: 'POST',
    body: formData,
  })
    .then((response) => {
      if (response.status === 201) {
        window.location.replace(`/user/edit`)
      } else {
        response.text().then((data) => {
          handledisplayError(data)
          if (response.status === 409) {
            window.location.replace(`/login`)
          }
        })
      }
    })
    .catch((error) => {
      console.error('Error during registration:', error)
    })
}
