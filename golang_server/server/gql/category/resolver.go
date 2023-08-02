package category

import (
	itf "com.fullstack.ecommerce/server/gql/generated"
)

type CategoryQuery struct {
	categoryRootQuery
}

type CategoryMutation struct {
	categoryRootMutation
}

type CategoryNamespace struct{}

func (CategoryNamespace) Category() itf.CategoryResolver {
	return &categoryType{}
}
