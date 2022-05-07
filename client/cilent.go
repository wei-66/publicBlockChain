package client

import (
	"btcionshow/block"
	"btcionshow/transaction"
	"btcionshow/utils"
	"flag"
	"fmt"
	"os"
)

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/3/28 10:53
 **/
/**
 * 用户的交互入口
 *只负责读取用户的命令和参数，并进行解析
 *解析后调用对应的功能
 */
type Cli struct {
}

/*确定系统有哪些功能，需要哪些参数
1.创建带有创世区块的区块链 ， 参数 1个 创世区块d的交易信息 string
2.添加新的区块d到区块链中，参数 1个  新区块的交易信息 string
3.打印所有区块信息 参数无
5.获取当前区块链中区块的个数
*/
func (cl *Cli) Run() {
	switch os.Args[1] {
	case "createchain":
		cl.CreateChain()
		break
	//case "addchain":
	//cl.AddChain()
	//break
	case "printchain":
		cl.PrintChain()
	case "sumchain":
		cl.SumChain()
	case "help":
		cl.Help()
		break
	default:
		fmt.Println("没有对应的指令，输入--help指令查看详细命令")
		os.Exit(1)
	}

}
func (cl *Cli) CreateChain() {
	createchain := flag.NewFlagSet("createchain", flag.ExitOnError)
	address := createchain.String("address", "", "账户名称")
	createchain.Parse(os.Args[2:])

	if utils.FileExist("./chain.db") {
		fmt.Println("文件已经存在")
		return
	}
	bc, err := block.NewChain(*address)
	defer bc.DB.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("创建区块链成功")
}

//func (cl *Cli) AddChain(){
//	addchain := flag.NewFlagSet("addchain", flag.ExitOnError)
//	//判断区块链是否存在
//	if !utils.FileExist("./chain.db") {
//		fmt.Println("区块链不存在")
//		return
//	}
//	//获取区块链对象
//	bc, err := block.NewChain(nil)
//	defer  bc.DB.Close()
//	if err != nil{
//		fmt.Println(err.Error())
//	}
//	add := addchain.String("data", "", "新增区块的交易信息")
//	addchain.Parse(os.Args[2:])
//	err = bc.AddBlock([]byte(*add))
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("添加区块成功")
//}

//发起一笔交易，在添加到区块链中
func (cl *Cli) send() {
	send := flag.NewFlagSet("send", flag.ExitOnError)
	from := send.String("from", "", "交易发起者的地址")
	to := send.String("to", "", "交易接收者的地址")
	amount := send.Uint("amount", 0, "交易的数量")
	err := send.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("解析失败", err.Error())
	}
	//1.创建一个普通得交易
	tx, _ := transaction.NewTransaction(*from, *to, *amount)
	//2.把交易放到区块中，然后再把区块存储到区块链中
	//在这个过程中，产生新区快得过程中，会产生一个coinbase交易
	bc, err := block.NewChain("")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	//创建一个coinbase交易
	cb, err := transaction.NewCoinbase(*from)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = bc.AddBlock([]transaction.Transaction{*tx, *cb})
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("添加成功")
}

func (cl *Cli) PrintChain() {
	//判断区块链是否存在
	if !utils.FileExist("./chain.db") {
		fmt.Println("区块链不存在")
		return
	}
	bc, err := block.NewChain("")
	defer bc.DB.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	blocks := bc.GetAllBlock()
	//获取每一个区块
	for _, v := range blocks {
		fmt.Printf("当前区块hash:%x : 交易个数:%d\n", v.Hash, len(v.Txs))
		//遍历交易集合
		for _, tx := range v.Txs {
			fmt.Printf("\t交易hash:%x\n", tx.TxId)
		}
	}
}
func (c *Cli) SumChain() {
	//判断区块链是否存在
	if !utils.FileExist("./chain.db") {
		fmt.Println("区块链不存在")
		return
	}
	bc, err := block.NewChain("")
	defer bc.DB.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	blocks := bc.GetAllBlock()
	fmt.Println("区块数量为", len(blocks), "个")
}
func (cl *Cli) Help() {
	fmt.Println("createchain: 创建带有创世区块的区块链--data参数1个")
	fmt.Println("addchaion:在区块链中添加新的区块--data参数1个")
	fmt.Println("printchain:打印区块链中区块的信息")
	fmt.Println("sumchain:查找区块链中有几个区块")
	fmt.Println("help:使用说明")
}
