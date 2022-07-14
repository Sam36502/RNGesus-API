EXE_DIR := ./dist/
EXE := rngesus
IMAGE := rngesus-api
CONTAINER := rngesus
PORT := 777

build:
	@echo " > Building..."
	@go build -o $(EXE_DIR)$(EXE) main.go

run: build
	@echo " > Hosting..."
	@sudo $(EXE_DIR)$(EXE)

image: build
	@echo " > Building docker image..."
	@docker build -t $(IMAGE) .

up: image
	@echo " > Starting container..."
	@docker run --name $(CONTAINER) -p $(PORT):$(PORT) -d $(IMAGE)

logs:
	@echo " > Following container logs:"
	@docker logs -f $(CONTAINER)

down:
	@echo " > Stopping container..."
	@-docker stop $(CONTAINER)
	@echo " > Removing container..."
	@-docker rm $(CONTAINER)