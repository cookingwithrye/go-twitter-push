package main

import (
    "github.com/jdiez17/go-pushover"
    "time"
)

func Notify(text string) {
    p := pushover.Pushover{
        Config.PushoverUser,
        Config.PushoverAPIKey,
    }
    n := pushover.Notification{
        Title:     "twitter alert",
        Message:   text,
        Timestamp: time.Now(),
    }
    p.Notify(n)
}
   
