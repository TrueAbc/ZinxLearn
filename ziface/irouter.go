package ziface

/*
	路由抽象接口
	路由的數據都是IRequest
*/

type IRouter interface {
	// 處理業務之前的Hook方法
	PreHandler(request IRequest)
	// 處理業務的主方法
	Handler(request IRequest)
	// 處理業務之後的方法
	PostHandler(request IRequest)
}
