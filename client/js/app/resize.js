define(['jquery', 'debounce', 'app/reporter'], function($, debounce, reporter) {
	"use strict";
	
	var oldSize = getSize();

	function getSize() {
		return {
			height: $(document).height(),
			width: $(document).width()
		}
	}

	return {
		registerHooks: function() {
			$(window).resize($.throttle(function() {
				var newSize = getSize();	// this is to prevent some strange behaviour where resize is triggered in phantomjs.
				if ( !(newSize['height'] == oldSize['height'] && newSize['width'] == oldSize['width']) ) { 
					var resizeEvent = {
						"eventType": "resize",
						"websiteUrl": "https://localhost:8182/resize",
						"sessionId": "123123-123123-123123123", // i suppose this is meant
						// to be a session cookie? setting to constant
						"resizeFrom": oldSize,
						"resizeTo": newSize
					};

					reporter.postData(resizeEvent);
					oldSize = newSize;
				}
			}, 200));

		}

	}
});