package transaction

import "bytes"

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/4/18 9:37
 **/
//交易输出的结构体
type Output struct {
	Value uint //描述输出币的金额

	ScriptSig []byte //锁定脚本
}

//判断某人是否能解开交易输出(判断这笔钱是不是某人的)
func (output *Output) IsUnlock(name string) bool {
	if name == "" {
		return false
	}
	return bytes.Compare(output.ScriptSig, []byte(name)) == 0
}
