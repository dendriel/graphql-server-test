## Passos

- Inicializar projeto
- Adicionar tools.go do site do `gqlgen.com`
  - Rodar `go mod tidy`
- Inicializar tool do gqlgen
  - `go run github.com/99designs/gqlgen init`
  - Aqui já conseguimos subir o server com o schema de exemplo
- Criar schema.graphqls
- Rodar geração do gqlgen
  - `go run github.com/99designs/gqlgen generate`
- Apagar funções relacionadas às entidades TODO dentro de `schema.resolvers.go`




## Referências
- https://gqlgen.com/
