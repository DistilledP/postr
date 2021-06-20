IMAGE_NAME=postr:latest
ARTIFACTS_OUTPUT=build/artifacts

build-app:
	docker build -f build/package/Dockerfile \
		--build-arg BUILD_ARTIFACTS=${ARTIFACTS_OUTPUT} \
		--force-rm \
		-t ${IMAGE_NAME} .

	@make download-cmd
	@make prune

download-cmd:
	@if [ ! -d ./bin ]; then mkdir bin; fi
	@if [ -f ./bin/image-upload ]; then rm ./bin/image-upload; fi
	@-docker rm cmd_copy >/dev/null
	@docker create --name cmd_copy ${IMAGE_NAME}
	@docker cp cmd_copy:/app/image-upload ./bin
	@docker rm cmd_copy >/dev/null
	@echo "Cmd downloaded to ${PWD}/bin"

run-fmt:
	go fmt ./...

run-test:
	go test -cover ./...

update-proto: build-app clean-proto
	@-docker run --rm -v ${PWD}:/app --name proto_compile ${IMAGE_NAME} /app/scripts/compile_proto.sh
	@make prune

clean-proto:
	@if [ -d internal/proto ]; then rm -fr internal/proto; fi

prune:
	@docker image prune -f >/dev/null
