package transaction

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


