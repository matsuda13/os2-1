package main

import (
	"fmt"
	"os2-1/FileWrite"
	"os"
	"strconv"
	"math"
	"log"
)

func main() {
	var list_buffersize [10] int
	var str string
	var speed float64
	filesize,_ := strconv.Atoi(os.Args[2])
	out_file := "./out/graphdata" + os.Args[2] + ".dat"
	for i:=0;i<10;i++{
		list_buffersize[i] = int(math.Pow(2,float64(i+2)))
	}
	fmt.Println(os.Args[1], os.Args[2])
	f, err := os.Create(out_file)
	if err != nil{
		log.Fatal(err)
	}
	for i:=0;i<len(list_buffersize);i++{
		buffersize := list_buffersize[i] 
	    average := filewrite.FWrite(buffersize, os.Args[1], filesize)
	    fmt.Printf("buffersize:%v の処理速度の平均秒数(μs)は、%vμs\n",buffersize,average)
		//b/μs = Mb/s
		speed = float64(filesize) / float64(average)
		str = strconv.Itoa(buffersize) + " " + strconv.FormatFloat(speed, 'f', 2, 64) + "\n"
		buf := []byte(str)
		f.Write(buf)
	}
	//変更2
}
