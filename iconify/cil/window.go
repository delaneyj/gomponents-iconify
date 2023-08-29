package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M424 440V48H88v392H48v32h416v-32ZM120 80h120v120H120Zm0 152h120v144H120Zm272 208H120v-32h272Zm0-64H272V232h120Zm0-176H272V80h120Z"/>`),
		g.Group(children),
	)
}