package ctx

import "github.com/wailsapp/wails/v2/pkg/runtime"

func OpenFileDialog(dialogOptions runtime.OpenDialogOptions) (string, error) {
	return runtime.OpenFileDialog(ctx, dialogOptions)
}
