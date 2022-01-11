package filewrite

import (
	//"fmt"
	"time"
	"os"
	"bufio"
	"log"
)


func FWrite(buffersize int, filepath string, filesize int) int {
	var t int
	n_loop := 20
	total := 0
	for i:=0;i<n_loop;i++{
	    start := time.Now()
    	f, err := os.Create(filepath)
	    if err!=nil{
    		log.Fatal(err)
	    }
    	defer f.Close()
	    buf := bufio.NewWriterSize(f, buffersize)
    	if buffersize == 0{//no buffering
	    	for i:=0;i<filesize;i++{
		    	f.Write([]byte("0"))
    		}
			t = int(time.Since(start).Microseconds())
	    	//fmt.Printf("経過:%vμs\n", t)
    	} else {
			for i := 0; i <filesize; i++ {
		        buf.Write([]byte("0"))
        	}
	        buf.Flush()
			t = int(time.Since(start).Microseconds())
	        //fmt.Printf("経過:%vμs\n", t)
		}
		total += t;
	}
	ave := total/n_loop
	return ave
}