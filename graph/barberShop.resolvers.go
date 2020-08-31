package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ken-mathenge/barber-shop-mgmt-system/graph/generated"
	"github.com/ken-mathenge/barber-shop-mgmt-system/graph/model"
)

func (r *mutationResolver) CreateEmployee(ctx context.Context, name string, gender string, idNumber int, position []string) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, id *string, timeIn time.Time, timeOut time.Time, service []string, pay int) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateService(ctx context.Context, service []string) (*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employees(ctx context.Context, employeeID string) ([]*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customers(ctx context.Context, customerID string) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Services(ctx context.Context) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
