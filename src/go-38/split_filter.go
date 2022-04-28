package go_38

import (
	"errors"
	"strings"
)

// SplitFilterWrongFormatError 定义异常信息
var SplitFilterWrongFormatError = errors.New("input data should be string")

// SplitFilter 定义分割符号
type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) //检查数据格式/类型，是否可以处理
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	//按照分割符号分割传入的参数
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
