package main

import (
    "bufio"
    "fmt"
    "os"
)

type Tasks map[int]string

func main() {
    id, tasks := 0, Tasks{}
    
    run(id, tasks)
}

func run(id int, tasks Tasks) {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var cmd string

    fmt.Println("Choose an option:\n c - create task\n q - quit")
    fmt.Scanln(&cmd)

    switch cmd {
    case "c":
        fmt.Print("What do you want to do? ")
        task, _ := in.ReadString('\n')
        tasks.Add(id, task)
        id++
        run(id, tasks)
    case "q":
        fmt.Println(tasks)
        return
    default:
        fmt.Println("Unknown command")
    }
}

func (t Tasks) Add(id int, task string) {
    t[id] = task
}
