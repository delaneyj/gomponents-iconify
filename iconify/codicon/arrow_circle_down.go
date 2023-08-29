package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="m5.369 8.08l2.14 2.14V4.468h1v5.68l2.066-2.067l.707.707l-2.957 2.956h-.707L4.662 8.788l.707-.707Z"/><path d="M14 8A6 6 0 1 0 2 8a6 6 0 0 0 12 0Zm-1 0A5 5 0 1 1 3 8a5 5 0 0 1 10 0Z"/></g>`),
		g.Group(children),
	)
}