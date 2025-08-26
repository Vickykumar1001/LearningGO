migrate:
	 echo "Running migrations"
	 cd pkg && \
	 go run migration.go

start-mongo:
	  go run main.go --dbtype=mongodb
start-mysqlx:
	  go run main.go --dbtype=mysqlx
start-sqlORM:
	  go run main.go --dbtype=sqlORM

mongo:
		docker run -d \
		  --name mongo-container \
		  -p 27017:27017 \
		  mongo
