package BLC
import (
	"time"
	"fmt"
	"strconv"
	"crypto/sha256"
	"bytes"
)
type Block struct{
	//1.区块高度asdfasdf
	Height int64
	//2.上一个区块hash
	PrevBlockHash []byte
	//3.交易数据
	Data []byte
	//4.时间戳
	Timestamp int64
	//5.hash
	Hash []byte
}

func (block *Block)  SetHash() {

	//1.Height 转 []byte
	heightBytes:=IntToHex(block.Height)
	fmt.Println(heightBytes)

	//2.将时间戳转 []byte
	//先将int64转字符串
	//第二个参数的范围为2～36代表进制
	timeString:=strconv.FormatInt(block.Timestamp,2)
	fmt.Println(timeString)
	//将字符串转字节数组
	timeBytes:=[]byte(timeString)
	fmt.Println(timeBytes)

	//3拼接所有属性
	blockBytes:=bytes.Join([][]byte{heightBytes,block.PrevBlockHash,block.Data,timeBytes,block.Hash},[]byte{})

	//4生成hash
	hash:=sha256.Sum256(blockBytes)  //返回32字节的字节数组
	block.Hash=hash[:]

}

//1.创建新区块函数
func NewBlock(data string,height int64,prevBlockHash []byte) *Block{
	//创建区块
	block:=&Block{Height:height,PrevBlockHash:prevBlockHash,Data:[]byte(data),Timestamp:time.Now().Unix(),Hash:nil}
	block.SetHash()
	return block
}