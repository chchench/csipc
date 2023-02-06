# Client-Server IPC

Demonstration of IPC between two processes using named pipe. In this demo, there are two programs. One is the client which generates a sequence of random integer numbers between 0 and 31 inclusive, which are sent to the server via a named pipe channel. The server retrieves the numbers from the channel and shows in a histogram the number of instances received for each integer.

![Histogram displayed in terminal](screenshots/histogram.png)