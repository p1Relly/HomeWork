package main

import (
	"fmt"
	"sync"
	"time"
)

// ExecutePipeline управляет конвейерной обработкой задач
func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, j := range jobs {
		wg.Add(1)
		out := make(chan interface{})
		go worker(wg, j, in, out)
		in = out
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup, j job, in, out chan interface{}) {
	defer wg.Done()
	defer close(out)
	j(in, out)
}

// SingleHash вычисляет хеш вида crc32(data)+"~"+crc32(md5(data))
func SingleHash(in, out chan interface{}) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for data := range in {
		wg.Add(1)
		go func(d interface{}) {
			defer wg.Done()
			value := strconv.Itoa(d.(int))

			// Вычисление md5 с защитой мьютексом
			mu.Lock()
			md5 := DataSignerMd5(value)
			mu.Unlock()

			// Параллельное вычисление двух crc32
			crcChan := make(chan string)
			crcMd5Chan := make(chan string)

			go func() { crcChan <- DataSignerCrc32(value) }()
			go func() { crcMd5Chan <- DataSignerCrc32(md5) }()

			crc := <-crcChan
			crcMd5 := <-crcMd5Chan

			out <- crc + "~" + crcMd5
		}(data)
	}
	wg.Wait()
}

// MultiHash вычисляет конкатенацию 6 crc32 хешей
func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for data := range in {
		wg.Add(1)
		go func(d interface{}) {
			defer wg.Done()
			value := d.(string)
			result := make([]string, 6)
			innerWg := &sync.WaitGroup{}

			for th := 0; th < 6; th++ {
				innerWg.Add(1)
				go func(i int) {
					defer innerWg.Done()
					result[i] = DataSignerCrc32(strconv.Itoa(i) + value
				}(th)
			}
			innerWg.Wait()

			out <- strings.Join(result, "")
		}(data)
	}
	wg.Wait()
}

// CombineResults объединяет и сортирует результаты
func CombineResults(in, out chan interface{}) {
	var results []string
	for data := range in {
		results = append(results, data.(string))
	}
	sort.Strings(results)
	out <- strings.Join(results, "_")
}

/*

0 SingleHash data 0
0 SingleHash md5(data) cfcd208495d565ef66e7dff9f98764da
0 SingleHash crc32(md5(data)) 502633748
0 SingleHash crc32(data) 4108050209
0 SingleHash result 4108050209~502633748
4108050209~502633748 MultiHash: crc32(th+step1)) 0 2956866606
4108050209~502633748 MultiHash: crc32(th+step1)) 1 803518384
4108050209~502633748 MultiHash: crc32(th+step1)) 2 1425683795
4108050209~502633748 MultiHash: crc32(th+step1)) 3 3407918797
4108050209~502633748 MultiHash: crc32(th+step1)) 4 2730963093
4108050209~502633748 MultiHash: crc32(th+step1)) 5 1025356555
4108050209~502633748 MultiHash result: 29568666068035183841425683795340791879727309630931025356555

func main() {
	data := "0"
	startTime := time.Now()
	fmt.Printf("0 SingleHash data %s\n", data)

	dataSignerMd5 := HWGorutina.DataSignerMd5(data)
	fmt.Printf("0 SingleHash md5(data) %s\n", dataSignerMd5)

	dataSignerCrc32 := HWGorutina.DataSignerCrc32(dataSignerMd5)
	fmt.Printf("0 SingleHash crc32(md5(data)) %s\n", dataSignerCrc32)

	dataSignerCrc32Data := HWGorutina.DataSignerCrc32(data)
	fmt.Printf("0 SingleHash crc32(data) %s\n", dataSignerCrc32Data)
	fmt.Printf("0 SingleHash result %s~%s\n", dataSignerCrc32Data, dataSignerCrc32)
	fmt.Printf("\n\nComplete. Tile Slapesed: %.2f second\n\n", time.Now().Sub(startTime).Seconds())
}

*/
