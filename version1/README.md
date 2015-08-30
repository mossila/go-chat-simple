
# โปรแกรม chat ง่ายๆ ด้วย golang #

## รอรับ connection ##
ใช้ package `net` เปิดรอรับ connecion ด้วยท่านี้  ถ้าเกิดว่ามี error ก็ให้ log ออกมาและ defer ไว้ว่าถ้าจบ function นี้ก็ให้ปิด socket ด้วย

```go

l, err := net.Listen("tcp", ":1234")
if err != nil {
    log.Fatal(err)
}
defer l.Close()
```

## client ต่อเข้ามา ##
เมื่อมี client ต่อเข้ามาจะเข้ามาผ่าน `l.Accept()` โดนเราจะได้ของออกมาเป็น connection และ error
```go
conn, err := l.Accept()
```

readline ที่ client ส่งเข้ามาด้วย bufio และ function ReadString
```go
msg, err := bufio.NewReader(conn).ReadString('\n')
```

จบ connection ด้วย
```
conn.Close()
```

## รองรับหลาย connecion ##
ครอบ `l.Accept` ด้วย `for{}`
```go
for {
    conn, _ := l.Accept()
    //readline
}
```



## Usage ##

```
go run main.go
```

## ทดสอบ connection ด้วย telnet ##

```
telnet localhost 1234
```

```bash
$ go run main.go
Hello I am client
```

```bash
$ telnet localhost 1234
$
Trying ::1...
Connected to localhost.
Escape character is '^]'.
Hello I am client.
Connection closed by foreign host.
$
$ telnet localhost 1234
Trying ::1...
Connected to localhost.
Escape character is '^]'.
Hello I am client 2
Connection closed by foreign host.
```
