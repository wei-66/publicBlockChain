package client

import (
	"btcionshow/block"
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
	case "addchain":
		cl.AddChain()
		break
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
func (cl *Cli) CreateChain(){
	createchain := flag.NewFlagSet("createchain", flag.ExitOnError)
	address := createchain.String("address","","账户名称")
	createchain.Parse(os.Args[2:])

	if utils.FileExist("./chain.db") {
		fmt.Println("文件已经存在")
		return
	}
	bc, err := block.NewChain(*address)
	defer bc.DB.Close()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("创建区块链成功")
}

func (cl *Cli) AddChain(){
	addchain := flag.NewFlagSet("addchain", flag.ExitOnError)
	//判断区块链是否存在
	if !utils.FileExist("./chain.db") {
		fmt.Println("区块链不存在")
		return
	}
	//获取区块链对象
	bc, err := block.NewChain(nil)
	defer  bc.DB.Close()
	if err != nil{
		fmt.Println(err.Error())
	}
	add := addchain.String("data", "", "新增区块的交易信息")
	addchain.Parse(os.Args[2:])
	err = bc.AddBlock([]byte(*add))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("添加区块成功")
}
func (cl *Cli) PrintChain(){
	//判断区块链是否存在
	if !utils.FileExist("./chain.db") {
		fmt.Println("区块链不存在")
		return
	}
	bc, err:= block.NewChain(nil)
	defer bc.DB.Close()
	if err != nil{
		fmt.Println(err.Error())
	}
	blocks := bc.GetAllBlock()
	for _,v := range blocks{
		fmt.Printf("hash:%x : 信息:%s\n",v.Hash,v.Data)
	}
 }
 func (c *Cli)SumChain(){
	 //判断区块链是否存在
	 if !utils.FileExist("./chain.db") {
		 fmt.Println("区块链不存在")
		 return
	 }
	 bc, err:= block.NewChain(nil)
	 defer bc.DB.Close()
	 if err != nil{
		 fmt.Println(err.Error())
	 }
	 blocks := bc.GetAllBlock()
	fmt.Println("区块数量为",len(blocks),"个")
 }
 func (cl *Cli)Help(){
 	fmt.Println("createchain: 创建带有创世区块的区块链--data参数1个")
 	fmt.Println("addchaion:在区块链中添加新的区块--data参数1个")
 	fmt.Println("printchain:打印区块链中区块的信息")
 	fmt.Println("sumchain:查找区块链中有几个区块")
 	fmt.Println("help:使用说明")
 }