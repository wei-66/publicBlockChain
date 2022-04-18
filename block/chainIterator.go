package block

import (
	"bytes"
	"github.com/boltdb/bolt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/3/14 9:39
 **/
type ChainIterator struct {
	DB *bolt.DB
	//标准位   表示当前迭代器所到位置
	currentHash []byte
}

//使用后迭代器找上一个区块
/**
  * @Description: 描述
  * @DateTime: 2022/4/11 11:57
  * @Param: param 参数
  * @return: 返回
  */
func (iterator *ChainIterator)Next()(*Block,error){
	var block *Block
	var err error
	 err = iterator.DB.View(func(tx *bolt.Tx) error {
		 bk := tx.Bucket([]byte(BUCKET_BLOCK))
		 if bk == nil{
		 	return errors.New("没有桶")
		 }
		 blockBytes := bk.Get(iterator.currentHash)
		 block, err = DeSerialize(blockBytes)
		 iterator.currentHash = block.PrevHash
		 return nil
	})
	 return block,err
}

//判断是否还有下一个区块
func (iterator *ChainIterator)HashNext()bool{
	compare := bytes.Compare(iterator.currentHash, nil)
	//迭代到创世区块
	return compare != 0
}

