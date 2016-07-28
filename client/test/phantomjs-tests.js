const exec = require('child_process').exec;

exports.resizeTest = function(done) {
	runPhantom('resize.ghost', done);
}

exports.copy = function(done) {
	runPhantom('copy.ghost', done);
}

exports.paste = function(done) {
	runPhantom('paste.ghost', done);
}

exports.delay = function(done) {
	runPhantom('delay.ghost',done);
}

exports.clearResults = function(done) {
	exec('rm -f  ' + __dirname + '/results/*', (error, stdout, stderr) => {
		if (error) {
			console.error(`exec error: ${error}`);
			return;
		}

		done();
	});

}


// each of these tests is capped at 2 seconds by phantomjs
function runPhantom(testFilename, done) {
	exec('phantomjs ' + __dirname + '/ghosts/' + testFilename, (error, stdout, stderr) => {
		if (error) {
			console.error(`exec error: ${error}`);
			return;
		}
	
		done();
	});
}