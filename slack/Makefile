.PHONY: clean build deploy

clean: 
	rm -rf ./kot/kot
	
build:
	GOOS=linux GOARCH=amd64 go build -o kot/kot ./kot

deploy: build
	sam package --template-file template.yaml --s3-bucket ${BUCKET} --output-template-file packaged.yaml
	aws cloudformation deploy --template-file packaged.yaml --stack-name kot-slack  --capabilities CAPABILITY_IAM --profile ${PROFILE}

