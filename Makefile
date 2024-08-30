db-up:
	docker run --rm --name my_postgres \
        	-e POSTGRES_HOST_AUTH_METHOD=trust \
        	-e POSTGRES_USER=postgres \
        	-e POSTGRES_DB=postgres \
        	-p 5432:5432 -d postgres:14.3