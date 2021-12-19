package main

import (
    "context"
    "fmt"
    "github.com/shurcooL/trayhost"
    "golang.design/x/hotkey"
    "golang.design/x/mainthread"
    "io/ioutil"
    
    "os"
    "runtime"
    
    "igoagent/util"
)


func main() {
   // EnterLoop must be called on the OS's main thread
   runtime.LockOSThread()
   
   go func() {
       mainthread.Init(fn)
   }()
   // Enter the host system's event loop
   ff, _ := ioutil.ReadFile("igo.ico")
   var item []trayhost.MenuItem
   item = append(item, trayhost.MenuItem{Title: "open igo", Enabled: enabled, Handler: openIgo})
   item = append(item, trayhost.MenuItem{Title: "exit", Enabled: enabled, Handler: trayhost.Exit})
   trayhost.Initialize("igo agent", ff, item)
   trayhost.EnterLoop()
}

func enabled() bool {
   return true
}

func openIgo() {
   util.ExecOSCmd("start igo -i")
}

func exit() {
    os.Exit(0)
}

func fn() { // Use fn as the actual main function.
    fmt.Println("Enter")
    var (
        mods = []hotkey.Modifier{hotkey.Modifier(hotkey.ModCtrl)}
        //k    = hotkey.KeyS
        k    = hotkey.Key1
    )
    
    // Register a desired hotkey.
    hk, err := hotkey.Register(mods, hotkey.Key(k))
    if err != nil {
        panic("hotkey registration failed")
    }
    
    // Start listen hotkey event whenever you feel it is ready.
    triggered := hk.Listen(context.Background())
    for range triggered {
        println("hotkey ctrl+1 is triggered")
        util.ExecOSCmd("start igo -i")
    }
}