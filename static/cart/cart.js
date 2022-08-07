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

function proceedtocheckout(RID,Total,Startdate,Endate){
  
  const UID = document.getElementById("userids")
  const userid = UID.innerHTML
  
  location.href="/user/payment/"+Total+"/"+userid+"/"+RID+"/"+Startdate+"/"+Endate
  
}


