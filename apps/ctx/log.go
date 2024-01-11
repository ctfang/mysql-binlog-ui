package ctx

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func LogDebug(msg string) {
	if ctx == nil {
		fmt.Println(msg)
		return
	}
	runtime.LogDebug(ctx, msg)
}

func LogError(msg string) {
	if ctx == nil {
		fmt.Println(msg)
		return
	}
	runtime.LogError(ctx, msg)
}
