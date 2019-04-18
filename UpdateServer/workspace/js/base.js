
// $.ajax(
//     {
// 		url: 'static/js/ConnIp.txt',
//         dataType: 'text',
//         success: function (data) {
// 			console.log(" log data is " + data)
//             localStorage.webIp = data;
//             localStorage.reqIPHeader = "http://" + localStorage.webIp + "/";
// 			console.log(" log localStorage.reqIPHeader is " + localStorage.reqIPHeader)
			
//         },
//         error: function (data) {
//             localStorage.webIp = "192.168.1.167:6001"; // frica conn
//             localStorage.reqIPHeader = "http://" + localStorage.webIp + "/";
//         }
//     })
function checkField(value){
     console.log("value is " + value)
	 //$("#featureLayerOptions")
	GetUserBindInfo(value)
}

//点击历史版本请求
$("#History_btn").click(function(event){
    event.preventDefault();
})
//点击上传跳转
$("#upload_btn").click(function(event){
    event.preventDefault();
    window.location.href = "/static/html/admin/index.html";
})
function GetPackageNames()
{
	console.log("time--> GetPakcageNames 111 ")
	$.get(
		"/GetAllPackageNames",
		function(data)
		{
			console.log("GetPackageNames-->" + data)
			           var i;
	         for (i = 0; i < data.length; i++) {
				 $("#featureLayerOptions").append("<option value='"+data[i]+"'>" +data[i] +"</option>")
			 }
		}
	);
}
GetPackageNames()
function GetUserBindInfo(version)
{
	console.log("time--> GetUserBindInfo " + version)
	$.post(
		"/GetVersion",
		{
			apkName : version
		},
		function(data)
		{
			var tempdata = data
			console.log("GetUserBindInfo-->" + data)
			$("#late_version_num").text(tempdata.VersionId)
			$("#late_version_name").text(tempdata.VersionName)
			$("#late_version_content").text(tempdata.VersionInfo)
		}
	);
}