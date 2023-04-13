* Go的`net/http`包已经实现了对 web 请求进行处理的功能。`http.ListenAndServe`可以对端口进行监听启动一个 web 服务，它的第二参数是`Handler`接口类型，所有的 http 请求都会被这个接口的`ServeHTTP`方法接收处理。  
* `Engine`实现了该接口，并内置路由器`router`，将请求方式和  请求路径组合起来作为唯一的 key ，将相应的处理函数储存在`router`中，将不同的 http 请求分发给对应的处理函数。
