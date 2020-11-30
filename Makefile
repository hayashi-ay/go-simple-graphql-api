up: 
	docker-compose up -d

down:
	docker-compose down

gqlgen: ## GraphQLのボイラープレートを作成する
	go run github.com/99designs/gqlgen
