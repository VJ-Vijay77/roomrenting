
function Frstname() {

    const firstname = document.getElementById('firstname')
    // const lastname = document.getElementById('lastname')
    // const email = document.getElementById('email')
    // const mobile = document.getElementById('mobile')
    const form = document.getElementById('form1')
    const errorElement = document.getElementById('error')
    
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
        if (messages.length > 0) {
            e.preventDefault()
            errorElement.innerText = messages.join(', ')
        }
    })
    //ends here
    }
    
    
    function Lstname() {
    
        // const firstname = document.getElementById('firstname')
        const lastname = document.getElementById('lastname')
        // const email = document.getElementById('email')
        // const mobile = document.getElementById('mobile')
        const form = document.getElementById('form2')
        const errorElement = document.getElementById('error2')
        
        form.addEventListener('submit', (e) => {
            let messages = []
            
            if (lastname.value==='' || lastname.value== null){
                messages.push('Last Name is required!')
            }else if(lastname.value.length < 1 ){
                messages.push('Last Name must have 1 characters')
            }else if (!/^[A-Za-z ]+$/.test(lastname.value)){
                messages.push('No special characters allowed!')
            }
            if (messages.length > 0) {
                e.preventDefault()
                errorElement.innerText = messages.join(', ')
            }
        })
        //ends here
        }
    
        function Mobile() {
    
            // const firstname = document.getElementById('firstname')
            // const lastname = document.getElementById('lastname')
            // const email = document.getElementById('email')
            const mobile = document.getElementById('mobile')
            const form = document.getElementById('form3')
            const errorElement = document.getElementById('error3')
            
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
            //ends here
            }
        
            function Email() {
    
                // const firstname = document.getElementById('firstname')
                // const lastname = document.getElementById('lastname')
                const email = document.getElementById('email')
                // const mobile = document.getElementById('mobile')
                const form = document.getElementById('form4')
                const errorElement = document.getElementById('error4')
                
                form.addEventListener('submit', (e) => {
                    let messages = []
                    
                    if (!/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email.value)){
                        messages.push('Enter valid email address!')
                    }else  if (email.value==='' || email.value== null){
                        messages.push('Email is required!')
                        
                        //password
                    }
                   
        
                    if (messages.length > 0) {
                        e.preventDefault()
                        errorElement.innerText = messages.join(', ')
                    }
                })
                //ends here
                }
        
        
                function Editaddress(ID){
    const housename = document.getElementById('housename')
    const place = document.getElementById('place')
    const mobile2 = document.getElementById('mobile2')
    const PIN = document.getElementById('pin')
    const state = document.getElementById('state')
    const form = document.getElementById('form5')
    const errorElement = document.getElementById('error5')
    var regName = /^[a-zA-Z]+ [a-zA-Z]+$/;
    
    form.addEventListener('submit', (e) => {
        let messages = []
        //first name
        if (housename.value==='' || housename.value== null){
            messages.push('House Name is required!')
        }else if(housename.value.length <= 2 ){
            messages.push('Houes Name must have atleast 3 characters')
        }else if (!/^[A-Za-z ]+$/.test(housename.value)){
            messages.push('No special characters allowed!')
        }
       else if (place.value==='' || place.value== null){
            messages.push('place is required!')
        }else if(place.value.length <= 2 ){
            messages.push('place must have 3 characters')
        }else if (!/^[A-Za-z ]+$/.test(place.value)){
            messages.push('No special characters allowed!')
        
       //mobile validation
        }else if(mobile2.value.length <= 9 || mobile2.value.length>=11 ){
            messages.push('mobile must/only have 10 digits')
        }else if(!/^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$/.test(mobile2.value)){
            messages.push('Enter valid Mobile number!')
        }
        else if (state.value==='' || state.value== null){
            messages.push('state is required!')
        }else if(state.value.length <= 2 ){
            messages.push('state must have 3 characters')
        }else if (!/^[A-Za-z ]+$/.test(state.value)){
            messages.push('No special characters allowed!')
        
        }else if(PIN.value.length <= 5 || PIN.value.length>=7 ){
            messages.push('PIN only have 6 digits')
        }else if(!/^\(?([0-9]{2})\)?[-. ]?([0-9]{2})[-. ]?([0-9]{2})$/.test(PIN.value)){
            messages.push('Enter valid PIN number!')
        }
        
    
        if (messages.length > 0) {
            e.preventDefault()
            errorElement.innerText = messages.join(', ')
        }else if (messages.length<=0){
            e.preventDefault()
            Checkaddress(ID)
        }
    
    })
                }



        
    function Checkaddress(ID){

    swal({
        title: "Are you sure the details entered is correct?",
        text: "Press OK to confirm.",
        icon: "warning",
        buttons: true,
        dangerMode: true,
      })
      .then((willDelete) => {
        if (willDelete) {
      
            $.post('/user/edit_profile/editaddress/'+ID, $('#form5').serialize())
      
    setTimeout(() => {swal("Updated Successfully", "We respect your privacy in sensitive informations.", "success"); },1000);
      setTimeout(() => { location.href="/user/edit_profile";  }, 3000);
          
      
      
          
        }
      });
      }
      

    

      //edit address two

      function Editaddresstwo(ID){
        const housename = document.getElementById('housename')
        const place = document.getElementById('place')
        const mobile2 = document.getElementById('mobile2')
        const PIN = document.getElementById('pin')
        const state = document.getElementById('state')
        const form = document.getElementById('form6')
        const errorElement = document.getElementById('error6')
        var regName = /^[a-zA-Z]+ [a-zA-Z]+$/;
        
        form.addEventListener('submit', (e) => {
            let messages = []
            //first name
            if (housename.value==='' || housename.value== null){
                messages.push('House Name is required!')
            }else if(housename.value.length <= 2 ){
                messages.push('Houes Name must have atleast 3 characters')
            }else if (!/^[A-Za-z ]+$/.test(housename.value)){
                messages.push('No special characters allowed!')
            }
           else if (place.value==='' || place.value== null){
                messages.push('place is required!')
            }else if(place.value.length <= 2 ){
                messages.push('place must have 3 characters')
            }else if (!/^[A-Za-z ]+$/.test(place.value)){
                messages.push('No special characters allowed!')
            
           //mobile validation
            }else if(mobile2.value.length <= 9 || mobile2.value.length>=11 ){
                messages.push('mobile must/only have 10 digits')
            }else if(!/^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$/.test(mobile2.value)){
                messages.push('Enter valid Mobile number!')
            }
            else if (state.value==='' || state.value== null){
                messages.push('state is required!')
            }else if(state.value.length <= 2 ){
                messages.push('state must have 3 characters')
            }else if (!/^[A-Za-z ]+$/.test(state.value)){
                messages.push('No special characters allowed!')
            
            }else if(PIN.value.length <= 5 || PIN.value.length>=7 ){
                messages.push('PIN only have 6 digits')
            }else if(!/^\(?([0-9]{2})\)?[-. ]?([0-9]{2})[-. ]?([0-9]{2})$/.test(PIN.value)){
                messages.push('Enter valid PIN number!')
            }
            
        
            if (messages.length > 0) {
                e.preventDefault()
                errorElement.innerText = messages.join(', ')
            }else if (messages.length<=0){
                e.preventDefault()
                Checkaddresstwo(ID)
            }
        
        })
                    }
    
    
    
            
        function Checkaddresstwo(ID){
    
        swal({
            title: "Are you sure the details entered is correct?",
            text: "Press OK to confirm.",
            icon: "warning",
            buttons: true,
            dangerMode: true,
          })
          .then((willDelete) => {
            if (willDelete) {
          
                $.post('/user/edit_profile/editaddresstwo/'+ID, $('#form6').serialize())
          
        // setTimeout(() => {swal("Updated Successfully", "We respect your privacy in sensitive informations.", "success"); },1000);
        //   setTimeout(() => { location.href="/user/edit_profile_two";  }, 3000);
              
          
          
              
            }
          });
          }
          
    