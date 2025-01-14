package main

import (
	"fmt"
	"log/slog"

	"github.com/alessioalex/grpc-course/module1/proto"
)

func main() {
	p := proto.Person{
		Name: "Chris",
	}

	slog.Info(fmt.Sprintf("The name's %q", p.GetName()))
}
