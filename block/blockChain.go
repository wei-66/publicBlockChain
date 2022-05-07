package block

import (
	"btcionshow/transaction"
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/2/21 10:34
 **/

const CHAIN_DB_PATH = "./chain.db"   //保存的文件地址
const BUCKET_BLOCK = "chain_block"   //存区块桶的名称
const BUCKET_STATUS = "chain_status" //用来存最后一个区块的hash值
const LAST_HASH = "lasthash"         //存桶2的key的名称

//区块链管理区块
type BlockChain struct {
	//Blocks []*Block
	DB       *bolt.DB
	LastHash []byte
}

//创建区块链
func NewChain(address string) (*BlockChain, error) {
	//打开数据库
	db, err := bolt.Open(CHAIN_DB_PATH, 0600, nil)
	if err != nil {
		return nil, err
	}
	var lastHash []byte
	//向数据库中添加数据
	//同一个时间内，只能有一个人来进行写操作
	err = db.Update(func(tx *bolt.Tx) error {
		//1.有桶
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil {
			//创建一个coinbase交易
			coinbase, err := transaction.NewCoinbase(address)
			genesis := GenesisBlock(*coinbase)
			bk, err := tx.CreateBucket([]byte(BUCKET_BLOCK))
			if err != nil {
				return err
			}
			serialize, err := genesis.Serialize()
			if err != nil {
				return err
			}
			bk.Put(genesis.Hash, serialize)
			//创建桶2
			bk2, err := tx.CreateBucket([]byte(BUCKET_STATUS))
			bk2.Put([]byte(LAST_HASH), genesis.Hash)
			lastHash = genesis.Hash
		} else {
			bk2 := tx.Bucket([]byte(BUCKET_STATUS))
			lastHash = bk2.Get([]byte(LAST_HASH))
		}
		return nil
	})
	bc := BlockChain{
		DB:       db,
		LastHash: lastHash,
	}

	return &bc, err
}

//把区块添加到区块链中
func (bc *BlockChain) AddBlock(tx []transaction.Transaction) error {

	new := NewBlock(bc.LastHash, tx)
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil {
			return errors.New("没有创建桶")
		}
		serialize, err := new.Serialize()
		if err != nil {
			return err
		}
		bk.Put(new.Hash, serialize)
		bk2 := tx.Bucket([]byte(BUCKET_STATUS))
		if bk2 == nil {
			return errors.New("没有创建桶2")
		}
		bk2.Put([]byte(LAST_HASH), new.Hash)
		bc.LastHash = new.Hash
		return nil
	})
	return err
}

//创建一个迭代器对象,迭代器只能在有区块链的情况下才可以使用迭代器
func (bc *BlockChain) Iterator() *ChainIterator {
	iterator := ChainIterator{
		DB:          bc.DB,
		currentHash: bc.LastHash,
	}

	return &iterator
}

//获取所有区块
func (bc *BlockChain) GetAllBlock() []*Block {
	iterator := bc.Iterator()
	bk := []*Block{}
	for {
		if iterator.HashNext() {
			bl, err := iterator.Next()
			if err != nil {
				fmt.Println("失败了")
				return nil
			}
			bk = append(bk, bl)
		} else {
			break
		}

	}
	return bk
}

//定义方法用于查找某个地址的所有收入
func (bc *BlockChain) FindAllOutput(address string) []transaction.UTXO {
	//先找所有的区块，在获取每一个区块中的所有交易，再找是address的收入
	blocks := bc.GetAllBlock()
	//map结构：map[key]value
	//key: 收入(input)所在的交易hash
	//value:[]int 表示output的位置下标
	allOutputs := make([]transaction.UTXO, 0)
	//获取每一个区块
	for _, block := range blocks {
		//获取每个区块中的每个交易
		for _, tx := range block.Txs {
			//找每一个交易中的所有的交易输出
			for outIndex, output := range tx.Output {
				if output.IsUnlock(address) {
					utxo := transaction.NewUTXO(tx.TxId, outIndex, &output)
					allOutputs = append(allOutputs, utxo)
				}
			}
		}
	}
	return allOutputs
}

//寻找某个人的所有的消费(input)
func (bc *BlockChain) FindAllInput(name string) []transaction.Input {
	//先找所有的区块，在获取每一个区块中的所有交易，再找是address的所有消费
	allInputs := make([]transaction.Input, 0)
	blocks := bc.GetAllBlock()
	//获取每一个区块
	for _, block := range blocks {
		//获取每个区块中的每个交易
		for _, tx := range block.Txs {
			//找每一个交易中的所有的input
			for _, input := range tx.Input {
				if input.IsLocked(name) {
					allInputs = append(allInputs, input)
				}
			}
		}
	}
	return allInputs
}

//接收某个地址的所有收入和消费记录，并从所有的收入中取出消费，剩下的就是可用交易输出(余额)
func (bc *BlockChain) FindSpendOutputs(alloutputs []transaction.UTXO,
	allInputs []transaction.Input, amount uint) []transaction.UTXO {
	//获取每一个交易收入
	for _, input := range allInputs {
		for index, utxo := range alloutputs {
			//比较输出和输入得txid和下标是否相等,相等则说明被消费了
			if bytes.Compare(utxo.Txid, input.TXid) == 0 || utxo.Index == input.Vout {
				//要从utxo中去掉这笔收入，剩下的utxo就是可用余额
				alloutputs = append(alloutputs[0:index], alloutputs[index+1:]...)
				break
			}
		}
	}
	var totalAmount uint = 0
	//记录一共有多少笔utxo
	outputs := make([]transaction.UTXO, 0)
	//遍历找到所有的余额
	for _, output := range alloutputs {
		totalAmount += output.Value
		outputs = append(outputs, output)
		//当金额累计到够用时就可用不用累加了
		if totalAmount >= amount {
			break
		}
	}
	return alloutputs
}
