var testEndpoint;
define(['jquery'], function() {
	"use strict";

	var serverResultEndpoint =  "https://localhost:8080/result/"
	
	return {
		postData: function(data, resource) {
			var url = (testEndpoint || (serverResultEndpoint+resource));
			
			$.ajax({
				url: url,
				type: "POST",
				data: JSON.stringify(data),
				contentType: "application/json; charset=utf-8",
				dataType: "json"
			}).done(function(response) {
				console.log("OK");
			}).fail(function(jqXHR, textStatus, errorThrown) {
				console.error(errorThrown);
			});
		}
	}
});