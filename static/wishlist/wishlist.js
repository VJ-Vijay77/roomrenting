function removefromwishlist(WID) {

    swal({
        title: "Are you sure you want to remove item from wishlist?",
        text: "!!",
        icon: "warning",
        buttons: true,
        dangerMode: true,
      })
      .then((willDelete) => {
        if (willDelete) {

            location.href='/user/wishlist/delete/'+WID;

            swal("Successfull", "Item Removed Successfully", "success");
         
         } // } else {
        //   swal("Your imaginary file is safe!");
        // }
      });


}



  
function wishlistTocart(RID,RoomName) {

           

    swal("Do you want to know more about this product","Room Name:"+RoomName, {           
buttons: {
cancel: "Cancel",
catch: {
text: "Check-In",
value: "catch",
},

},
})
.then((value) => {
switch (value) {


case "catch":
//location.href="/user/cart/"+RID;
$.ajax({
    
    url: "/user/wishlist/cart/"+RID,
    method:"get", // serializes the form's elements.
    success: function(k) {
        if (k=="login"){
            swal("Login First!", "You have to login first to checkout.You are redirecting to login page.", "warning");
            setTimeout(() => { location.href="/user_login";  }, 1500);
            
        }else if (k=="fine"){
            // swal("Added Successfully", "Added to Cart!.", "success");
            location.href="/user/room_info/"+RID;
        }// }else if (k=="sameroom"){
        //     swal("Already in Cart", "The item is already in cart", "error");
        // }


    }
 
   
    
});
    

}
});


}
