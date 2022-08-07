function isavailable(){
  let startdate = document.getElementById("sdate").value
    
  location.href = "/user/all_room_search/filter/"+startdate
}

function range (){
  let range = document.getElementById("customRange3").value
  
  if (range == 1){
    document.getElementById("rangeid").innerHTML="Price Range : Above 5k"
    location.href='/user/filter/abovefive';
  }else if (range == 0){
    document.getElementById("rangeid").innerHTML="Price Range : Below 5k"
    location.href='/user/filter/belowfive';
    
  }else if ( range == 2 ){
    document.getElementById("rangeid").innerHTML="Price Range : Above 7k"
    location.href='/user/filter/aboveseven';

  }
}

  function avrange (){
    let range = document.getElementById("customRange3").value
    
    if (range == 1){
      document.getElementById("rangeid").innerHTML="Price Range : Above 5k"
      location.href='/user/filter/avabovefive';
    }else if (range == 0){
      document.getElementById("rangeid").innerHTML="Price Range : Below 5k"
      location.href='/user/filter/avbelowfive';
      
    }else if ( range == 2 ){
      document.getElementById("rangeid").innerHTML="Price Range : Above 7k"
      location.href='/user/filter/avaboveseven';
  
    }


}

function bkrange (){
  let range = document.getElementById("customRange3").value
  
  if (range == 1){
    document.getElementById("rangeid").innerHTML="Price Range : Above 5k"
    location.href='/user/filter/bkabovefive';
  }else if (range == 0){
    document.getElementById("rangeid").innerHTML="Price Range : Below 5k"
    location.href='/user/filter/bkbelowfive';
    
  }else if ( range == 2 ){
    document.getElementById("rangeid").innerHTML="Price Range : Above 7k"
    location.href='/user/filter/bkaboveseven';

  }


}