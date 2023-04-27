package graph

import "github.com/andreis3/graphql-goexpert/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryService *database.Category
}
