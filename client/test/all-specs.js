var seleniumTests = require('./selenium-tests.js'),
	fs = require('fs'),
	path = require('path');

var should = require('should');


describe('end-to-end fingerprinter tests',function(){

before(function() {
	seleniumTests.clearResults();
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
	beforeEach(function() {
		seleniumTests.resizeTest();
	});

	it('reports the correct json event report', function() {
		matchesFixtures("resize.json");
	});
});

})