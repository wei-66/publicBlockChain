package transaction

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/4/18 9:41
 **/
//整个交易的结构体
type Transaction struct {

	TxId []byte //交易的唯一标识

	Output []Output //多个交易输出

	Input []Input //多个交易输入
}

