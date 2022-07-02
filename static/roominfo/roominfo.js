// function sleep(ms) {
//     return new Promise(resolve => setTimeout(resolve, ms));
// }



function date(Roomprice,save,Oroomprice){
 
  let startdate = document.getElementById("startdate").value
  let endate = document.getElementById("endate").value
  const t = new Date(startdate);
  const r = new Date(endate);

  let dif = Math.abs(r-t)
  const days = Math.ceil(dif / (1000 * 60 * 60 * 24)); 
  document.getElementById("price").innerHTML='Rs '+days*Roomprice
  
 
  
  document.getElementById("savingprice").innerHTML= days*save
  
  // const roomcost = document.getElementById("roomorginalprice").innerText
  
   document.getElementById("roomorginalprice").innerHTML=" "+days*Oroomprice+" "

}
         
         
         function checkin(RID,RName,Price) {
          let startdate = document.getElementById("startdate").value
  let endate = document.getElementById("endate").value
 
  const t = new Date(startdate);
  const r = new Date(endate);

  let dif = Math.abs(r-t)
  const days = Math.ceil(dif / (1000 * 60 * 60 * 24)); 

           

            swal("Room Price: Rs "+Price*days+"/-","Room Name:"+RName+"\nRoom Id:"+RID+"\n1 Bedroom 2 Bathroom 1 Hall", {           
  
              
    buttons: {
    cancel: "Cancel",
    catch: {
      text: "Add to Cart",
      value: "catch",
    },
    
  },
  
})
.then((value) => {
  switch (value) {
 
    
    case "catch":
       //location.href="/user/cart/"+RID;
        $.ajax({
          

            
            url: "/user/cart/"+RID+"/"+days+"/"+startdate+"/"+endate,
            method:"get", // serializes the form's elements.
            success: function(k) {
                if (k=="login"){
                    swal("Login First!", "You have to login first to checkout.You are redirecting to login page.", "warning");
                    setTimeout(() => { location.href="/user_login";  }, 1500);
                    
                }else if (k=="added"){
                    swal("Added Successfully", "Added to Cart!.", "success");
                    setTimeout(() => { location.href="/user/room_info/"+RID;  }, 1500);
                }else if (k=="sameroom"){
                    swal("Already in Cart", "The item is already in cart", "error");
                }


            }
         
           
            
        });
            


    //   swal("Successful", "Added to Cart Successfully", "success");
      
    //   break;
 
    // default:
    //   swal("Got away safely!");
  }
});
        

}




function wishlist(RID,RName) {

  swal("Add to Wishlist?","Room Name:"+RName, {           
    buttons: {
      cancel: "Cancel",
      
      catch: {
        text: "Add to Wishlist",
        value: "catch",
      },
      
    },
  })
  .then((value) => {
    switch (value) {
   
      
      case "catch":
         //location.href="/user/cart/"+RID;
          $.ajax({
              
              url: "/user/wishlist/"+RID,
              method:"get", // serializes the form's elements.
              success: function(k) {
                  if (k=="login"){
                      swal("Login First!", "You have to login first to checkout.You are redirecting to login page.", "warning");
                      setTimeout(() => { location.href="/user_login";  }, 1500);
                      
                  }else if (k=="added"){
                      swal("Added Successfully", "Added to Wishlist!.", "success");
                      setTimeout(() => { location.href="/user/room_info/"+RID;  }, 1500);
                  }else if (k=="sameroom"){
                      swal("Already in Wishlist", "The item is already in wishlist", "error");
                  }
  
  
              }
           
             
              
          });
                 
        }
      });
          




  // swal("Click on either the button or outside the modal.")
  // .then((value) => {

  //   swal(`The returned value is: ${value}`);
  // });

}

function booked(){

  swal("Sorry!", "The Room is already booked!");
}