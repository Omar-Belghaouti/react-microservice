rundb:
	@echo "Running database..."
	@docker run --name pg -e POSTGRES_PASSWORD=ramo -e POSTGRES_USER=omar -p 5432:5432 -d postgres:14-alpine
	@echo "Done"

createdb_posts:
	@echo "Creating database..."
	@docker exec -it pg createdb posts -U omar
	@echo "Done"

createdb_comments:
	@echo "Creating database..."
	@docker exec -it pg createdb comments -U omar
	@echo "Done"

run_posts_microservice:
	cd posts; go run main.go

run_comments_microservice:
	cd comments; go run main.go

run_frontend:
	cd frontend; yarn start

.PHONY: rundb createdb_posts createdb_comments run_posts_microservice run_comments_microservice run_frontend