package model

type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	// Relation created via graphql resolvers
	//Courses     []*Course `json:"courses"`
}
