# Go Login Application That Scales on AWS
This repository is based on the Frontend Masters course ["Build Go Apps That Scale on AWS"](https://frontendmasters.com/courses/go-aws/) with which I created my first go project (login/register) using a AWS Infrastructure.

## Functionality
There is an API Gateway built in front of the lambda function. The lambda can handle:
- `/register`: saves the user with a hashed password in DynamoDB
- `/login`: compares the plain text password with the hashed one from DynamoDB
- `/protected`: is an example for a page that just a logged in user with a valid session (which gets checked via JWT Tokens) can access

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
