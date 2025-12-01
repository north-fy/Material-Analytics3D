package application

import (
	"github.com/north-fy/Material-Analytics3D/internal/application/layout"
)

// init all windows in layout
func windowManager() *layout.SpecificWindow {
	auth := layout.NewAuthWindow()
	sw := layout.NewSpecificWindow(auth)
	return sw
}
