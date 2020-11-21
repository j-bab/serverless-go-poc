# Serverless Infrastructure and API 

A proof of concept Api written in Go, and the serverless configuration to deploy it and provision the necessary infrastructure in a CloudFormation stack.

## Getting started

You'll need the following installed on your system:

Node  
Go  
Aws Cli  
Serverless Framework

Assuming you use windows, you'll need to install make as well (with chocolatey: choco install make)

run `npm install` in the project directory

## Configuration

configure AWS cli and serverless:  
https://www.serverless.com/framework/docs/providers/aws/guide/credentials/

In the project directory, update config.json and change the AwsAccountId to your own.  
Change the region if you wish.  
Change the serviceName to something relevant to you but likely to be unique, as it is used to generate bucket names which must be globally unique.

If you are NOT using Windows you will also need to adjust the Makefile as follows:  
  
Change the command for clean by replacing 	`rd /s /q  "./bin"` with `rm -rf ./bin`  
Change the build commands as follows:  
remove the line `	set GOOS=linux`  
add `env GOOS=linux ` to the beginning of each remaining line in this section.
it should resemble something like this:  
`env GOOS=linux go build -ldflags="-s -w" -o bin/note functions/note/main.go`

## Deployment

In the project directory, run:

### `make deploy`

This cleans up the Bin dir, builds the Go Binaries, runs serverless deploy in verbose mode  
By default serverless will use the stage "dev" when deploying the stack

After a (long) while your cloudFormation will have been created and run, and if you look in your AWS console you should see your freshly created S3 buckets, dynamoDb table, Lambdas, Api Gateway and cloudFront distribution.
You should also see log groups in cloudWatch.

## Next steps
Infrastructure and API are now deployed.  

Now we can pull the [repository for the client app](https://github.com/j-bab/serverless-go-poc-client), configure it and deploy it to the app bucket we just created to serve from cloudFront