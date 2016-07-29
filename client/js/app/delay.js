define(['jquery', 'app/reporter'], function($, reporter) {
	var timeStart = null;

	return {
		registerHooks: function() {
			$('form input').on('input',function(elem) {

					if (!Boolean(timeStart)) {
						timeStart = Date.now();
					}
				});

			$("button[type='submit']").on('submit click', function(e) {
				e.preventDefault();
				var eventData = {
					"eventType": "timeTaken",
					"websiteUrl": window.location.origin,
					"sessionId": "123123-123123-123123123",
					"time": parseInt((Date.now() - timeStart) / 1000)
				};

				reporter.postData(eventData, "delay");
			});
		}
	}
})