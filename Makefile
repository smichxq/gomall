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
.PYTHON: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/cloudwego/gomall/app/frontend -I ../../idl && go work use . && go mod tidy

# 根据IDL生成auth对应的代码
.PYTHON: gen-frontend-auth-page
gen-frontend-auth-page:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/cloudwego/gomall/app/frontend -I ../../idl && go work use . && go mod tidy

# 热重载
.PYTHON: hot-reload-run-forntend
hot-reload-run-forntend:
	@cd app/frontend && air

# 生成user客户端(idl)代码到rpc_gen文件夹下方便复用
.PYTHON: gen-user-rpc-client
gen-user-rpc-client:
	@ cd rpc_gen && cwgo client --type RPC --service user --module github.com/cloudwego/gomall/rpc_gen --I ../idl --idl ../idl/user.proto && go work use . && go mod tidy

# 生成user服务端(idl)代码到rpc_gen文件夹下方便复用
# --pass 向底层工具（hz 或 Kitex）传递额外的参数
# -use 配置kitex 不生成 kitex_gen 目录并使用指定的目录
# 生成完毕后目录依赖-use会去远程查找
# 请手动新增replace github.com/cloudwego/gomall/rpc_gen => ../../rpc_gen到app/user/go.mod 文件并刷新依赖
.PYTHON: gen-user-rpc-server
gen-user-rpc-server:
	@ cd app/user && cwgo server --type RPC --service user --module github.com/cloudwego/gomall/app/user --pass "-use github.com/cloudwego/gomall/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/user.proto && go work use . && go mod tidy


# 热重载
.PYTHON: hot-reload-run-user
hot-reload-run-user:
	@cd app/user && air



# 简化从环境变量加载MySQL配置
# 仅用于开发
# 测试环境使用docker或其他安全的方式
# 环境变量是临时的且仅用于启动命令
.PHONY: app-user-server-boot-start
app-user-server-boot-start:
	@echo "Load config from env"
	@cd app/user && \
	env MYSQL_USER=root \
	    MYSQL_PASSWORD=123 \
	    MYSQL_HOST=192.168.3.6 \
	    MYSQL_PORT=3306 \
	    MYSQL_DATABASE=user \
	air