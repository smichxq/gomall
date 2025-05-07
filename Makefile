.PHONY: gen-demo-proto
gen-demo-proto:
	@mkdir demo/demo_proto && cd demo/demo_proto && cwgo server --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto
	@echo "Generating demo proto files..."
	# @mkdir -p demo
	# @protoc --proto_path=demo --proto_path=../proto --proto_path=../third_party/protobuf/src --proto_path=../third_party/protobuf/src/google/protobuf --proto_path=../third_party/protobuf/src/google/api --proto_path=../third_party/protobuf/src/google/rpc --proto_path=../third_party/protobuf/src/google/type --proto_path=../third_party/protobuf/src/google/iam/v1 --proto_path=../third_party/protobuf/src/google/api/annotations.proto --go_out=. demo/demo.proto
	@echo "Demo proto files generated successfully."
	@echo "go module init"
	@go mod tidy
	@echo "go mod tidy success"
	@echo "add workspace"
	@go work use .
	@echo "add workspace success"


.PHONY: gen-demo-thrift
gen-demo-thrift:
	@mkdir demo/demo_thrift && cd demo/demo_thrift && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_proto --idl ../../idl/echo.proto
	@echo "Generating demo thrift files..."
	# @mkdir -p demo
	# @thrift --gen go -out demo demo.thrift
	@echo "Demo thrift files generated successfully."
	@echo "go module init"
	@cd demo/demo_thrift && go mod tidy
	@echo "go mod tidy success"
	@echo "add workspace"
	@cd demo/demo_thrift && go work use .
	@echo "add workspace success"