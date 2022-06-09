function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
         
         
         function checkin(RID,RName,Price) {

           

            swal("Room Name:"+RName+"\nRoom Price:"+Price+"/-","Room Id:"+RID+"\n1 Bedroom 2 Bathroom 1 Hall", {           
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
            
            url: "/user/cart/"+RID,
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