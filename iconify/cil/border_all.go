package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M56 472h400a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16H56a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16ZM272 72h168v168H272Zm0 200h168v168H272ZM72 72h168v168H72Zm0 200h168v168H72Z"/>`),
		g.Group(children),
	)
}