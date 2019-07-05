# Go-Micro-Service
Simple Golang Micro Manager is useful to create RestFul web services using Golang.

## Features:

1. Environment setup (Like Development, QA, Production)
2. API Versioning (WIP)
3. Middleware configuration (WIP)
4. Basic Security (WIP)
5. Supports for HTTP request methods (GET, PUT, POST, DELETE)


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

Configure the API version numbe based on the environments JSON file.

E.g., `DEV.json` is helpful achieve the API versioning to register web services based on the version number.

```golang
    http://<webserver>/<version>/<webservice>        
```