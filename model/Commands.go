package model

import (

)

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
  Data string `json:"data"`
}
