package main

type Plant struct {
  Name string `json:"name,omitempty"`
  DisplayName string `json:"displayName,omitempty"`
  CurrentPolls Polls `json:"currentPolls,omitempty"`
}

type Plants []Plant
