# Docker Commands

Basic generic set of commands to reference when working with docker

## Containers

### Create Container

```
$ sudo docker run --name mydb -e MYSQL_ROOT_PASSWORD=pass -d mysql:latest
```

### List Containers

```
$ docker ps
```

### Container Info

```
$ docker inspect <container name>
```

IP Address

```
$ docker inspect <container name> | grep IPAddress
```

### Docker Compose

Spin up docker instancres from a `docker-compose.yml` file.

```
$ sudo docker-compose up -d
```
`-d` *predicate detaches the docker process*

### Misc

Run against a MySQL instance

```
$ mysql -uroot -ptest-pass -h 172.17.0.2 -P 3306
```