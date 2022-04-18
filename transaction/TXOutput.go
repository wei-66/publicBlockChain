package transaction

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/4/18 9:37
 **/
//交易输出的结构体
type Output struct {

   Value int //描述输出币的金额

   ScriptSig []byte //锁定脚本
}


