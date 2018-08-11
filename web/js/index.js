(function() {
  window.onload = function() {
    let fetchButton = document.querySelector('#fetch')
    let clearButton = document.querySelector('#clear')
    let resultsList = document.querySelector('#body ul')

    function API() {
      this.get_request = function(endpoint) {
        let xhr = new XMLHttpRequest()
        xhr.open('GET', `/api/${endpoint}`)
        xhr.responseType = 'json'
        return xhr
      }

      this.gibberish = function() {
        var req = this.get_request('gibberish')
        req.onload = function(e) {
          console.log(this.response)
          let li = document.createElement('li')
          li.textContent = this.response.text
          resultsList.appendChild(li)
        }
        req.send()
      }
      return this
    }
    var api = API()

    fetchButton.onclick = function() {
      api.gibberish();
    }

    clearButton.onclick = function() {
      let lis = resultsList.getElementsByTagName('li')
      for (var i = lis.length - 1; i >= 0; --i) {
        lis[i].remove()
      }
    }
  }
})()
