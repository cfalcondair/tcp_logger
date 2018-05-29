# tcp_logger
A simple tcp logger that runs on port 2000 and just prints back what you send it.

# Setup
- Ensure that go is installed.

# How to run

 - Simple run the go code from within the repo, `go run tcp_logger.go`.
 This will listen on port 2000 of localhost.
 - In a seperate terminal on the same machine, make a connection to this listener. We could use the TCP swiss armyknife, netcat(nc), for example:
    `nc localhost 2000`
 - Type any data into the command line and it will send that back to you.
 - Feel the rush.
