proto-gen:
	@echo "Options: foo | bar"
	@read -p "Enter service name:" service; \
	protoc -Iproto \
	--go_out=services \
	--go_opt=paths=import \
	--go-grpc_out=services \
	--go-grpc_opt=paths=import \
	--validate_out="lang=go:services" \
	--grpc-gateway_out=services \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=import \
	proto/$$service/*.proto

clean:
	rm pb/*.go