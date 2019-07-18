# godocker

Repository containing a simple docker container, that runs a go server which listens on port 8080 (of host machine) and prints Hello World.


## Build godocker
Browse the folder cloned, then run:

    go build main.go
    
Or:

    go build .

This will generate a godocker executable.


## Create & run Docker image
Note: no need to build godocker executable before creating the container. The container builds the executable during it own building.

    # build docker image
    docker build -t godocker .

    # run godocker container, map host port 8080 to container port 8080
    docker run --rm -p 8080:8080 godocker


