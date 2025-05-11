package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Window"
	"graphics.gd/variant/Object"
)

type FullScreen struct {
	Node.Extension[FullScreen] `gd:"FullScreenHandler"`
}

func NewFullScreenHandler() *FullScreen {
	handler := new(FullScreen)
	handler.AsNode().SetProcessMode(Node.ProcessModeAlways)
	return handler
}

func (fs FullScreen) Input(event InputEvent.Instance) {
	if OS.HasFeature("HTML5") {
		if Object.Is[InputEventMouseButton.Instance](event) && Input.MouseMode() != Input.MouseModeCaptured {
			Input.SetMouseMode(Input.MouseModeCaptured)
		}
	} else {
		if key, ok := Object.As[InputEventKey.Instance](event); ok && event.IsPressed() && key.Keycode() == Input.KeyF11 {
			var window Window.Instance = SceneTree.Get(fs.AsNode()).Root()
			if window.Mode() == Window.ModeFullscreen {
				window.SetMode(Window.ModeWindowed)
			} else {
				window.SetMode(Window.ModeFullscreen)
			}
		}
	}
}
