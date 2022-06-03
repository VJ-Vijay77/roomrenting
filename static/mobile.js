const mobile = document.getElementById('mobilevalidation')
const form = document.getElementById('form')
const errorElement = document.getElementById('error')

form.addEventListener('submit', (e) => {
    let messages = []

     if(mobile.value.length <= 9 || mobile.value.length>=11 ){
        messages.push('mobile must/only have 10 digits')
    }else if(!/^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$/.test(mobile.value)){
        messages.push('Enter valid Mobile number!')
    }
    
    

    if (messages.length > 0) {
        e.preventDefault()
        errorElement.innerText = messages.join(', ')
    }

})