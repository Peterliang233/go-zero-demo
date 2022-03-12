generate-sql-model:
	cd ./service/user/model; goctl model mysql ddl -src user.sql -dir . -c
generate-user-api:
	cd ./service/user/api; goctl api go -api user.api -dir .
curl:
	curl -i -X POST \
	http://localhost:8888/user/login \
	-H 'Content-Type: application/json' \
	-d '{ \
	"username":"666", \
	"password":"123456" \
	}'
run-user-api:
	cd service/user/api; go run user.go -f etc/user-api.yaml