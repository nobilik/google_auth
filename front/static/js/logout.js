function logout() {
  fetch('/logout', {
    method: 'DELETE',
  })
    .then((response) => {
      if (response.status === 200) {
        document.cookie =
          'test_session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
        window.location.replace(`/login`)
      } else {
        response.text().then((data) => {
          handledisplayError(data)
        })
      }
    })
    .catch((error) => {
      console.error('Error during logout:', error)
    })
}
