generate-sql-model:
	cd ./service/user/model; \
	goctl model mysql ddl -src user.sql -dir . -c
generate-user-api:
	cd ./service/user/api; \
	goctl api go -api user.api -dir .
generate-search-api:
	cd ./service/search/api; \
	goctl api go -api search.api -dir .
run-user-api:
	cd service/user/api; \
	go run user.go -f etc/user-api.yaml
run-search-api:
	cd service/search/api; \
	go run search.go -f etc/search-api.yaml
generate-user-rpc:
	cd service/user/rpc; \
	goctl rpc proto -src user.proto -dir .
run-user-rpc:
	cd service/user/rpc; \
	go run user.go -f etc/user.yaml
