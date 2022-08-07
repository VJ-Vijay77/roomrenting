function blockuser(ID){

    swal({
        title: "Are you sure?",
        text: "Once the user is blocked he/she cannot login.",
        icon: "warning",
        buttons: true,
        dangerMode: true,
      })
      .then((willDelete) => {
        if (willDelete) {
            if (ID==1){
                swal("Cannot Block!", "User has ADMIN Privileges!", "error");
                location.reload
                return
            }

            location.href='/admin/user_management/block/'+ID;

            swal("Successfull", "User Blocked Successfully", "success");
         
         } // } else {
        //   swal("Your imaginary file is safe!");
        // }
      });

    }
     
     
    function unblockuser(ID){

      swal({
          title: "Are you sure?",
          text: "Once the user is unblocked he/she can get all user privileges.",
          icon: "info",
          buttons: true,
          dangerMode: true,
        })
        .then((willDelete) => {
          if (willDelete) {
              if (ID==1){
                  swal("Cannot do anything", "User has ADMIN Privileges!", "error");
                  location.reload
                  return
              }
  
              location.href='/admin/user_management/unblock/'+ID;
  
              // swal("Successfull", "User UnBlocked Successfully", "success");
           
           } // } else {
          //   swal("Your imaginary file is safe!");
          // }
        });
  
      }
       
      function deleteuser(ID){

        swal({
            title: "Are you sure you want to delete?",
            text: "Once the user is deleted,all data will be erased!.",
            icon: "warning",
            buttons: true,
            dangerMode: true,
          })
          .then((willDelete) => {
            if (willDelete) {
                if (ID==1){
                    swal("Cannot Delete!", "User has ADMIN Privileges!", "error");
                    location.reload
                    return
                }
    
                location.href='/admin/user_management/delete/'+ID;
    
                swal("Successfull", "User Deleted Successfully", "success");
             
             } // } else {
            //   swal("Your imaginary file is safe!");
            // }
          });
    
        }
       