import tcp from 'k6/tcp';

export function setup() {
	return {
		format: ['<', 'i', 'h', 'i'] // '<ihi'
	};
}

export default function(data) {
	const response = tcp.connect("10.18.98.194", 22);

	let msg = [147, 8, 1, 20, 25]
	let headers = [msg.length, 1, 3]
	let body = tcp.pack(data.format, headers, msg)
	tcp.send(body);
};

export function teardown(data) {
	tcp.close();
};
