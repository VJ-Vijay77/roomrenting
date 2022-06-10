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
