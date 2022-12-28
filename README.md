
start up database
```
$ docker run \
	--detach \
	--name=snipp-db \
	--env="MYSQL_ROOT_PASSWORD=password" \
	--publish 3306:3306 \
	--volume=/root/docker/snipp-db/conf.d:/etc/mysql/conf.d \
	mysql:latest

$ mysql -u root -h 127.0.0.1 -p
```
