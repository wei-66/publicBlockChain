package utils

import "crypto/sha256"

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/2/21 9:25
 **/
//计算哈希值
func GetHash(data []byte)[]byte{
	 hash := sha256.New()
	 hash.Write(data)
	return hash.Sum(nil)
}


