package main

import "C"

import (
	"github.com/daspoet/gowinkey"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"github.com/moutend/go-hook/pkg/mouse"
	"github.com/moutend/go-hook/pkg/types"
	"os"
	"time"
)

func Kill_Process() {
	kEvent := Usedkeys(gowinkey.VK_K)
	for e := range kEvent {
		if e.HasControl() && e.State == gowinkey.KeyDown {
			os.Exit(1)
		}
	}
}

// LOX
func Cheat_Clicker() {
	go Clicker()
	kEvent := Usedkeys(gowinkey.VK_Y)
	for e := range kEvent {
		if e.HasControl() && e.State == gowinkey.KeyDown {
			clickerState = !clickerState
		}
	}
}
func Clicker() {
	//intVar, err := strconv.Atoi(value)
	//fmt.Println (intVar, err, reflect.TypeOf(intVar))
	for {
		if clickerState {
			robotgo.MilliSleep(iDelay)
			robotgo.Click()
		}
	}
}
func Cheat_Burstmacro() {
	go MouseListening()
	kEvent := Usedkeys(gowinkey.VK_B)
	for e := range kEvent {
		if e.HasControl() && e.State == gowinkey.KeyDown {
			recoilMode = !recoilMode
		}
	}
}
func MouseListening() {
	go FakeClickLoop()
	mouseChan := make(chan types.MouseEvent, 100)

	if err := mouse.Install(nil, mouseChan); err != nil {
		return
	}

	for {
		mEvents := <-mouseChan
		bFlagStatus := mEvents.Flags & 1

		if recoilMode {
			if mEvents.Message == win.WM_LBUTTONDOWN {
				if bFlagStatus == 1 {
					continue
				}
				recoilState = true
			} else if mEvents.Message == win.WM_LBUTTONUP {
				if bFlagStatus == 1 {
					continue
				}
				recoilState = false
			}
		}

	}
}
func FakeClickLoop() {
	for {
		if recoilState {
			SendFakeClick()
		}
	}
}

func SendFakeClick() {
	SendMouseLeftDownKey()
	time.Sleep(58 * time.Millisecond)
	SendMouseLeftUpKey()
	time.Sleep(40 * time.Millisecond)
}
