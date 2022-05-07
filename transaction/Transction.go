package transaction

import (
	"btcionshow/block"
	"btcionshow/utils"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
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

//创建一个普通的交易
func NewTransaction(from, to string, amount uint) (*Transaction, error) {
	//1.构建Input交易输入
	//1.1在所有的交易中，去寻找可以使用
	//所有可用的交易输出(余额) = 所有的收入 - 所有的消费
	bc, err := block.NewChain("")
	if err != nil {
		return nil, nil
	}
	input := bc.FindAllInput(from)
	output := bc.FindAllOutput(from)
	//获取了from这个地址的所有可用的交易输出
	//因为只需要找到满足本次交易的余额就可以了
	spendableOutputs := bc.FindSpendOutputs(output, input, amount)
	//1.2从所有可用的交易输出中，取出一部分，判断是否够用
	if spendableOutputs == nil {
		return nil, errors.New("没有可用的余额")
	}
	inputs := make([]Input, 0)
	for _, output := range spendableOutputs {
		//构建交易输入，就要引入交易输出，因为交易输入本质就是之前历史的交易中未消费的交易输出
		input := NewInput(output.Txid, output.Index, []byte(from))
		inputs = append(inputs, input)
	}
	//2.构建Output交易输出
	//3.给TxId赋值

	tx := Transaction{
		Output: nil,
		Input:  inputs,
	}
	txBytes, err := tx.Serialize()
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	hash := utils.GetHash(txBytes)
	tx.TxId = hash
	return &tx, nil
}

//创建coinbase交易
func NewCoinbase(address string) (*Transaction, error) {
	cb := Transaction{
		Output: []Output{
			{
				Value:     50,
				ScriptSig: []byte(address),
			},
		},
		Input: nil,
	}
	txsBytes, err := cb.Serialize()
	if err != nil {
		return nil, err
	}
	hash := utils.GetHash(txsBytes)
	cb.TxId = hash
	return &cb, nil
}

//序列化
func (txs *Transaction) Serialize() ([]byte, error) {
	var result bytes.Buffer
	en := gob.NewEncoder(&result)
	err := en.Encode(txs)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil

}
