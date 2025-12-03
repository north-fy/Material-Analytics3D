package application

import (
	"fyne.io/fyne/v2"
)

type ScreenManager struct {
	window        fyne.Window
	screens       map[screenName]fyne.CanvasObject
	currentScreen screenName
	onSwitch      func(screenName)
}

func NewScreenManager(window fyne.Window) *ScreenManager {
	return &ScreenManager{
		window:  window,
		screens: make(map[screenName]fyne.CanvasObject),
	}
}

func (sm *ScreenManager) addScreen(name screenName, obj fyne.CanvasObject) {
	sm.screens[name] = obj
}

func (sm *ScreenManager) setCurrentScreen(name screenName) {
	sm.currentScreen = name
	sm.window.SetContent(sm.screens[name])
}

func (sm *ScreenManager) setOnSwitch() {
	sm.onSwitch = func(name screenName) {
		sm.window.SetTitle(string(name))
	}
}
