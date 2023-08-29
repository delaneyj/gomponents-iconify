package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 7a1 1 0 0 0-1 1v1.69l-4-2.31a2 2 0 0 0-3 1.73v5.78a2 2 0 0 0 1 1.73a2 2 0 0 0 2 0l4-2.31V16a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Zm-1 5l-5 2.89V9.11l5 2.88ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}