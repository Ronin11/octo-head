const express = require('express');
const http = require('http');

const config = require('./config');
const clog = require('./util/clog');

const app = express();
const server = http.createServer(app);
const io = require('socket.io')(server);

io.on('connection', function(socket){
	socket.on('chat message', function(msg){
		console.log('message: ' + msg);
	});
});

server.listen(config.port, function listening() {
	clog.i('Listening on port ', server.address().port);
});
