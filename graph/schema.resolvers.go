package graph

import (
	"context"
	"fmt"

	"github.com/vyctor/fc-graphql/graph/model"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: CreateCourse - createCourse"))
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: Courses - courses"))
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
