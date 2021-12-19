package util

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

const (
    windows          = "windows"
    winOSCommand     = "cmd.exe"
    winCommandOption = "/c"
    
    linux              = "linux"
    linuxOSCommand     = "/bin/bash"
    linuxCommandOption = "-c"
    
    macos              = "darwin"
    macOSCommand       = "/usr/bin/open"
    macOSCommandOption = "-a"
    
)

type OperationSystem interface {
    ExecOSCmd(command string)
}

type Windows struct{}

type Linux struct{}

type MacOS struct{}

func (windows *Windows) ExecOSCmd(command string) {
    cmd := buildWindowsCmd(command)
    doExecOSCmd(cmd)
}

func (linux *Linux) ExecOSCmd(command string) {
    cmd := buildLinuxCmd(command)
    doExecOSCmd(cmd)
}

func (macos *MacOS) ExecOSCmd(command string) {
    cmd := buildMacOSCmd(command)
    doExecOSCmd(cmd)
}

func chooseOS() OperationSystem {
    switch runtime.GOOS {
    case windows:
        return new(Windows)
    case linux:
        return new(Linux)
    case macos:
        return new(MacOS)
    default:
        fmt.Println("Error: Operation system is not supported!")
        os.Exit(1)
    }
    return nil
}

func ExecOSCmd(command string) {
    operationSystem := chooseOS()
    operationSystem.ExecOSCmd(command)
}

func buildMacOSCmd(command string) *exec.Cmd {
    commands := []string{macOSCommandOption, command}
    return exec.Command(macOSCommand, commands...)
}

func buildLinuxCmd(command string) *exec.Cmd {
    commands := []string{linuxCommandOption, command}
    return exec.Command(linuxOSCommand, commands...)
}

func buildWindowsCmd(command string) *exec.Cmd {
    commands := []string{winCommandOption, command}
    return exec.Command(winOSCommand, commands...)
}

func doExecOSCmd(cmd *exec.Cmd) bool {
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        return false
    }
    return true
}
