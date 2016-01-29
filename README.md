Hello Server [![Build Status](https://travis-ci.org/breerly/hello-server.svg?branch=master)](https://travis-ci.org/breerly/hello-server)
============

A Programmable, 6mb, Hello World Docker Appliance.


Run with default settings:

```
$ docker run -d -p 8080:8080 breerly/hello-server
$ curl `docker-machine ip default`:8080

Hello World!
```

Customize with environmental variables:

```
$ docker run -d -p 8081:8081 -e HELLO_MESSAGE="Hello Johnson!" -e HELLO_PORT=8081 breerly/hello-server
$ curl `docker-machine ip default`:8081

Hello World!
```

Run from Docker Compose:

```
hello:
    image: breerly/hello-server
    ports:
        - 8082:8082
    environment:
        - HELLO_PORT=8082
        - HELLO_MESSAGE=hello!
```

Note the this image exposes port 8000-8050, so the internal port chosen with `HELLO_PORT` should be within that range.

