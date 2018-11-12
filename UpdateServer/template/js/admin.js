$("#DiffPackage_btn").click(function(event){
    event.preventDefault();
    GetUserBindInfo(20);
})
$.ajax(
    {
        url: '../../js/ConnIp.txt',
        dataType: 'text',
        success: function (data) {
			console.log(" log data is " + data)
            localStorage.webIp = data;
            localStorage.reqIPHeader = "http://" + localStorage.webIp + "/";
			console.log(" log localStorage.reqIPHeader is " + localStorage.reqIPHeader)
			
        },
        error: function (data) {
            localStorage.webIp = "192.168.1.167:6001"; // frica conn
            localStorage.reqIPHeader = "http://" + localStorage.webIp + "/";
        }
    })
function GetUserBindInfo(version)
{
	console.log("time--> GetUserBindInfo ")
	$.get(
		localStorage.reqIPHeader + "GeneratePatch",
		function(data)
		{
			console.log("time-->" + data)
		}
	);
}