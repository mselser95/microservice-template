# Microservice

## Install protoc

To install protoc:

```bash
brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```
## Generate GQL actions
1. Update Schema
```bash
gq https://<hasura url> -H "X-Hasura-Admin-Secret: <admin secret>" --introspect > schema.graphql
```
2. Write Query in `internal/gql/queries` and `internal/gql/mutations`
3. Add filename to `internal/gql/genqclient.yaml`
4. Run Genqlient
```bash
genqlient
```
This will generate the file `internal/gql/generated.go` with the new queries and mutations.