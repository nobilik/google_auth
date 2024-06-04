document.addEventListener('DOMContentLoaded', () => {
  checkError()
  document.getElementById('user-form').addEventListener('submit', (event) => {
    event.preventDefault()
    updateUser()
  })
})

function updateUser() {
  const form = document.getElementById('user-form')
  const email = form.elements.email.value
  const full_name = form.elements.full_name.value
  const telephone = form.elements.telephone.value
  const data = { email: email, full_name: full_name, telephone: telephone }

  fetch('/user/edit', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (response.status === 200) {
        window.location.replace(`/user/show`)
      } else {
        response.text().then((data) => {
          handledisplayError(data)
          if (response.status === 401) {
            window.location.replace(`/login`)
          }
        })
      }
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}
