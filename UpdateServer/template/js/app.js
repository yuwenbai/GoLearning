$.ajax(
    {
		url: 'static/js/ConnIp.txt',
        dataType: 'text',
        success: function (data) {
			console.log(" log data is " + data)
            webIp = data;
            reqIPHeader = "http://" + localStorage.webIp + "/";
			console.log(" log localStorage.reqIPHeader is " + localStorage.reqIPHeader)
			
        },
        error: function (data) {
            localStorage.webIp = "192.168.1.167:6001"; // frica conn
            localStorage.reqIPHeader = "http://" + localStorage.webIp + "/";
        }
    })