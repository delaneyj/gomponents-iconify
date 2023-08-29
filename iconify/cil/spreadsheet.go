package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64v392h480V64Zm448 360H48V96h416Z"/><path fill="currentColor" d="M88 136h88v32H88zm0 72h88v32H88zm0 72h88v32H88zm0 72h88v32H88zm200-216h136v32H288zm0 72h136v32H288zm0 72h136v32H288zm0 72h136v32H288zm-72-216h32v32h-32zm0 72h32v32h-32zm0 72h32v32h-32zm0 72h32v32h-32z"/>`),
		g.Group(children),
	)
}