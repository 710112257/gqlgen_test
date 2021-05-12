package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/710112257/gqlgen_test/model"
)

func (r *queryResolver) Article(ctx context.Context) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) People(ctx context.Context) (*People, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SayHello(ctx context.Context, name string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
