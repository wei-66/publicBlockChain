package transaction

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/5/2 10:00
 **/
//用于描述用户可以消费得交易信息
//1.该交易输出在哪个交易中：txid
//2.该交易输出在交易中是第几个：index
//3.该交易输出属于谁
//4.该交易输出得金额
type UTXO struct {
	Txid    []byte
	Index   int
	*Output //匿名字段 utxo中结构体就会默认包含output结构体得字段
}

//创建UTXO结构体
func NewUTXO(txid []byte, index int, output *Output) UTXO {
	return UTXO{txid, index, output}
}
