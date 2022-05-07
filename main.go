package main

import (
	"btcionshow/client"
)

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/2/21 9:08
 **/
func main() {
	cl := client.Cli{}
	cl.Run()

	/*	return
		bc, err := block.NewChain([]byte("创世区块"))
		defer bc.DB.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = bc.AddBlock([]byte("2154"))
		if err != nil {
			fmt.Println("添加失败")
		} else {
			fmt.Println("添加成功")
		}
		iterator := bc.Iterator()
		for {
			if iterator.HashNext() {
				bk, err := iterator.Next()
				if err != nil {
					fmt.Println("失败了")
					return
				}
				fmt.Println(string(bk.Data))
			} else {
				break
			}
		}
	*/
	//err = bc.AddBlockChain([]byte("2154"))
	//if err != nil{
	//	fmt.Println("添加失败")
	//}else{
	//	fmt.Println("添加成功")
	//}

	//fmt.Println("第一个区块的值为：",string(chain.Blocks[0].Data))
	//fmt.Println("第一个区块的随机数为：",chain.Blocks[0].Nonce)
	//block := NewBlock([]byte("hello"), nil)
	//fmt.Println(block.)
	//newBlock := NewBlock([]byte("sdafadsf"), block.Data)
	//fmt.Println(block.Honce)
	//chain.AddBlockChain(block.Data)
	//chain.AddBlockChain(newBlock.Data)
	//fmt.Println("第个区块名称："+string(chain.Blocks[0].Data))
	//for i := 0 ; i <len(chain.Blocks) ; i++{
	//	fmt.Println("第",i+1,"个区块名称："+string(chain.Blocks[i].Data))
	//}
}
