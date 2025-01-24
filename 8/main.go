// go routine example

package main
import (
	"fmt"
	"time"
	"sync"
)
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go dbCall(dbData[i])
	}
	wg.Wait()
	fmt.Println("Time taken:", time.Since(t0))
	fmt.Println("Results:", results)
}

func dbCall(i string) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Data from db: ", i)
	save(i)
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("The current results are: ", results)
	m.RUnlock()
}

