package znet

import "trueabc.top/zinx/ziface"

// 可以作爲其他router實現的基礎, 可以嵌入這個基類
type BaseRouter struct {
}

func (b *BaseRouter) PreHandler(request ziface.IRequest) {

}

func (b *BaseRouter) Handler(request ziface.IRequest) {

}

func (b *BaseRouter) PostHandler(request ziface.IRequest) {

}

func NewBaseRouter() *BaseRouter {
	return &BaseRouter{}
}
