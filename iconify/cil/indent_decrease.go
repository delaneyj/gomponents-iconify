package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M72 63.998h424v32H72zm128 88h296v32H200zm0 88h296v32H200zm0 88h296v32H200zm-128 88h424v32H72zm88-271.089L4.473 256L160 367.091Zm-32 160L59.527 256L128 207.091Z"/>`),
		g.Group(children),
	)
}