package kekasigohelper

import "context"

func GetDataFromContext(ctx context.Context, key string) string {
	lr := ctx.Value(key)
	if l, ok := lr.(string); ok {
		return l
	} else {
		return AnyToString(lr)
	}
}
