define(['jquery', 'app/reporter'], function($, reporter){
	
	return {
		registerHooks: function(){
			$('form input').each(function(_, elem) {
				 	
				 	$(elem).on('copy',function(ev){
  						eventReport = {
  						"eventType": "copyAndPaste",
  						"websiteUrl": window.location.origin,
  						"sessionId": "123123-123123-123123123",
  						"pasted": false, // from what is suggested in doc we assume this means copy event.
  						"formId": $(ev.target).attr('id') };

  						reporter.postData(eventReport, "copypaste");

				 	});

				 	$(elem).on('paste',function(ev){
				 		eventReport = {
  						"eventType": "copyAndPaste",
  						"websiteUrl": window.location.origin,
  						"sessionId": "123123-123123-123123123",
  						"pasted": true, // from what is suggested in doc we assume this means copy event.
  						"formId": $(ev.target).attr('id') };

  						reporter.postData(eventReport, "copypaste");

				 	});

			});
		}
	}
});