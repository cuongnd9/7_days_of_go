docker run -it \
--name database \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=103cuong \
-e MYSQL_DATABASE=grpc_go \
-v ${PWD}/scripts/database/sql-scripts:/docker-entrypoint-initdb.d/ \
mysql:8
