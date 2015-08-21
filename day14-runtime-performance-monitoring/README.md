#go performance monitoring and displaying them with StatsD
-install node.js
-download and make StatsD
-Configure config file
-start the deamon
-start the gocode and go performance matrix will stream the performance data on the UDP port to StatsD

##Better option (use StatsD & Graphite Docker image):
-```docker run -d --name graphite --restart=always -p 80:80 -p 2003:2003 -p 8125:8125/udp hopsoft/graphite-statsd```

-go to your boot2docker IP or localhost to get to Graphite UI

-start the go code with the boot2docker IP or if it's localhost then dont use any flags 
```go run statMe.go -statsd=192.168.59.103:8125```

-grab a redbull and enjoy staring at the pretty performance charts :)

![goroutines and gc](https://raw.githubusercontent.com/gunjan5/go30days-challenge/master/day14-runtime-performance-monitoring/goroutines_gc.png)
![memory](https://raw.githubusercontent.com/gunjan5/go30days-challenge/master/day14-runtime-performance-monitoring/memory.png)
