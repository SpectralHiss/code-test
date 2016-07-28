var endpoint;
define(['jquery'], function() {
	"use strict";
	
	return {
		postData: function(data) {
			var url = (endpoint || "https://localhost:8080/resize");
			
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