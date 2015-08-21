#Go performance monitoring and displaying them with StatsD
##Better option (use StatsD & Graphite Docker image):
-```docker run -d --name graphite --restart=always -p 80:80 -p 2003:2003 -p 8125:8125/udp hopsoft/graphite-statsd```

-go to your boot2docker IP or localhost to get to Graphite UI

-```go get github.com/bmhatfield/go-runtime-metrics```

-start the go code with the boot2docker IP or if it's localhost then dont use any flags 
```go run statMe.go -statsd=192.168.59.103:8125```

-go performance matrix will stream the performance data on the UDP port to StatsD

-grab a redbull and enjoy staring at the pretty performance charts :)

![goroutines and gc](https://raw.githubusercontent.com/gunjan5/go30days-challenge/master/day14-runtime-performance-monitoring/goroutines_gc.png)
![memory](https://raw.githubusercontent.com/gunjan5/go30days-challenge/master/day14-runtime-performance-monitoring/memory.png)
