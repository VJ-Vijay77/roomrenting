const firstname = document.getElementById('form3Example1')
const lastname = document.getElementById('form3Example2')
const email = document.getElementById('form3Example3')
const password = document.getElementById('form3Example4')
const mobile = document.getElementById('form3Example5')
const form = document.getElementById('form')
const errorElement = document.getElementById('error')
var regName = /^[a-zA-Z]+ [a-zA-Z]+$/;

form.addEventListener('submit', (e) => {
    let messages = []
    //first name
    if (firstname.value==='' || firstname.value== null){
        messages.push('Name is required!')
    }else if(firstname.value.length <= 2 ){
        messages.push('Name must have 3 characters')
    }else if (!/^[A-Za-z ]+$/.test(firstname.value)){
        messages.push('No special characters allowed!')
    }
    
    //last name
  else  if (lastname.value==='' || lastname.value== null){
        messages.push('Last Name is required!')
    }

    else if (!/^[A-Za-z ]+$/.test(lastname.value)){
        messages.push('No special characters allowed!')
   
        //email
    }else if (!/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email.value)){
        messages.push('Enter valid email address!')
    }else  if (email.value==='' || email.value== null){
        messages.push('Email is required!')
        
        //password
    }else if(password.value ==='' || password.value==null){
        messages.push('Password should not be empty!')
    }else if(password.value.length <= 4 ){
        messages.push('password must have 5 characters')
   
   //mobile validation
    }else if(mobile.value.length <= 9 || mobile.value.length>=11 ){
        messages.push('mobile must/only have 10 digits')
    }else if(!/^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$/.test(mobile.value)){
        messages.push('Enter valid Mobile number!')
    }
    
    

    if (messages.length > 0) {
        e.preventDefault()
        errorElement.innerText = messages.join(', ')
    }

})