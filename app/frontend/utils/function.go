package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) int32 {
	useId := ctx.Value(SessionUserId)
	if useId == nil {
		return 0
	}
	return useId.(int32)
}