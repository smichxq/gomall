.PHONY: gen-demo-proto
gen-demo-proto:
	@ echo "Generating demo proto project"
	@ mkdir -p demo/demo_proto && cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto
	@echo "Generating demo proto files..."
	@echo "Demo proto files generated successfully."
	@echo "go module init"
	@cd demo/demo_proto && go mod tidy
	@echo "go mod tidy success"
	@echo "add workspace"
	@cd demo/demo_proto && go work use .
	@echo "add workspace success"


.PHONY: gen-demo-thrift
gen-demo-thrift:
	@ echo "Generating demo thrift project"
	@mkdir -p demo/demo_thrift && cd demo/demo_thrift && cwgo server --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_proto --idl ../../idl/echo.thrift
	@echo "Generating demo thrift files..."
	@echo "Demo thrift files generated successfully."
	@echo "go module init"
	@cd demo/demo_thrift && go mod tidy
	@echo "go mod tidy success"
	@echo "add workspace"
	@cd demo/demo_thrift && go work use .
	@echo "add workspace success"