package item

import (
	"context"
	"github.com/dealmaker/procedure/item/model"
)

type WorkerInstance struct {
	FuncGetItem func(ctx context.Context, filter model.QueryFilter) ([]model.Item, error)
	FuncUpdateItem func(ctx context.Context, item *model.Item) error
	FuncInsertItem func(ctx context.Context, item *model.Item) (string, error)
	FuncDeleteItem func(ctx context.Context, objId string) error
}
func (w WorkerInstance) Init() *WorkerInstance {
	return &w
}