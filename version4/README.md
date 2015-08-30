## แยกการทำงานให้เป็นสัดส่วน ##
แยกส่วนของการทำงานที่รับ client เข้ามาเป็นอีกส่วน

```go
func clientHandler(c net.Conn) {
	defer c.Close()
	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Print(msg)
	}
}
```
ที่ `main` ก็จะเหลือเพียงแค่นี้

```go
func main() {
    l, err := net.Listen("tcp", ":1234")
    ...
    for {
        ...
        go clientHandler(conn)
    }
}
```

## ย้ายส่วนของการ start server ออกจาก `main` ##
และประกาศให้ ให้ ChatServer เป็น exported function ด้วย(ที่อื่นเรียกใช้งานได้)

```go
/*ChatServer start chatserver on tcp with specific port
 * Because this is exported function
 * if it has no comment here you will get warning like this
 * > exported function ChatServer should have comment or be unexported
 */
func ChatServer(port string) {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go clientHandler(conn)
	}
}
```

main ก็จะเหลือเพียงเท่านี้
```go
func main() {
    ChatServer(":1234")
}
```

และถ้าเรารับ port มาจาก command line args ด้วยก็จะไม่มีอะไรที่เรา hardcode เข้าไปแล้ว
```go
import "os"
func main() {
    port := ":" + os.Args[1]
    ChatServer(port)
}
