package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	hypriot "github.com/hypriot/docker-machine-hypriot"
)

func main() {
	plugin.RegisterDriver(hypriot.NewDriver("", ""))
}
