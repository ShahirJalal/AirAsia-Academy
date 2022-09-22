
// import http package
let http = require("http")

function iAmCallBackFunction(request, response){
    console.log(request);
    response.writeHead(200, {"Content-Type": "text/html"});
    response.write("<h1>Hello from node server </h1>");
    response.end();
}

// use http to create server
// createServer will create the server using call back function
// callback function does the actual job
var firstServer = http.createServer(iAmCallBackFunction);

firstServer.listen(8181);