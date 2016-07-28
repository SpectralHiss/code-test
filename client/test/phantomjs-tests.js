const exec = require('child_process').exec;

var debug = process.env['DEBUG_TESTS'];

exports.resizeTest = function(done) {
	runPhantom('resize.ghost', done);
}

exports.copy = function(done) {
	runPhantom('copy.ghost', done);
}

exports.paste = function(done) {
	runPhantom('paste.ghost', done)
}

exports.clearResults = function(done) {
	exec('rm -f  ' + __dirname + '/results/*', (error, stdout, stderr) => {
		if (error) {
			console.error(`exec error: ${error}`);
			return;
		}

		if (debug == true) {
			console.log(`stdout: ${stdout}`);
		}

		done();
	});

}

function runPhantom(testFilename, done) {
	exec('phantomjs ' + __dirname + '/ghosts/' + testFilename, (error, stdout, stderr) => {
		if (error) {
			console.error(`exec error: ${error}`);
			return;
		}
		if (debug == true) {
			console.log(`stdout: ${stdout}`);
		}
		done();
	});
}