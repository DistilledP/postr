IMAGE_NAME=test:builder

build-app:
	docker build -f build/package/Dockerfile --force-rm -t ${IMAGE_NAME} .

run-app:
	docker run --rm -it --name app ${IMAGE_NAME} sh

run-fmt:
	go fmt ./...

run-test:
	go test ./...

update-proto: build-app clean-proto
	@-docker run -v ${PWD}:/app --name proto_compile ${IMAGE_NAME} /app/scripts/compile_proto.sh
	@-docker cp proto_compile:/app/internal/proto ./
	@docker rm -f proto_compile
	@make prune

clean-proto:
	@if [ -d internal/proto ]; then rm -fr internal/proto; fi

prune:
	@docker image prune -f
