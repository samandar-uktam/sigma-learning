package abstractfactory

type GUIFactory interface {
	CreateButton() Button
	CreateTextField() TextField
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateTextField() TextField {
	return &MacTextField{}
}

type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (f *WinFactory) CreateTextField() TextField {
	return &WinTextField{}
}
