
<h1>
  <p align="center" style="text-align: center">  
    GoCap
  </p>
</h1>


<p align="center" style="text-align: center">  
  <img src="https://user-images.githubusercontent.com/110062350/202498091-4c5c4852-fbe1-49ab-986f-4d1c5bbd65b7.png">
</p>

<p align="center">
   API Wrapper for Capmonster for Golang
</p>


## Contents

- [Introduction](#introduction)
- [Features](#features)
- [Package](#Package)
- [Docs](#Docs)

## Introduction

this is an api wrapper made in golang this can be used with capmonster
to use this in your project make sure its in golang and go to the docs 
down below

## Features

- fast & Lightweight
- Very Simple to use
- Can be used in any project
- has get Task/Solve funcs

## Package

import the package into your project:


```go
import "github.com/yaboipy/gocap"
```


## Docs

Setting The Payload:
```go
package main

import (
    "github.com/yaboipy/gocap"
    "log"
)

func main() {
    var (
      url = "discord.com" // site domain
      Key = "4c672d35-0701-42b2-88c3-78380b0db560" //site key
      Agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) " //useragent
      
      gocap := gocap.GoCap()
    )
    err := gocap.SetPayload(Key, Url, Agent)
    if err != nil {
        log.Fatal(err)
    }
}
```





Get Task ID Reponse:
```go
package main

import (
    "github.com/yaboipy/gocap"
    "log"
)

func main() {
     // make sure to set payload before running the functions!!!
    gocap := gocap.GoCap()
    apikey := "jkeargh98w4hq365tyq25h3h" // Capmonster API Key
    TaskID := gocap.GetTaskID(apikey)
    print("Task ID: ", task)
}
```




Solve Captcha using Task ID:
```go
package main

import (
    "github.com/yaboipy/gocap"
    "log"
)

func main() {
    var (
      url = "discord.com" // site domain
      Key = "4c672d35-0701-42b2-88c3-78380b0db560" //site key
      apikey = "jkeargh98w4hq365tyq25h3h" // Capmonster API Key
      Agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) " //useragent
      
      gocap := gocap.GoCap()
    )
    
    err := gocap.SetPayload(Key, Url, Agent)
    if err != nil {
        log.Fatal(err)
    }
    
    TaskID := gocap.GetTaskID(apikey)
    Solved := gocap.Solve(apikey)
    print("Solved Captcha String: ", Sovled)
}
```
