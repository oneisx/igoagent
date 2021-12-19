package agent

import (
    "context"
    "github.com/shurcooL/trayhost"
    "golang.design/x/hotkey"
    "golang.design/x/mainthread"
    "igoagent/util"
    "io/ioutil"
)

const openIgoCommand = "start igo -i"

func Start() {
    util.FirstLaunch()
    
    registerHotKey()
    
    initTray()
}

func registerHotKey() {
    go func() {
        mainthread.Init(fn)
    }()
}

func initTray() {
    ico, _ := ioutil.ReadFile("ico/igo.ico")
    var item []trayhost.MenuItem
    item = append(item, trayhost.MenuItem{Title: "Open Igo", Enabled: enabled, Handler: OpenIgo})
    item = append(item, trayhost.MenuItem{Title: "Exit Agent", Enabled: enabled, Handler: trayhost.Exit})
    trayhost.Initialize("igo agent", ico, item)
    trayhost.EnterLoop()
}

func enabled() bool {
    return true
}

func OpenIgo() {
    util.ExecOSCmd(openIgoCommand)
}

func fn() {
    mods := []hotkey.Modifier{hotkey.Modifier(hotkey.ModCtrl)}
    k := hotkey.Key1
    
    hk, err := hotkey.Register(mods, hotkey.Key(k))
    if err != nil {
        panic("hotkey registration failed")
    }
    
    triggered := hk.Listen(context.Background())
    for range triggered {
        OpenIgo()
    }
}
