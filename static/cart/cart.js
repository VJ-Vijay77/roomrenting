function removecartitem(CID) {

    swal({
        title: "Are you sure you want to remove item from cart",
        text: "!!",
        icon: "warning",
        buttons: true,
        dangerMode: true,
      })
      .then((willDelete) => {
        if (willDelete) {

            location.href='/user/cart/delete/'+CID;

            swal("Successfull", "Item Removed Successfully", "success");
         
         } // } else {
        //   swal("Your imaginary file is safe!");
        // }
      });


}

function proceedtocheckout(totalprice,userID){
  location.href="/user/payment/"+totalprice+"/"+userID
  
}

