package BLC
import (
    "time"
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

//1.创建新区块函数
func NewBlock(data string,height int64,prevBlockHash []byte) *Block{
	//创建区块
	block:=&Block{Height:height,PrevBlockHash:prevBlockHash,Data:[]byte(data),Timestamp:time.Now().Unix(),Hash:nil}
	
	return block
}