package util

import (
    "gopkg.in/toast.v1"
    "log"
)

func FirstLaunch() {
    notification := toast.Notification{
        AppID: "Igo Agent",
        Title: "First start reminder",
        Message: "Igo agent start successful!",
        Audio: toast.Reminder,
        Icon: "D:\\igoagent\\ico\\igo.ico", // This file must exist (remove this line if it doesn't)
        Actions: []toast.Action{
            {"protocol", "Star Igo", "https://github.com/oneisx/igo"},
            {"protocol", "For Help", "https://github.com/oneisx/igoagent"},
        },
    }
    err := notification.Push()
    if err != nil {
        log.Fatalln(err)
    }
}