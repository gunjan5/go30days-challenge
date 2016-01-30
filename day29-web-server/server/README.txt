Name: Gunjan Patel
Assignment 1: web server (written in Go)
Date: Jan 28






Test cases: 
================
HTTP/1.0, 200 OK
	telnet localhost 8080
	GET / HTTP/1.0

HTTP/1.0, 400 BAD REQUEST
	telnet localhost 8080
	GET bad HTTP/1.0

HTTP/1.0, 404 NOT FOUND
	telnet localhost 8080
	GET /nothing.html HTTP/1.0

HTTP/1.0, 403 FORBIDDEN
	telnet localhost 8080
	GET /nopermission.html HTTP/1.0	

================
HTTP/1.1, 200 OK
	telnet localhost 8080
	GET / HTTP/1.0

HTTP/1.1, 400 BAD REQUEST

	GET bad HTTP/1.0

HTTP/1.1, 404 NOT FOUND

	GET /nothing.html HTTP/1.0

HTTP/1.1, 403 FORBIDDEN

	GET /nopermission.html HTTP/1.0	


List of files:

├── Santa Clara University_files
│   ├── -livewhale-plugins-jquery-jquery.lw-widget.js
│   ├── 311_levis.jpg
│   ├── BryanStevenson760480-760x481.jpg
...  ... ...
│   ├── scu.css
│   ├── search-autocomplete.js
│   └── sinskywchild760480.jpg
├── nopermission.html			//no read permission
├── index.html
├── minion.gif
├── r2d2.jpg
├── server
└── server.go

1 directory, 42 files



Instructions for running:
-This assignment is written in Google's Go programming language (with professor's prior permission)
-