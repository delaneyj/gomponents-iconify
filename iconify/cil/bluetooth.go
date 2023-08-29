package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M150.627 376L232 294.627V496h38.627l120-120l-120-120l120-120l-120-120H232v201.373L150.627 136h-45.254l120 120l-120 120ZM264 54.627L345.373 136L264 217.373Zm0 240L345.373 376L264 457.373Z"/>`),
		g.Group(children),
	)
}