package abstractfactory

import "fmt"

type MacButton struct{}

func (m *MacButton) Render() {
	fmt.Println("Render Mac Button")
}

type WinButton struct{}

func (m *WinButton) Render() {
	fmt.Println("Render Win Button")
}

type MacTextField struct{}

func (mtx *MacTextField) Render() {
	fmt.Println("Render Mac Text Field")
}

type WinTextField struct{}

func (wintx *WinTextField) Render() {
	fmt.Println("Render Win Text Field")
}
