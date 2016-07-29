The fingerprint collection server:

The solution progressively fills up session id reports as instructed, we added some CORS to allow
our test-dev-server in client.

run static page server with `node ../client/test-dev-server`
and then run the collection server binary which starts the server in 8080 

browse to `localhost://8081` and trigger events, see reports in command line.

I tried to write tests to run the assertions on the command line, I got somewhere 
but couldn't finish in time, the test code was used to help with developement but tests should be run manually as explained.

