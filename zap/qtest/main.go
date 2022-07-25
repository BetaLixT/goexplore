package main

import "go.uber.org/zap"

func main() {
  lgr, _ := zap.NewProduction()
  lgr.Info("test")
}
