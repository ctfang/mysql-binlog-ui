package ctx

import "context"

var ctx context.Context

func SetCtx(c context.Context) {
	ctx = c
}
