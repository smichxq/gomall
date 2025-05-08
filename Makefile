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

# 简化从环境变量加载MySQL配置
# 仅用于开发
# 测试环境使用docker或其他安全的方式
# 环境变量是临时的且仅用于启动命令
.PHONY: demo-proto-server-boot-start
demo-proto-server-boot-start:
	@echo "Load MySQL config from env"
	@cd demo/demo_proto && \
	env MYSQL_USER=root \
	    MYSQL_PASSWORD=123 \
	    MYSQL_HOST=192.168.3.6 \
	    MYSQL_PORT=10123 \
	    MYSQL_DATABASE=demo_proto \
	go run .


# 数据填充CRUD
.PHONY: demo-proto-db-record-generate
demo-proto-db-record-generate:
	@echo "MySQL record generate"
	@cd demo/demo_proto/cmd/dbop && \
	env MYSQL_USER=root \
	    MYSQL_PASSWORD=123 \
	    MYSQL_HOST=192.168.3.6 \
	    MYSQL_PORT=10123 \
	    MYSQL_DATABASE=demo_proto \
	go run .

# 根据IDL生成Http项目
# cwgo命令解析
# --idl 指定参考的proto
# --service 服务名称
# -module 指定服务module名称(与目录一致)
# -I 指定参考的proto的依赖文件的文件夹
.PYTHON: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl && go work use . && go mod tidy