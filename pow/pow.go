package pow

import (
	"btcionshow/utils"
	"bytes"
	"math/big"
	"strconv"
)

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/2/28 9:39
 **/

const BITS = 2 //难度系数，前面有多少个0
/**
 * 区块的hash值 < 系统给定的hash值
 */

type ProofOfWork struct {
	//PrevHash []byte //上一个区块hash
	//TimeStamp int64  //时间戳
	//Data []byte
	Block  BlockeInterface
	Target *big.Int
}
type BlockeInterface interface {
	GetTimeStamp() int64
	GetPrevHash() []byte
	GetData() []byte
}

/**
实例化一个pow结构体，并且返回
*/

func NewPow(block BlockeInterface) *ProofOfWork {
	target := big.NewInt(1) //声明一个大整数类型的1
	//hash值256位 向左移BITS位
	target = target.Lsh(target, 255-BITS)
	pow := ProofOfWork{
		//PrevHash : prevHash, //上一个区块hash
		//TimeStamp : timeStamp,//时间戳
		//Data : data,
		Block: block,
		Target: target,
	}
	return &pow
}
//用来寻找随机数
func (pow *ProofOfWork)Run()([]byte,int64){
	var nonce int64 //随机数
	nonce = 0
	//block := pow.Block
	//block.Nonce = nonce
	timeBytes:=[]byte(strconv.FormatInt(pow.Block.GetTimeStamp(),10))
	num:=big.NewInt(0)

	//转型  []byte转成大整数
	for{
		nonceBytes:=[]byte(strconv.FormatInt(nonce,10))
		hashByets:=bytes.Join([][]byte{pow.Block.GetData(),pow.Block.GetPrevHash(),timeBytes,nonceBytes},[]byte{})
		hash:=utils.GetHash(hashByets)//当前区块的hash值
		//fmt.Println("正在寻找nonce,当前的nonce为",nonce)
		num = num.SetBytes(hash)//用来转换成大整数的
		if(num.Cmp(pow.Target)==-1 ){
			return hash,nonce
		}
		nonce++
	}
	return nil,0
}
