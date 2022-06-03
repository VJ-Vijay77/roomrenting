const otp = document.getElementById('validateotp')
const form = document.getElementById('form')
const errorElement = document.getElementById('error')


form.addEventListener('submit', (e) => {
    let messages = []

    if (otp.value == "" || otp.value == null) {
        messages.push('OTP cannot be empty!')
    } else if (otp.value.length <= 3 || otp.value.length >= 5) {
        messages.push('OTP has only 4 digits!')
    } else if (! /^[0-9]+$/.test(otp.value)) {
        messages.push('OTP has only numeric digits!')
    }




    if (messages.length > 0) {
        e.preventDefault()
        errorElement.innerText = messages.join(', ')
    }

})