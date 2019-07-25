# Go-Micro-Service

Simple Golang Micro Manager is useful to create RestFul web services using Golang.

## Features:

1. Environment setup (Like Development, QA, Production)
2. API Versioning
3. Middleware Authentication Layer
4. Supports for HTTP request methods (GET, PUT, POST, DELETE)
5. Host Web Pages
6. Code Coverage (Unit testing)


#### 1. Environment setup

Web Service basics have to support multiple environments like development, QA, Production environment and Configured in the generic format.
Under environments folder created three different config files to achieve the Environment setup.

```json
{
	"title" : "Golang Micro Web Service Manager (DEV)",
	"version" : "v1",
	"http": {
		"enable" : true,
		"port" : 3000	
	},
	"https": {
		"options" : {
			"key" : null,
			"cert" : null
		},
		"enable" : false,
		"port" : 3043
	},
	"auth" : "bypass"
}
```

#### 2. API Version

Configure the API version number based on the environments JSON file.

E.g., `DEV.json` is helpful achieve the API versioning to register web services based on the version number.

```golang
	http://<webserver>/<version>/<webservice>        
```

#### 3. Middleware Authentication Layer

Every RESTful Web services need the security and validation before going to hit the server business logic's. Go Micro Service Manager helps to implement the basic security and validate the request headers.

For the different environments, We can able to provide a different kind of security implementation using over security implementation.

	-   DEV environment is by-pass the security validation.
	-   QA and Production environments are required the security validation. 

If it's QA and Production environments, Request Header must be.

	- authorization = "bybass"
	- content-type  = "application/json"
	- date          = "1454577924104"     - UTC time.

#### 4. Supports for HTTP request methods (GET, PUT, POST, DELETE)    

Go Micro Service Manager helps to support the four types of HTTP request methods.


#### 5. Host Web pages    

Go Micro Service Manager helps to host static web pages in the HTTP server with dynamic web title based on the environments.

![alt text][api_test_environment]

[api_test_environment]: https://github.com/vigneshuvi/Go-Micro-Service/blob/master/public/images/api_test_environment.png

#### How to run this Go Micro Services server?

-   Need to install Go & Set the GOPATH
-   Open terminal and check the Go version
-   Run the Go HTTP server by 'go run main.go DEV' command (Like DEV, QA and PROD)

![alt text][start-go-http-server]

[start-go-http-server]: https://github.com/vigneshuvi/Go-Micro-Service/blob/master/public/images/start-go-http-server.png

#### How to run unit test and see the code coverage?

-   Open terminal and move to Go-Micro-Service folder
-	Run 'go test ./... -coverprofile=coverage.out' to generate coverage profile
-   Seeing coverage output in html file use 'go tool cover -html=coverage.out'

![alt text][code-coverage]

[code-coverage]: https://github.com/vigneshuvi/Go-Micro-Service/blob/master/public/images/code-coverage.png


#### How to test the RESTful web service?

-  If you set up the node environment is `DEV`, then you can directly test it from RESTful Client.

		http://localhost:3000/v1/user             - Show All -  Fetch all based on request.
		http://localhost:3000/v1/user?id="test"   - Show     -  Fetch based on request id. 
		http://localhost:3000/v1/user?id="test"   - PUT      -  Update based on request id. 
		http://localhost:3000/v1/user             - POST     -  Create based on request.
		http://localhost:3000/v1/user:id="test"   - DELETE   -  Update based on request id.


#### Do you like it?

Do you like this repo? Share it on Twitter, Facebook, Google+ or anywhere you like so that more of us can use it and help. Thanks!

Created by [Vignesh](http://vigneshuvi.github.io/) 

![alt text][logo]

[logo]: https://github.com/vigneshuvi/vigneshuvi.github.io/blob/master/favicon.ico/android-icon-48x48.png