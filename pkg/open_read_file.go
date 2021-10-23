package pkg

import (
	"HTTP_server/internal"
	"fmt"
	"os"
	"sync"
)

type Pool interface {
	Put(func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{})
	Get() func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{}
}
type PoolImpl struct {
	chans chan func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{}
}

func NewPool(poolSize int) *PoolImpl {
	return &PoolImpl{
		chans: make(chan func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{}, poolSize),
	}
}

func (p *PoolImpl) Put(f func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{}) {
	p.chans <- f
}

func (p *PoolImpl) Get() func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{} {
	return <-p.chans
}
func ReadFileByChunk(file *os.File, fileSize int64, defaultChunkSize int, goRoutinsCount int) []byte {
	pool := NewPool(goRoutinsCount)
	//number := 0
	for i := 0; i < goRoutinsCount; i++ {
		pool.Put(func(file os.File, fileReader []byte, startPos int, chunksize int, endPos int) interface{} {
			file.ReadAt(fileReader[startPos:endPos], int64(startPos))
			return struct{}{}
		})
	}
	var wg sync.WaitGroup
	n := int(fileSize) / defaultChunkSize
	fileReader := make([]byte, fileSize, fileSize)
	i := 0
	for {
		if i*defaultChunkSize > int(fileSize) {
			fmt.Println(i*defaultChunkSize, n)
			break
		}
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup, fileReader []byte) {
			defer wg.Done()
			object := pool.Get()
			startPos := defaultChunkSize * i
			endPos := startPos + defaultChunkSize
			if startPos+defaultChunkSize > int(fileSize) {
				endPos = int(fileSize)
			}
			object(*file, fileReader, startPos, defaultChunkSize, endPos)
			fmt.Println("read from ", startPos, " in go routin ", i)
			pool.Put(object)
		}(i, &wg, fileReader)
		i += 1
	}
	wg.Wait()
	return fileReader
}

func OpenReadFile(directory string, fileName string) ([]byte, error) {
	filePath := "./data/" + directory + "/" + fileName
	file, err := os.Open(filePath)
	file_stat, err := file.Stat()
	if err != nil {
		return nil, internal.UnsuccessfulUpload
	}
	n := file_stat.Size()
	fmt.Println(n)
	defer file.Close()
	defaultChunkSize := 1024
	fileReader := ReadFileByChunk(file, n, defaultChunkSize, 2)
	return fileReader, nil
}
