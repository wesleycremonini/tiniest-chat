The tiniest chat I could come with in golang.<br>
Its using a tcp network to receive and forward messages to all connected users.

How to use:

1. Install Go<br>
2. `git clone https://github.com/wesleycremonini/tiniest-chat.git`<br>
3. `cd /path/to/repo`<br>
4. `go mod tidy`<br>
5. `go run .`<br>

At this point, the tcp network is running and listening on port 5000

6. Open multiple terminal windows
7. `telnet 127.0.0.1 5000` on each of them
8. type any message in any terminal and press enter

All connected terminals will receive the message from the others

Example:

![Example Image](https://github.com/wesleycremonini/tiniest-chat/blob/main/example.jpeg)

Here, 4 terminals are connected, when you connect, you receive a random integer as a nickname, so you can differentiate multiple users