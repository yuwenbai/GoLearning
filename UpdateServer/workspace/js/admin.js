$("#DiffPackage_btn").click(function(event){
    event.preventDefault();
    BtnGeneratePatch();
})

function BtnGeneratePatch()
{
	console.log("time--> BtnGeneratePatch ")
	$.get(
        "/GeneratePatch",
		function(data)
		{
			$("#currentTradeModel").text(data)
			console.log("time-->" + data)
		}
	);
}