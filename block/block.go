package block

import (
	pow2 "btcionshow/pow"
	"btcionshow/utils"
	"bytes"
	"encoding/gob"
	"strconv"
	"time"
)

/**
* @author: SuZhiXiaoWei
* @DateTime: 2022/2/21 11:03
**/
//区块结构体
type Block struct {
	PrevHash  []byte //上一个区块hash
	TimeStamp int64  //时间戳
	Data      []byte //数据
	Nonce     int64  //随机数
	Hash      []byte //hash值
}
func (block *Block)GetTimeStamp()int64{
	return block.TimeStamp
}
func (block *Block)GetPrevHash()[]byte{
	return block.PrevHash
}
func (block *Block)GetData()[]byte{
	return block.Data
}
/*
1.获取当前区块的哈希值
2.由上一个区块hash值，数据data，和时间戳拼接 再加上随机数
*/
func (block *Block)getBlockHash()[]byte{
	//将int类型转为string类型
	time := []byte( strconv.FormatInt(block.TimeStamp, 10))
	random := []byte( strconv.FormatInt(block.Nonce, 10))
	//拼接字符串 第一个参数要拼接的内容，第二个是以什么形式拼接
	hash := bytes.Join([][]byte{block.PrevHash, block.Data, time,random}, []byte{})
	return utils.GetHash(hash)
}
//func (block *Block) SetHash() []byte {
//	//区块的hash ：时间戳+上一个区块的hash值+交易信息+随机数
//
//	time := []byte(strconv.FormatInt(block.TimeStamp, 10))
//	nonce := []byte(strconv.FormatInt(block.Nonce, 10))
//	hash := bytes.Join([][]byte{block.PrevHash, block.Data, time, nonce}, []byte{})
//	return utils.GetHash(hash)
//}
//创建新的区块
func NewBlock(prevHash []byte, data []byte) *Block {
	block := Block{
		PrevHash:  prevHash,
		TimeStamp: time.Now().Unix(),
		Data:      data,
	}
	pow := pow2.NewPow(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block

}

//序列化  把目标结构体转成一个有序的排序
//序列化：把结构体block转成[]byte
func (block *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	en := gob.NewEncoder(&result)
	err := en.Encode(block)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil

}

//把[]byte转成block
//反序列化
func DeSerialize(data []byte) (*Block, error) {
	//var result bytes.Buffer
	reader := bytes.NewReader(data)
	de := gob.NewDecoder(reader)
	var block *Block
	err := de.Decode(&block)
	if err != nil {
		return nil, err
	}
	return block, nil
}


//创世区块
func GenesisBlock(data []byte) *Block {
	return NewBlock(nil, data)
}




