define(['jquery'], function() {
	var testData = {
		"eventType": "resize",
		"websiteUrl": "https://localhost:8182/resize",
		"sessionId": "123123-123123-123123123",
		"beforeDimensions": {
			"width": "1000px",
			"length": "1000px"
		},
		"afterDimensions": {
			"width": "100px",
			"length": "100px"
		}
	};

	var url = "https://localhost:8080/resize" || endpoint; 


	window.onload = function(){
		$.ajax({
			type: "POST",
			url: url,
			data: testData,
			success: success,
			dataType: dataType
		});
	}
});