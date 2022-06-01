const email = document.getElementById('form2Example17')
const password = document.getElementById('form2Example27')
const form = document.getElementById('form')
const errorElement = document.getElementById('error')



form.addEventListener('submit', (e) => {
    let messages = []
    //email
    if (email.value==='' || email.value==null){
        messages.push('Email cannot be empty!')
    }else if (/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email.value)){
        messages.push('Enter valid email address!')
   
        //password
    }else if(password.value ==='' || password.value==null){
        messages.push('Password should not be empty!')
    }else if(password.value.length <= 4 ){
        messages.push('password must have 5 characters')
    }

if (messages.length > 0) {
    e.preventDefault()
    errorElement.innerText = messages.join(', ')
}

})