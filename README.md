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


## Relacionamentos

Podemos realizar o relacionamento entre as entidades manualmente, direto nos resolvers, ou cadastrando novos models em `gqlden.yml`.
No segundo caso é necessário rodar o comando `go run github.com/99designs/gqlgen generate` que adicionará um novo tipo de resolver
dentro de `schema.resolvers.go`.

Por exemplo, se queremos retornar os cursos relacionados a uma categoria, devemos adicionar os models de curso e cateogria no `gqlden.yml`,
remover o relacionamento explícito da categoria com os cursos (a prop navigacional `[]Courses`), regerar os arquivos e implementar o novo resolver
de cursos associado ao novo CategoryResolver.

```golang
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: Courses - courses"))
}
```



## Referências
- https://gqlgen.com/
