var express = require('express');
var app = express(),
    port = 8081,
    bodyParser = require('body-parser'),
    path = require('path'),
    fs = require('fs');

app.use(bodyParser.json());
app.use(express.static(__dirname));

app.get('/tests/:testName', function(req, resp) {
    var testName = req['params']['testName'];

    fs.readFile('index.html', 'utf8', function(err, contents) {
        // we are appeding tiny test code that sets urlpostvariable 
        // unfortunately it  means it has to be global..
        if (testName != "") {
            patchedHTML = contents + "<script> testEndpoint='/testData/" + testName + "';</script>"
            console.log('serving patched page for functional tests');
            resp.send(patchedHTML)
        } else {
            resp.send(contents)
        }
    });

});


app.post('/testData/:testname', function(req, resp) {

    var testName = req['params']['testname'];

    fs.writeFile(path.join(__dirname, 'test/results/', testName + '.json'), JSON.stringify(req.body), (err) => {
            if (err != null) {
                console.log(err);
            }

            console.log("test results for " + testName + " written to test/results/")
        }
    );
});

app.listen(port, function() {
    console.log("test/dev server listening on port:" + port);
});