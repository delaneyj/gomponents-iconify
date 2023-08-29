package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Proxy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.013 9.395l4.57-4.57L21 6.243V2h-4.243l1.411 1.41l-4.834 4.835a3.938 3.938 0 0 0-5.191 2.75H5.72a2 2 0 1 0 .005 2H8.14a3.94 3.94 0 0 0 5.204 2.757l4.83 4.83L16.758 22H21v-4.243l-1.41 1.411l-4.571-4.57a3.967 3.967 0 0 0 .841-1.603L18 13v2l3-3l-3-3v2l-2.143-.005a3.968 3.968 0 0 0-.844-1.6Z"/>`),
		g.Group(children),
	)
}