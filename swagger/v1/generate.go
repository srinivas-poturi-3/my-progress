//go:generate oapi-codegen -package v1server -o server.generated.go -generate gorilla-server .\swagger-bundle.yaml
//go:generate oapi-codegen -package v1server -o types.generated.go -generate types .\swagger-bundle.yaml

package v1server
