requirejs.config({
    "baseUrl": "/js/lib",
    "paths": {
      "app": "/js/app"
    },
    "shim": {
    	debounce: {
    		deps:['jquery'],
    		exports: 'debounce'
    	}
    },
});

// Load the main app module to start the app
requirejs(["app/main"]);