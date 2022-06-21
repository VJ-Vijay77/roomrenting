function isavailable(){
  let startdate = document.getElementById("sdate").value
    
  $.ajax({
    URL : '/user/all_room_search/'+startdate,
    METHOD : 'GET',
    success : function(){
        
    }



  })
}