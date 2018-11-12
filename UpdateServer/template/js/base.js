// $("#logout").click(function(event){
    // event.preventDefault();
    // del_cookie("admin_id");
    // window.location.href = "/login/index";
// })

// function del_cookie(name)
// {
    // // document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;';
// }

// $("form[data-type=formAction]").submit(function(event){
    // event.preventDefault();
    // var target = event.target;
    // var action = $(target).attr("action");
    // $.post(action, $(target).serialize(), function(ret){
        // if(ret.Ret == "0") {
            // alert(ret.Reason);
        // } else {
            // location.href = $(target).attr("form-rediret");
        // }
    // },"json")
// })
$.ajax(
    {
		url: 'static/js/ConnIp.txt',
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
	console.log("time--> GetPakcageNames ")
	$.get(
		localStorage.reqIPHeader + "GetAllPackageNames",
		function(data)
		{
			console.log("GetPackageNames-->" + data)
			           var i;
	         for (i = 0; i < data.length; i++) {
				 $("#featureLayerOptions").append("<option value='"+data[i]+"'>" +data[i] +"</option>")
			 }
			// $("#featureLayerOptions").append("<option value='"+tempdata.VersionId+"'>" +tempdata.VersionId +"</option>")
			// $("#featureLayerOptions").append("<option value='"+tempdata.VersionName+"'>");
			// $("#featureLayerOptions").append("<option value='"+100+"'>"+"text"+"</option>");
		}
	);
}
GetPackageNames()
function GetUserBindInfo(version)
{
	console.log("time--> GetUserBindInfo " + version)
	$.post(
		localStorage.reqIPHeader + "GetVersion",
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