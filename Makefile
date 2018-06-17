run-server:
	docker build -t jsonrpc-usage-example-server server
	docker run --rm --name jsonrpc-server -p 5000:5000 jsonrpc-usage-example-server

run-client:
	docker build -t jsonrpc-usage-example-client client
	docker run --rm --link jsonrpc-server jsonrpc-usage-example-client