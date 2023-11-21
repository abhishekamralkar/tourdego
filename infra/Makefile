build:
	go build -o infra .

run:
	./infra

dockerbuild:
	sudo docker build --target runner . -t infra

dockerrun:
	sudo docker run infra

dockerdeleteimage:
	sudo docker rmi -f $(sudo docker images | awk '{print $3}' | tail -n +2)
        
all: build run dockerbuild dockerrun
