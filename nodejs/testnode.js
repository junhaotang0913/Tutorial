var http = require('http');

http.createServer(function(req,res){
	res.writeHead(200,{'content-type':'text'});
	res.end("Hello");
	console.log("BYE");
}).listen(10080);