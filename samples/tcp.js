import tcp from 'k6/tcp';

export function setup() {
	return {
		format: ['<', 'i', 'h', 'i'] // '<ihi'
	};
}

export default function(data) {
	const response = tcp.connect("10.18.98.194", 22);
	tcp.send(data.format, [4, 1, 3], [147, 8, 1, 20, 25]);
};

export function teardown(data) {
	tcp.close();
};
