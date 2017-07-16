var concatArgs = function(args){
	var msg = "";
	for(arg in args){
		if(typeof args[arg] === 'object')
			msg += JSON.stringify(args[arg]);
		else
			msg += args[arg];
	}
	return msg;
}

var createMessage = function(msg, tag){
	var date = new Date();
	return "[" + date.getTime() + ' - ' + date.toLocaleString() + ' - ' + tag +"]: " + msg + "\n";
}

exports.i = function(){
	var tag = "INFO";
	var message = createMessage(concatArgs(arguments), tag);
	console.log(message);
}

exports.d = function(){
	var tag = "DEBUG";
	var message = createMessage(concatArgs(arguments), tag);
	console.log(message);
}

exports.e = function(){
	var tag = "ERROR";
	var message = createMessage(new Error(concatArgs(arguments)).stack, tag);
	console.log(message);
}