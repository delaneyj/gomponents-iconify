package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignFactorialFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.001 15.046a1 1 0 0 0 1.998 0l.477-10.501a1.478 1.478 0 1 0-2.952 0l.477 10.5ZM12 18a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5H12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}