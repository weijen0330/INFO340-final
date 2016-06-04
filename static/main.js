$(function() {
    $.get("/ping", function(data){
        if(data.error == "true"){
            $("#results").prepend("<div class='alert alert-danger'><strong>Error!</strong> "+ data.message +"</div>");
        }
    }, "json"),
    
    $.get("/query1", function(data){
        $("#firstQuery").append(data);
    }, "html")

    $("#submit").click(function() {
        $.post("/insert", 
        { 
            firstName: $("#firstName").val(), 
            middleName: $("#middleName").val(),
            lastName: $("#lastName").val(),
            description: $("#description").val()
        })
        .done(function(data){
          if(data){
            console.log(data)
            alert("Cargo request submitted!")
            $("#firstName").val(''); 
            $("#middleName").val('');
            $("#lastName").val('');
            $("#description").val('');
            } else {
            console.log("Failed grab data!")
            }          
        });
    });

    $("#search").click(function() {


    }),/*
    $.get("/query2", function(data){
        $("#secondQuery").append(data);
    }, "html")

    $.get("/query3", function(data){
        $("#thirdQuery").append(data);
    }, "html")

    $("#submit").click(function(){
      $.post("/insert", {
        firstName: $("#firstName").val(), 
        middleName: $("#middleName").val()
        lastName: $("#lastName").val(),
        description: $("#description").val()
        })
        .done(function(data){
          if(data){
            console.log(data)
          } else {
            console.log("Failed grab data!")
          }
        });
    }); */

})
