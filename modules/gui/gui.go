package gui


import (
    "os"

    "github.com/therecipe/qt/widgets"
)


func CreateBaseApp() *widgets.QMainWindow {
    app := widgets.NewQApplication(len(os.Args), os.Args)

    mainWindow := widgets.NewQMainWindow(nil, 0)
    mainWindow.SetWindowTitle("File Explorer")
    mainWindow.SetMinimumSize2(800, 600)

    title := widgets.NewQLabel2("File Explorer", mainWindow, 0)
    title.SetStyleSheet("font-size: 20px; font-weight: bold;")

    mainLayout := widgets.NewQVBoxLayout()
    mainLayout.AddWidget(title, 100, 0)

    mainWindow.SetLayout(mainLayout)
    mainWindow.Show()

    app.Exec()

    return mainWindow
}
