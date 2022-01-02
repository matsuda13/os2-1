package main

import (
	"fmt"
	"os2-1/FileWrite"
	"os"
	"strconv"
	"math"
)

func main() {
	var list_buffersize [10] int
	var str string
	out_file := "./out/graphdata" + os.Args[2]
	for i:=0;i<10;i++{
		list_buffersize[i] = int(math.Pow(2,float64(i+2)))
	}
	fmt.Println(os.Args[1], os.Args[2])
	f, _ := os.Create(out_file)
	for i:=0;i<len(list_buffersize);i++{
		buffersize := list_buffersize[i]
		filesize,_ := strconv.Atoi(os.Args[2]) 
	    //num, _ := strconv.Atoi(os.Args[2])
	    average := filewrite.FWrite(buffersize, os.Args[1], filesize)
	    fmt.Printf("buffersize:%v の処理速度の平均秒数(μs)は、%vμs\n",buffersize,average)
		str = strconv.Itoa(buffersize) + " " + strconv.Itoa(average) + "\n"
		buf := []byte(str)
		f.Write(buf)
	}
}
