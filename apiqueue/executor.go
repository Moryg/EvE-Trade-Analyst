package apiqueue

import (
)


func executor() {
  var r IRequest
  for {
    r = <- queue

    r.Execute()
  }
}