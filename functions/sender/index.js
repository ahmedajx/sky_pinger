console.log('starting function')

var twilio = require('twilio');
var http = require('http');

var client = new twilio.RestClient(
    process.env.ACCOUNTSID,
    process.env.AUTHTOKEN
);

exports.handle = function() {
    http.get("http://skysportsapi.herokuapp.com/sky/football/getteamnews/chelsea/v1.0/", function (response) {
        response.setEncoding('utf8');
        var news = '';
        response.on('data', function(data){
            news += data;
        });
        response.on('end',function(){
            client.messages.create({
                body: JSON.parse(news)[0].shortdesc,
                to: process.env.FROMTEL,
                from: process.env.TOTEL
            }, function(err, message) {
                console.log(message.sid);
            });
        });
    })
}