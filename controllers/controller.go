package controllers

type Handlers interface {
	CreateUser()
	ResizePicture()
	SecondResizePicture()
	GetRequestAllResizeObjs()
}
