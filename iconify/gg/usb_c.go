package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8Z"/><path fill-rule="evenodd" d="M3 12a5 5 0 0 1 5-5h8a5 5 0 0 1 0 10H8a5 5 0 0 1-5-5Zm5-3h8a3 3 0 1 1 0 6H8a3 3 0 1 1 0-6Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}