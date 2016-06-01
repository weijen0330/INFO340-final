$(function(){
    $.get("/ping", function(data){
        if(data.error == "true"){
            $("#results").prepend("<div class='alert alert-danger'><strong>Error!</strong> "+ data.message +"</div>");
        }
    }, "json")
    /*
    $.get("/query1", function(data){
        $("#firstQuery").append(data);
    }, "html")

    $.get("/query2", function(data){
        $("#secondQuery").append(data);
    }, "html")

    $.get("/query3", function(data){
        $("#thirdQuery").append(data);
    }, "html")
*/
    $("#submit").click(function(){
      $.post("/insert", {
        username: $("#username").val(), 
        password: $("#password").val(),
        
        })
        .done(function(data){
          if(data){
            console.log(data)
            $("#result").text("Logged in as: " + data.username);
          } else {
            console.log("Failed to log in!")
            $("#result").text("Username / pasword combination invalid!");
          }
        });
    });

})
