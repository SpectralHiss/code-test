var phantomjsTests = require('./phantomjs-tests.js'),
	fs = require('fs'),
	path = require('path');

var should = require('should');


describe('end-to-end fingerprinter tests',function(){
this.timeout(20000);

before(function(done) {
	phantomjsTests.clearResults(done);
});

var matchesFixtures = function(testName) {
	
	var result = fs.readFileSync(path.join(__dirname, "results", testName));
	var testOutput = JSON.parse(result)


	var fixtureBuffer = fs.readFileSync(path.join(__dirname, "fixtures", testName));


	var expected = JSON.parse(fixtureBuffer)
	should.exist(testOutput);
	should.deepEqual(testOutput, expected);
}

describe('resizing the window', function() {
	
	beforeEach(function(done) {
		phantomjsTests.resizeTest(done);
	});


	it('reports the correct json event report', function(done) {
		matchesFixtures("resize.json")
		done();
	});
});

describe('copying form elements',function(){
	beforeEach(function(done) {
		phantomjsTests.copy(done);
	});

	it('reports the correct json event report', function(done) {
		matchesFixtures("copy.json")
		done();
	});
})


describe('pasting form elements ',function(){
	beforeEach(function(done) {
		phantomjsTests.paste(done);
	});

	it('reports the correct json event report', function(done) {
		matchesFixtures("paste.json")
		done();
	});
})




});
