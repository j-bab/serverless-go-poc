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

run `go get ./...` in the project directory to fetch the go dependencies

## Configuration

configure AWS cli and serverless:  
https://www.serverless.com/framework/docs/providers/aws/guide/credentials/

In the project directory, update config.json and change the AwsAccountId to your own.  
Change the region if you wish.  
Change the serviceName to something relevant to you but likely to be unique, as it is used to generate bucket names which must be globally unique.

If you are NOT using Windows you will also need to adjust the Makefile as follows:  
  
Change the command for clean by replacing 	`rd /s /q  "./bin"` with `rm -rf ./bin`  
Change the build commands as follows:  
replace the line `	set GOOS=linux`  with `env GOOS=linux `   
replace the line `	set GOARCH=amd6`  with `env GOARCH=amd6 `   



## Deployment

In the project directory, run:

### `make deploy`

This cleans up the Bin dir, builds the Go Binaries, runs serverless deploy in verbose mode  
By default serverless will use the stage "dev" when deploying the stack

After a (long) while your cloudFormation will have been created and run, and if you look in your AWS console you should see your freshly created S3 buckets, dynamoDb table, Lambdas, Api Gateway and cloudFront distribution.
You should also see log groups in cloudWatch.


If the "bin" directory does not exist, it may cause an error when deploy is run as the clean operation will try to delete a non-existent directory.  
Simply create a directory called "bin" in the project root if this is the case (or run `make build`)


## Next steps
Infrastructure and API are now deployed.  

Now we can pull the [repository for the client app](https://github.com/j-bab/serverless-go-poc-client),
 configure it and deploy it to the app bucket we just created to serve from cloudFront
 
 To run the tests, simply run `go test`from the directory containing the tests to run.
 ### Thanks to
 
 https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/dynamodb/  

 https://github.com/nerdguru/go-sls-crudl/blob/master/functions/get.go
 
 https://www.softkraft.co/aws-lambda-in-golang/
 
 https://blog.alexellis.io/golang-writing-unit-tests/
 
and an intelli GoLand free trial 

and for CORS:
https://adamdrake.com/serverless-with-lambda-api-gateway-and-go.html