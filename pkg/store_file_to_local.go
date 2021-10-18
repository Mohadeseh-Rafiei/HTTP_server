package pkg

import (
	"fmt"
	"os"
	"sync"
)

func StoreByChunkToLocal(filename string, accessKey string, defaultChunkSize int, content []byte){
	//fmt.Println(content)
	os.Mkdir("./data/" + accessKey, 0777)
	filePath := "./data" + "/" +accessKey + "/" + filename

	outFile, _ := os.Create(filePath)
	goRoutines := len(content) / defaultChunkSize
	//fmt.Println(len(content))
	wg := sync.WaitGroup{}
	wg.Add(goRoutines + 1)
	for i:=0; i< goRoutines+ 1; i++{
		//fmt.Println(x)
		go func(i int) {
			var x []byte
			if i == goRoutines {
				x = content[i *defaultChunkSize: ]
			}
			if i < goRoutines {
				x = content[i *defaultChunkSize : (i+1) *defaultChunkSize]
			}
			outFile.WriteAt(x, int64(i*defaultChunkSize))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("file downloaded successfully!")
	return

}
