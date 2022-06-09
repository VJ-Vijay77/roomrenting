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

            swal("Successfull", "User Deleted Successfully", "success");
         
         } // } else {
        //   swal("Your imaginary file is safe!");
        // }
      });


}