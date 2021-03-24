.PHONY: proto-user
proto-user:
	protoc -I protos/ protos/user.proto --go_out=protos/user --go-grpc_out=protos/user

.PHONY: proto-email
proto-email:
	protoc -I protos/ protos/email.proto --go_out=protos/email --go-grpc_out=protos/email

.PHONY: proto-order
proto-order:
	protoc -I protos/ protos/order.proto --go_out=protos/order --go-grpc_out=protos/order

.PHONY: proto-fare
proto-fare:
	protoc -I protos/ protos/fare.proto --go_out=protos/fare --go-grpc_out=protos/fare

proto-push:
	protoc -I protos/ protos/push.proto --go_out=protos/push --go-grpc_out=protos/push

.PHONY: proto-node
proto-node:
	cp /protos/notify.proto /node/protos

.PHONY: grpc-web-order
grpc-web-order:
	protoc -I protos/ protos/order.proto --js_out=import_style=commonjs:webapp/protos/order --grpc-web_out=import_style=commonjs,mode=grpcwebtext:webapp/protos/order

.PHONY: grpc-web-fare
grpc-web-fare:
	protoc -I protos/ protos/fare.proto --js_out=import_style=commonjs:webapp/protos/fare --grpc-web_out=import_style=commonjs,mode=grpcwebtext:webapp/protos/fare

.PHONY: grpc-web-user
grpc-web-fare:
	protoc -I protos/ protos/user.proto --js_out=import_style=commonjs:webapp/protos/user --grpc-web_out=import_style=commonjs,mode=grpcwebtext:webapp/protos/user

.PHONY: run-prettier-admin-panel
run-prettier-admin-panel:
	cd adminapp && npx prettier --write .

.PHONY: run-prettier-webapp
run-prettier-webapp:
	cd webapp && npx prettier --write .

