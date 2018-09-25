import tcp from 'k6/tcp';
import proto from 'k6/tcp/proto';

export let options = {
	rps: 1,
	vus: 1,
	duration: '3s'
};

export function setup() {
	tcp.connect("10.18.98.194", 22);
	return {
		format: ['<', 'i', 'h', 'i'] // '<ihi'
	};
}

export default function(data) {
	let loginRequest = JSON.stringfy({
		"account":"test1",
		"platform":"1",
		"device":"1",
		"serverID":1454420002,
		"gameID":"1",
		"opID":"1",
		"opGameID":"1",
		"clientVer":"2018.07.31",
		"passwd":"1",
		"timeShift":0,
		"debug":1,
		"sdkVer":"4.0",
		"channelId":"",
		"mobID":"",
		"os":"",
		"deviceType":""
	});
	let msg = proto.newMessage(loginRequest, 'proto.ClientLoginRequest');
	let headers = [msg.length, 1, 3];
	let request = tcp.pack(data.format, headers, msg);
	tcp.send(request);
}

export function teardown(data) {
	tcp.close();
}
