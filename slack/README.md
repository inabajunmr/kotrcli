# slack

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Deploy

### Config

You must set environment variables VERIFICATION_TOKEN and SLACK_ID on template.yaml.

* VERIFICATION_TOKEN=Slack app token
* SLACK_ID=Your slack id
* USER_TOKEN=Your KOT user_token
* TOKEN=Your KOT token

#### Example
```yaml
Variables:
  VERIFICATION_TOKEN: xxxxx
  SLACK_ID: "USER_TOKEN:TOKEN"
```

### Deploy Command
```
make deploy BUCKET="your S3 bucket name" PROFILE="your aws profile"
```
