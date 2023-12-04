## Templates-swagger-generation

### Pre-requisites

- Install Node.js and npm. Install **openapi** using below command
```shell
npm i -g @redocly/cli@1.0.0-beta.117
```
- Install **openapi-codegen** using below command
```shell
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest 
```

### Update the swagger-bundle.yaml

```shell
openapi bundle .\server.yaml --output swagger-bundle.yaml
```

### Generate code

```shell
oapi-codegen -package v1server -o server.generated.go -generate gorilla-server .\swagger-bundle.yaml
oapi-codegen -package v1server -o types.generated.go -generate types .\swagger-bundle.yaml
```
  