package main

import (
	"directoryReader"
	"fmt"
	"gui"
)


func main() {
    app, mainWindow := gui.CreateBaseApp()
    drives := directoryReader.GetDrives()
    gui.ShowDrives(drives, mainWindow)

    fmt.Println(directoryReader.TestDiskSpace())

    app.Exec()
}
