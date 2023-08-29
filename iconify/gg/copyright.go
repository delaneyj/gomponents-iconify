package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m13.392 10.436l1.419-1.418a4 4 0 1 0 .176 5.798l-1.313-1.313h-.21a2 2 0 1 1-.073-3.067Z"/><path fill-rule="evenodd" d="M12 3a9 9 0 1 1 0 18a9 9 0 0 1 0-18Zm0 2a7 7 0 1 1 0 14a7 7 0 0 1 0-14Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}