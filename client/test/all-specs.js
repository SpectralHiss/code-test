var phantomjsTests = require('./phantomjs-tests.js'),
	fs = require('fs'),
	path = require('path'),
	should = require('should');

describe('end-to-end fingerprinter tests', function() {
	this.timeout(30000); // 30 second for whole suite maximum.

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

			// each of these tests is capped at 2 seconds by phantomjs
			phantomjsTests.resizeTest(done);

		});

		it('reports the exact json event report in fixtures/ with correct dimensions delta etc..', function(done) {
			matchesFixtures("resize.json")
			done();
		});
	});

	describe('copying form elements', function() {
		beforeEach(function(done) {
			phantomjsTests.copy(done);
		});

		it('reports the exact json event report in fixtures/ with correct type and target etc..', function(done) {
			matchesFixtures("copy.json")
			done();
		});
	})


	describe('pasting form elements ', function() {
		beforeEach(function(done) {
			phantomjsTests.paste(done);
		});

		it('reports the exact json event report in fixtures/ with correct type and target, etc...', function(done) {
			matchesFixtures("paste.json")
			done();
		});
	})

	// this test could last up to 15 seconds
	describe('user delay to submit', function() {
		beforeEach(function(done) {
			phantomjsTests.delay(done);
		});

		it('reports the correct json with a delay aproximatively in expected range', function() {
			// because we cannot guarantee exact times in test and javascript..

			var result = fs.readFileSync(path.join(__dirname, "results/delay.json"));
			var testOutput = JSON.parse(result)

			testOutput["eventType"].should.be.equal("timeTaken");
			testOutput["websiteUrl"].should.be.equal("http://localhost:8081")
			testOutput["sessionId"].should.be.equal("123123-123123-123123123")
			testOutput["time"].should.be.approximately(7,1) // seconds
		})

	})
});