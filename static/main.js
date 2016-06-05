$(function() {
    $.get("/ping", function(data){
        if(data.error == "true"){
            $("#results").prepend("<div class='alert alert-danger'><strong>Error!</strong> "+ data.message +"</div>");
        }
    }, "json"),

    $.get("/query1", function(data){
        $("#firstQuery").append(data);
        }, "html")

    $("#reset").click(function() { 
        $("#firstQuery").empty();
        $.get("/query1", function(data){
            $("#firstQuery").append(data);
        }, "html")
    })

    $("#submit").click(function() {
        $.post("/insert", 
        { 
            firstName: $("#firstName").val(), 
            middleName: $("#middleName").val(),
            lastName: $("#lastName").val(),
            description: $("#description").val()
        })
        .done(function(data) {
            if(data){
                alert("Cargo request submitted!")
                $("#firstName").val(''); 
                $("#middleName").val('');
                $("#lastName").val('');
                $("#description").val('');
            } else {
                console.log("Failed to grab data!")
            }          
        });
    })

    $("#search").click(function() {
        $.post("/update",
        {   
            searchBox: $("#searchbox").val()
        }, "html")
        .done(function(data) {
            if (data) {
                console.log(data)
                $("#firstQuery").empty()
                $("#firstQuery").append(data)
            } else {
                console.log("Failed  to grab data!")
            }
        });
    });
})
