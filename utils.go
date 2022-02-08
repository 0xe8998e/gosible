/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-02-07 20:54:17
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/gosible/utils.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

import (
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/0xe8998e/gosible/ssh"
)

func CSTLayoutString() string {
	var cst *time.Location
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
	ts := time.Now()
	return ts.In(cst).Format("2006-01-02 15:04:05")
}

// 合并map,数据合并到dst中
func mergeMap(src map[interface{}]interface{}, dst map[interface{}]interface{}) {
	for variable, value := range src {
		if _, present := dst[variable]; !present {
			dst[variable] = value
		}
	}
}

func InStringSlice(haystack []string, needle string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}

	return false
}

func GetProjectAbsPath(filename string) (projectAbsPath string) {

	programPath, _ := filepath.Abs(filename)

	projectAbsPath = path.Dir(programPath)

	return projectAbsPath

}

// 模拟耗时任务
func takeUpTimeTask(done <-chan interface{}, inStream <-chan interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case val, ok := <-inStream:
				if !ok {
					return
				}

				executer := Executer{
					SshConfig: &ssh.Config{
						Host: val.(ExecuterParam).Host.Ip,
						Port: val.(ExecuterParam).Host.Port,
						User: val.(ExecuterParam).Host.UserName,
					},
				}

				c, err := executer.Run(val.(ExecuterParam).Task)

				x := ExecuterResult{
					HostResult: c,
					Error:      err,
				}
				outStream <- x

				// time.Sleep( time.Second)

			}
		}
	}()
	return outStream
}

// 扇出处理耗时任务
func fanOut(done <-chan interface{}, chanStream chan interface{}) []<-chan interface{} {
	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		// 耗时任务的分流管道
		finders[i] = takeUpTimeTask(done, chanStream)
	}
	return finders
}

// 扇入汇流结果通道
func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	// 管道汇流处理
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}
	// 从所有的通道中取数据
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	// 等待所有数据汇总完毕
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}
