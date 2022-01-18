package filewrite

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"log"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

//write "0" to file
func FWrite(buffersize int, filepath string, filesize int) int {
	var t int
	n_loop := 20
	total := 0
	//loop 20 times and calculate average time
	for i:=0;i<n_loop;i++{
	    start := time.Now()
    	f, err := os.Create(filepath)
	    if err!=nil{
    		log.Fatal(err)
	    }
    	defer f.Close()
	    buf := bufio.NewWriterSize(f, buffersize)
		//no buffering
		if buffersize == 0{
	    	for i:=0;i<filesize;i++{
		    	f.Write([]byte("0"))
    		}
			t = int(time.Since(start).Microseconds())
	    	//fmt.Printf("progress:%vμs\n", t)
		//buffering
    	} else {
			for i := 0; i <filesize; i++ {
		        buf.Write([]byte("0"))
        	}
	        buf.Flush()
			t = int(time.Since(start).Microseconds())
	        //fmt.Printf("progress:%vμs\n", t)
		}
		total += t;
	}
	ave := total/n_loop
	return ave
}