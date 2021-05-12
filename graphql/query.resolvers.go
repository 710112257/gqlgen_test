package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
)

func (r *queryResolver) Article(ctx context.Context) (*Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("hello,%v", name), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
