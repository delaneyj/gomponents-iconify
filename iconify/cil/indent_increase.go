package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M72 63.998h424v32H72zm128 88h296v32H200zm0 88h296v32H200zm0 88h296v32H200zm-128 88h424v32H72zM16 144.909v222.182L171.527 256Zm32 62.182L116.473 256L48 304.909Z"/>`),
		g.Group(children),
	)
}