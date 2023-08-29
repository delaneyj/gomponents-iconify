package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFromTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4h12v2H6zm6 16.414l6.707-6.707l-1.414-1.414L13 16.586V8h-2v8.586l-4.293-4.293l-1.414 1.414z"/>`),
		g.Group(children),
	)
}