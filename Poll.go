package main

import (
  "time"
)

type Poll struct {
  Id string `json:"id"`
  Action string `json:"action"`
  DisplayName string `json:"displayName"`
  Deadline time.Time `json:"deadline"`
  Yee []string `json:"yee"`
  OrNah []string `json:"orNah"`
  YeeCount uint64 `json:"yeeCount"`
  OrNahCount uint64 `json:"orNahCount"`
}

type Polls []Poll
