let startBtn = document.getElementById('start')
let stopBtn = document.getElementById('stop')

let serverInput = document.getElementById('server')
let rateInput = document.getElementById('rate')
let timer

startBtn.addEventListener('click', function () {
    timer = setInterval(sendData, Number(rateInput.value))
    startBtn.disabled = true
})

stopBtn.addEventListener('click', function () {
    clearInterval(timer)
    startBtn.disabled = false
})

function myRand() {
    return Math.floor(Math.random() * Math.floor(10)) + 1
}

function sendData() {
    console.log('sending to: ')
    console.log(serverInput.value)
    let xhr = new XMLHttpRequest()
    xhr.open("POST", serverInput.value)

    let data = []

    for (let i = 0; i <= myRand(); i++) {
        //max 10 objects
        data.push(
            {
                price: myRand(),
                quantity: myRand(),
                amount: myRand(),
                object: myRand(),
                method: myRand()
            }
        )
    }

    xhr.send(JSON.stringify(data))

    xhr.onload = function() {
        if (xhr.status !== 200) {
            console.log(`Err: ${xhr.status}: ${xhr.statusText}`)
        } else { 
            console.log('ok')
        }
      }
}