package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCornerDownLeftOutline0"><g id="evaCornerDownLeftOutline1"><path id="evaCornerDownLeftOutline2" fill="currentColor" d="M20 6a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 1-.29.71A1 1 0 0 1 17 12H8.08l2.69-3.39a1 1 0 0 0-1.52-1.17l-4 5a1 1 0 0 0 0 1.25l4 5a1 1 0 0 0 .78.37a1 1 0 0 0 .62-.22a1 1 0 0 0 .15-1.41l-2.66-3.36h8.92a3 3 0 0 0 3-3Z"/></g></g>`),
		g.Group(children),
	)
}