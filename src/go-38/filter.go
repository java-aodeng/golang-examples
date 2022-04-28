package go_38

// Request 入参定义
type Request interface{}

// Response 响应参数定义
type Response interface{}

// Filter interface is the definition of the data processing components
// Pipe-Filter structure
type Filter interface {
	Process(data Request) (Response, error)
}
