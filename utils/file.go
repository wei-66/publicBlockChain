package utils

import "os"

/**
 * @author: SuZhiXiaoWei
 * @DateTime: 2022/3/28 11:25
 **/

//判断文件是否存在
func FileExist(path string)bool{
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

