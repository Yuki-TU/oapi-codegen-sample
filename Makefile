# makeを打った時のコマンド
.DEFAULT_GOAL := help

.PHONY: gencode
# gencode: ## コード自動生成
# 	oapi-codegen -package gen -generate types -o gen/types.gen.go ./docs/openapi.yml
# 	oapi-codegen -package gen -generate strict-server,gin -templates ./_tool/oapi/templates -o gen/server.gen.go ./docs/openapi.yml
gencode: openapi ## コード自動生成
	oapi-codegen -package gen -generate types -o gen/types.gen.go ./docs/openapi.yml
	oapi-codegen -package gen -generate strict-server,gin -templates ./_tools/oapi/templates -o gen/server.gen.go ./docs/openapi.yml

.PHONY: openapi
openapi: ## OpenAPI bundle
	redocly bundle ./docs/index.openapi.yml --output ./docs/openapi.yml

new-mg: ## 新規作成
	sql-migrate new -config=./_tools/sql-migrate/config.yaml -env=development $(name)

mg-up: ## マイグレーション実行
	# sql-migrate up -config=./_tools/sql-migrate/config.yaml -env=development
	sql-migrate up -config=./_tools/sql-migrate/config.yaml -env=test

mg-down: ## マイグレーション戻す
	sql-migrate down -config=./_tools/sql-migrate/config.yaml -env=development

.PHONY: sqlc
sqlc:
	sqlc generate -f ./_tools/sqlc/config.yaml


.PHONY: mock
mock: ## mock作成
	mockgen -source=./handler/usecase.go -destination=./handler/_mock/mock_usecase.go
	mockgen -source=./repository/repository.go -destination=./repository/_mock/mock_repository.go
	mockgen -source=./repository/querier.go -destination=./repository/_mock/mock_querier.go
	mockgen -source=./repository/db.go -destination=./repository/_mock/mock_db.go

.PHONY: test
test: 
	go test -cover -race -shuffle=on ./...

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
