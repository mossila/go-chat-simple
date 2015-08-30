## แยกการทำงาน ของแต่ละ connection ออกไป##
ด้วยคำสัั่ง `go` ที่แตกการทำงานของแต่ละ `conn` ของแต่ละ client ออกไป

```go
go func() { //แตกการทำงานออกไป
    for {
        msg, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println(err.Error())
            break
        }
        fmt.Print(msg)
    }
} ()// () คือการเรียกใช้งาน
```

และแยก connection ของแต่ละ client ให้ handle กันเองใน `func` ที่แตกออกไปด้วย

```go
go func(c net.Conn) { //ส่งเข้าไปในชื่อใหม่ว่า c
    for {
        //นำ c มาใช้แทน conn เดิม
        msg, err := bufio.NewReader(c).ReadString('\n')
        if err != nil {
            fmt.Println(err.Error())
            break
        }
        fmt.Print(msg)
    }
}(conn) //conn ที่ต่อเข้ามาล่าสุด
```

## จบการ ทำงานของ connection ให้ครบถ้วน ##


```go
go func(c net.Conn) {
    defer c.Close()
    ...
}(conn)
```
