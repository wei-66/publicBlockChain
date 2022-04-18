package transaction

import (
	"btcionshow/utils"
	"bytes"
	"encoding/gob"
)

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

func (txs *Transaction) Serialize() ([]byte, error) {
	var result bytes.Buffer
	en := gob.NewEncoder(&result)
	err := en.Encode(txs)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil

}
//创建coinbase交易
func NewCoinbase(address string)(*Transaction ,error) {
	cb := Transaction{
		Output: []Output{
			{
				Value: 50,
				ScriptSig: []byte(address),
			},
		},
		Input: nil,
	}
	txsBytes, err := cb.Serialize()
	if err != nil {
		return nil ,err
	}
	hash := utils.GetHash(txsBytes)
	cb.TxId = hash
	return  &cb,nil
}