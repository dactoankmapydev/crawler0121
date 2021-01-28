# ioc-news
https://medium.com/@Mardiniii/make-it-real-week-3-search-engine-indexer-crawling-the-internet-160769bb95bb

https://github.com/Mardiniii/go_search_engine_indexer

https://tutorialedge.net/golang/go-rabbitmq-tutorial/

{'CVE', 'URL', 'IPv6', 'FileHash-SHA1', 'email', 'FileHash-PEHASH', 'IPv4', 'YARA', 'FileHash-SHA256', 'hostname', 'Mutex', 'URI', 'BitcoinAddress', 'domain', 'FileHash-IMPHASH', 'FileHash-MD5', 'FilePath', 'JA3', 'SSLCertFingerprin
t'}

https://otx.alienvault.com/user/toannd_96/subscribing
https://otx.alienvault.com/user/AlienVault/pulses
https://otx.alienvault.com/user/SOTPOTUSALIENWAREVOLT/pulses


```
package main

import (
    "fmt"
    "time"
)

func convertTimestamp(strTime string) int64 {
	layout := "2006-01-02T15:04:05.000000"
	t,_ := time.Parse(layout, strTime)
	return t.Unix()
}

func main() {
   loc, _ := time.LoadLocation("Europe/London")
   time1 := convertTimestamp("2021-01-27T23:03:37.543000")
   time2 := convertTimestamp("2021-01-11T16:02:35.285000")
   fmt.Println("timestamp1", time1)
   fmt.Println("timestamp2", time2)
   
   convertTime1 := time.Unix(time1, 0).In(loc).Format(time.RFC3339)
   convertTime2 := time.Unix(time2, 0).In(loc).Format(time.RFC3339)
   fmt.Println("UTC1", convertTime1)
   fmt.Println("UTC2", convertTime2)
}
```

```
 package main

 import "fmt"

 func main() {
         arr := []int64{1611788617, 1610380955}

         max := arr[0] // assume first value is the smallest

         for _, value := range arr {
                 if value > max {
                         max = value // found another smaller value, replace previous value in max
                 }
         }

         fmt.Println("The biggest/largest value is : ", max)
 }
 ```
