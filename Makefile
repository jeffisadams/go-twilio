STACK_NAME := "twilio-stack"

.PHONY: test
test:
	aws cloudformation validate-template --template-body file://template.yaml

.PHONY: clean
clean:
	rm -rf ./dist
	rm -rf template_deploy.yaml

.PHONY: deps
deps: clean
	go get github.com/kevinburke/twilio-go
	go get github.com/aws/aws-sdk-go/aws

.PHONY: build
build: deps
	go build -o dist/main ./src/main.go

.PHONY: buildPi
buildPi: deps
	GOOS=linux GOARCH=arm GOARM=5 go build -o dist/sendText ./src/main.go

.PHONY: deploy
deploy:
	aws cloudformation deploy \
		--no-fail-on-empty-changeset \
		--template-file template.yaml \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_AUTO_EXPAND CAPABILITY_IAM \
		--parameter-overrides  "Bucket=$(BUCKET_NAME)"

.PHONY: teardown
teardown:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
	clean
