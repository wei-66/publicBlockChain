package block

import (
	"btcionshow/transaction"
	"fmt"
	"github.com/boltdb/bolt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/2/21 10:34
 **/

const CHAIN_DB_PATH = "./chain.db"    //保存的文件地址
const BUCKET_BLOCK = "chain_block"   //存区块桶的名称
const BUCKET_STATUS = "chain_status" //用来存最后一个区块的hash值
const LAST_HASH = "lasthash"           //存桶2的key的名称

//区块链管理区块
type BlockChain struct {
	//Blocks []*Block
	DB       *bolt.DB
	LastHash []byte
}
//创建区块链
func NewChain(address string)(*BlockChain,error){
	//打开数据库
	db, err := bolt.Open(CHAIN_DB_PATH, 0600, nil)
	if err !=nil{
		return nil,err
	}
	var lastHash []byte
	//向数据库中添加数据
	//同一个时间内，只能有一个人来进行写操作
	err = db.Update(func(tx *bolt.Tx) error {
		//1.有桶
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil{
			coinbase, err:= transaction.NewCoinbase(address)
			genesis := GenesisBlock(*coinbase)
			bk, err := tx.CreateBucket([]byte(BUCKET_BLOCK))
			if err !=nil{
				return err
			}
			serialize, err := genesis.Serialize()
			if err !=nil{
				return err
			}
			bk.Put(genesis.Hash,serialize)
			//创建桶2
			bk2, err := tx.CreateBucket([]byte(BUCKET_STATUS))
			bk2.Put([]byte(LAST_HASH),genesis.Hash)
			lastHash = genesis.Hash
		}else{
			bk2 := tx.Bucket([]byte(BUCKET_STATUS))
			lastHash = bk2.Get([]byte(LAST_HASH ))
		}
		return nil
	})
	bc:=BlockChain{
		DB: db,
		LastHash: lastHash,
	}

	return &bc,err
}
//把区块添加到区块链中
func (bc *BlockChain)AddBlock(data []byte)error{

	new:=NewBlock(bc.LastHash,data)
	err :=bc.DB.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(BUCKET_BLOCK))
		if bk == nil{
			return errors.New("没有创建桶")
		}
		serialize, err := new.Serialize()
		if err != nil {
			return err
		}
		bk.Put(new.Hash,serialize)
		bk2 := tx.Bucket([]byte(BUCKET_STATUS))
		if bk2 == nil{
			return errors.New("没有创建桶2")
		}
		bk2.Put([]byte(LAST_HASH),new.Hash)
		bc.LastHash = new.Hash
		return nil
	})
	return err
}

//创建一个迭代器对象,迭代器只能在有区块链的情况下才可以使用迭代器
func (bc *BlockChain)Iterator()*ChainIterator{
	iterator:=ChainIterator{
		DB: bc.DB  ,
		currentHash: bc.LastHash,
	}

	return &iterator
}

//获取所有区块
func (bc *BlockChain)GetAllBlock()[]*Block{
	iterator := bc.Iterator()
	 bk := []*Block{}
	for {
		if iterator.HashNext() {
			bl, err := iterator.Next()
			if err != nil {
				fmt.Println("失败了")
				return nil;
			}
			bk = append(bk, bl)
		}else {
			break
		}

	}
	return bk
}