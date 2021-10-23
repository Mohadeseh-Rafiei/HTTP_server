package pkg

import (
	"HTTP_server/internal"
	"fmt"
	"os"
	"sync"
)

type Worker interface {
	Put(func(buffer []byte, outFile *os.File, startPos int) interface{})
	Get() func(buffer []byte, outFile *os.File, startPos int) interface{}
}
type WorkerImpl struct {
	chans chan func(buffer []byte, outFile *os.File, startPos int) interface{}
}

func NewWorker(poolSize int) *WorkerImpl {
	return &WorkerImpl{
		chans: make(chan func(buffer []byte, outFile *os.File, startPos int) interface{}, poolSize),
	}
}

func (p *WorkerImpl) Put(f func(buffer []byte, outFile *os.File, startPos int) interface{}) {
	p.chans <- f
}

func (p *WorkerImpl) Get() func(buffer []byte, outFile *os.File, startPos int) interface{} {
	return <-p.chans
}

func SaveByChunk(goRoutinsCount int, defaultChunkSize int, content []byte, outFile *os.File) {
	worker := NewWorker(goRoutinsCount)
	for i := 0; i < goRoutinsCount; i++ {
		worker.Put(func(buffer []byte, outFile *os.File, startPos int) interface{} {
			outFile.WriteAt(buffer, int64(startPos))
			return struct{}{}
		})
	}
	wg := sync.WaitGroup{}
	i := 0
	for {
		if i*defaultChunkSize > len(content) {
			fmt.Println(i*defaultChunkSize, len(content), i)
			break
		}
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			var x []byte
			startPos := i * defaultChunkSize
			endPos := startPos + defaultChunkSize
			if startPos+defaultChunkSize > len(content) {
				endPos = len(content) - 1
			}
			x = content[startPos:endPos]
			object := worker.Get()
			object(x, outFile, startPos)
			worker.Put(object)
		}(i, &wg)
		i += 1
	}
	wg.Wait()
}

func StoreByChunkToLocal(filename string, accessKey string, defaultChunkSize int, content []byte) error {
	if os.Mkdir("./data/"+accessKey, 0777) != nil {
		return internal.UnsuccessfulDownload
	}
	filePath := "./data" + "/" + accessKey + "/" + filename

	outFile, err := os.Create(filePath)
	if err != nil {
		return internal.UnsuccessfulDownload
	}
	SaveByChunk(2, defaultChunkSize, content, outFile)
	fmt.Println("file downloaded successfully!")
	return nil

}
