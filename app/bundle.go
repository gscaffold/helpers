package app

import "context"

type defaultBundle struct {
	name   string
	fn     func(context.Context)
	stopFn func(context.Context)
}

func NewDefaultBundle(name string, fn, stopFn func(context.Context)) IBundle {
	return &defaultBundle{
		name:   name,
		fn:     fn,
		stopFn: stopFn,
	}
}

func (b *defaultBundle) GetName() string {
	return b.name
}
func (b *defaultBundle) Run(ctx context.Context) {
	b.fn(ctx)
}

func (b *defaultBundle) Stop(ctx context.Context) {
	b.stopFn(ctx)
}
