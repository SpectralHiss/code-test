var webdriver = require('selenium-webdriver');

const exec = require('child_process').exec;


exports.resizeTest = function() {
	var driver = new webdriver.Builder().forBrowser('firefox').build();
	driver.get("http://localhost:8081/tests/resize");
	driver.manage().window().setSize('1000', '1000');
	driver.wait(function() {
		return true;
	}, 100000);
	driver.quit();
}

exports.clearResults = function() {
	exec('rm -f  '+ __dirname + '/results/*', (error, stdout, stderr) => {
		if (error) {
			console.error(`exec error: ${error}`);
			return;
		}
		console.log(`stderr: ${stderr}`);
	});

}