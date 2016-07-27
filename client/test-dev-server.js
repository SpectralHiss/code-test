var express = require('express');
var app = express(),
	port=8081,
	bodyParser = require('body-parser'),
    fs = require('fs');

app.use(bodyParser.json());
app.use(express.static('.'));

app.get('/tests/:testName',function(req,resp) {
	var testName = req['params']['testName'];

    fs.readFile('index.html', 'utf8', function(err, contents) {
    	// we are appeding tiny test code that sets urlpostvariable 
    	// unfortunately it  means it has to be global..
    	if (testName != "") {
    		patchedHTML = contents + "<script> endpoint='/testData/"+testName+"';</script>" 		
    		console.log('serving patched page for functional tests');
    		resp.send(patchedHTML)
    	} else {
    		resp.send(contents)	
    	}
    });

});


app.post('/testData/:testnumber',function (req, resp){
	console.log("got test json output"+resp.body);
	fs.writeFile('./spec/results/'+ req['params']['testnumber']+'.json');
});

app.listen(port,function(){
	console.log("test/dev server listening on port:"+port);
});




