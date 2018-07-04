# docker image build
```
docker build -t pyapp:latest .
```

# docker 컨테이너 실행
```
docker run --name pyapp-test -p 8080:5000 -v $(pwd):/python-docker pyapp
```

# docker 접속..?
```
$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS PORTS               NAMES
fddaad0145d6        f04e40655008        "/bin/sh -c 'python3..."   9 minutes ago       Up 2 minutes     0.0.0.0:8080->5000/tcp      pyapp-test

$ docker exec -it $container_id
```

참고 URL : https://mingrammer.com/setup-the-python-development-environment-with-pycharm-and-docker/
