The fingerprinting library



1. To run tests:

ensure node is installed
run `npm install`
ensure mocha and phantomjs are installed globally
run `npm test`


test-dev-server.js serves static page at port 8081 for normal user interactions
the browser code posts results to `/result/test`

It makes use of special endpoints to facilitate acceptance testing, see below.


1. Testing approach:

A special endpoint exists under `/test/:testname`.
This serves a patched index page that makes the fingerprint collector code post to another special endpoint: `/testData/:testname` 
That endpoint saves the JSON it receives to `test/results/testname.json`.

when mocha runs it wipes previous results in ,`test/results/*` then calls the appropriate phantom 
script in `test/ghosts/*.ghost` . This functionality is in `test/phantomjs-tests.js` **it works by shelling out to global phantomjs**.

these **phantomjs scripts** load previously mentioned patched page and simulate each user interactions.

**After phantom the mocha suite compares the results on disk with the fixtures in fixtures/testname.json**

Note. functional tests can be a bit messy, the solution forces me to add a global variable 
where we communicate the value of the "endpoint" it is not fatal to approach because
in production environment this can be fixed with very little post-processing.

1. Fingerprint collector

The collector is composed of 2 main parts:

1. a reporter that is in charge of posting data
1. specialsed scripts under name of each kind of event it is registering listeners for

there were a few challenges with resize triggering too much, we made use of jquery throttle to remedy
this which simply tapers invocation to a period of 200ms.

The solution is pretty straightfoward, the bulk of the work was writting the testing 'rig'