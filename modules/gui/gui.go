package gui


import (
    "os"

    "github.com/therecipe/qt/widgets"
)


func CreateBaseApp() (*widgets.QApplication, *widgets.QMainWindow) {
    app := widgets.NewQApplication(len(os.Args), os.Args)

    mainWindow := widgets.NewQMainWindow(nil, 0)
    mainWindow.SetWindowTitle("Roald")

    mainWindow.Show()

    return app, mainWindow
}


func ShowDrives(drives []rune, mainWindow *widgets.QMainWindow) {
    driveLayout := widgets.NewQGridLayout2()

    for i, drive :=  range drives {
        driveButton := widgets.NewQPushButton2(string(drive), nil)

        driveLayout.AddWidget2(driveButton, 0, i, 0)
    }

    driveWidget := widgets.NewQWidget(nil, 0)
    driveWidget.SetLayout(driveLayout)
    mainWindow.SetCentralWidget(driveWidget)
}
