package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardThirtyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12C2 6.477 6.477 2 12 2a9.985 9.985 0 0 1 8 4V3.5h2v6h-4.834c.212.368.334.795.334 1.25v2.5a2.5 2.5 0 0 1-5 0v-2.5A2.5 2.5 0 0 1 16 8.458V7.5h2.616A8 8 0 1 0 20 12h2c0 5.523-4.477 10-10 10S2 17.523 2 12Zm13-2.25a1 1 0 0 0-1 1v2.5a1 1 0 1 0 2 0v-2.5a1 1 0 0 0-1-1Zm-5.625 3a.625.625 0 1 1 0 1.25H6.75v1.5h2.625a2.125 2.125 0 0 0 1.62-3.5a2.125 2.125 0 0 0-1.62-3.5H6.75V10h2.625a.625.625 0 1 1 0 1.25H7.5v1.5h1.875Z"/>`),
		g.Group(children),
	)
}