package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/layer5io/meshery/internal/graphql/generated"
	"github.com/layer5io/meshery/internal/graphql/model"
)

func (r *mutationResolver) ChangeAddonStatus(ctx context.Context, selector *model.MeshType, targetStatus *model.Status) (*model.Status, error) {
	if selector != nil && targetStatus != nil {
		return r.changeAddonStatus(ctx)
	}

	return nil, ErrInvalidRequest
}

func (r *mutationResolver) ChangeOperatorStatus(ctx context.Context, targetStatus *model.Status) (*model.Status, error) {
	if targetStatus != nil {
		return r.changeOperatorStatus(ctx, targetStatus)
	}

	return nil, ErrInvalidRequest
}

func (r *queryResolver) GetAvailableAddons(ctx context.Context, selector *model.MeshType) ([]*model.AddonList, error) {
	if selector != nil {
		return r.getAvailableAddons(ctx, selector)
	}

	return nil, ErrInvalidRequest
}

func (r *queryResolver) GetControlPlanes(ctx context.Context, filter *model.ControlPlaneFilter) ([]*model.ControlPlane, error) {
	if filter != nil {
		return r.getControlPlanes(ctx)
	}

	return nil, ErrInvalidRequest
}

func (r *queryResolver) GetOperatorStatus(ctx context.Context) (*model.OperatorStatus, error) {
	return r.getOperatorStatus(ctx)
}

func (r *subscriptionResolver) ListenToAddonEvents(ctx context.Context, selector *model.MeshType) (<-chan []*model.AddonList, error) {
	if selector != nil {
		return r.listenToAddonEvents(ctx)
	}

	return nil, ErrInvalidRequest
}

func (r *subscriptionResolver) ListenToControlPlaneEvents(ctx context.Context, filter *model.ControlPlaneFilter) (<-chan []*model.ControlPlane, error) {
	if filter != nil {
		return r.listenToControlPlaneEvents(ctx)
	}

	return nil, ErrInvalidRequest
}

func (r *subscriptionResolver) ListenToOperatorEvents(ctx context.Context) (<-chan *model.OperatorStatus, error) {
	return r.listenToOperatorEvents(ctx)
}

func (r *subscriptionResolver) SubscribeToMeshSync(ctx context.Context) (<-chan *model.OperatorControllerStatus, error) {
	return r.subscribeToMeshSync(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
