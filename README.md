# chat-websocket-gin

<a href="https://www.buymeacoffee.com/tinkerbaj"><img src="https://img.buymeacoffee.com/button-api/?text=Buy me a coffee&emoji=&slug=tinkerbaj&button_colour=5F7FFF&font_colour=ffffff&font_family=Cookie&outline_colour=000000&coffee_colour=FFDD00" /></a>
  
Example of using gorilla websocket with gin (chat with rooms)

*   Clone repo
*   Enter in console "go mod tidy" to fetch all packages project need
*   Enter "go run main.go" to start server

Websocket route will be ws://localhost:8080/ws/ + any name

for example: ws://localhost:8080/ws/room1

You can use WebSocket King Client Chrome extension to test it
If you use room1 id for example like ws://localhost:8080/ws/room1 
you can use this JSON 
```
{
"type": "message",
"sender": "user1",
"recipient": "user2",
"content": "How are you",
"id": "room1"
}
```
run 2 clients on the same URL ws://localhost:8080/ws/room1  and you can use the same JSON just important is to send on the same id in our case room1
