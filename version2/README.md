## เพิ่มการรองรับ หลายข้อความ ##
เพิ่มการรองรับหลายข้อความด้วยการครอบ `for{}` เข้าไป

```go
for {
    msg, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println(err.Error())
        break
    }
    fmt.Print(msg)
}
```

แต่ว่าเมื่อครอบเข้าไปแบบนี้กลับทำให้ connection นี้เกิดการ block คือรับ connection ได้แค่ 1 เท่านั้น

```go
for { //for1
    conn, _ := l.Accept()
    for { //การทำงานเข้ามาติดที่ for นี้ถ้าไม่ break ออกไปก็จะไม่สามารถออกไปทำงานที่ for1 ได้
        msg, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            break
        }
        fmt.Print(msg)
    }
}
```
