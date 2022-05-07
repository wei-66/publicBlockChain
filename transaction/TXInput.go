package transaction

import "bytes"

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/4/18 9:39
 **/
//交易输入结构体
type Input struct {
	TXid []byte //确定要消费的交易输出在哪个交易中

	Vout int //j交易输出的索引位置

	ScriptSig []byte //解锁脚本

}

func NewInput(txid []byte, vout int, scriptSig []byte) Input {
	return Input{txid, vout, scriptSig}
}

//判断input是某人的消费
func (input *Input) IsLocked(name string) bool {
	if name == "" {
		return false
	}
	return bytes.Compare(input.ScriptSig, []byte(name)) == 0
}
