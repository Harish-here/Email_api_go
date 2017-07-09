

var app = new Vue({
    el:'#app',
    data:{
        message : "harish",
        result : "lassnlakdlaksdalkdj"
    },
    methods:{
        postData : function(){
            var settings = {
                              "async": true,
                            //   "crossDomain": true,
                              "url": "http://localhost:3306/template",
                              "method": "POST",
                              "headers": {
                                "content-type": "application/x-www-form-urlencoded",
                                "cache-control": "no-cache",
                                "postman-token": "0d6e1286-349b-5c1f-ab66-f830e39c5da7"
                              },
                              "data": {
                                "customerName": "harish",
                                "phoneNumber": "8870072346",
                                "customerEmail": "harishamudha@gmail.com",
                                "message": "hi this sample thing i'm doing"
                              }
                            }

                      $.ajax(settings).done(function (response) {
                        this.data.result = "<h1>lool</h1>"
                      });
            
        }
    }
})