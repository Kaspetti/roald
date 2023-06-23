package gui


import (
    "os"

    "github.com/therecipe/qt/widgets"
)


func CreateBaseApp() *widgets.QApplication {
    app := widgets.NewQApplication(len(os.Args), os.Args)

    window := widgets.NewQMainWindow(nil, 0)
    window.SetWindowTitle("File Explorer")
    window.SetMinimumSize2(800, 600)

    title := widgets.NewQLabel2("File Explorer", window, 0)
    title.SetStyleSheet("font-size: 20px; font-weight: bold;")

    mainLayout := widgets.NewQVBoxLayout()
    mainLayout.AddWidget(title, 100, 0)

    window.SetLayout(mainLayout)
    window.Show()

    return app
}
