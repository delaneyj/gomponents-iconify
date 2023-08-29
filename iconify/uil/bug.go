package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14h2a1 1 0 0 0 0-2h-2v-1a5.15 5.15 0 0 0-.21-1.36A5 5 0 0 0 22 5a1 1 0 0 0-2 0a3 3 0 0 1-2.14 2.87A5 5 0 0 0 16 6.4a2.58 2.58 0 0 0 0-.4a4 4 0 0 0-8 0a2.58 2.58 0 0 0 0 .4a5 5 0 0 0-1.9 1.47A3 3 0 0 1 4 5a1 1 0 0 0-2 0a5 5 0 0 0 3.21 4.64A5.15 5.15 0 0 0 5 11v1H3a1 1 0 0 0 0 2h2v1a7 7 0 0 0 .14 1.38A5 5 0 0 0 2 21a1 1 0 0 0 2 0a3 3 0 0 1 1.81-2.74a7 7 0 0 0 12.38 0A3 3 0 0 1 20 21a1 1 0 0 0 2 0a5 5 0 0 0-3.14-4.62A7 7 0 0 0 19 15Zm-8 5.9A5 5 0 0 1 7 15v-4a3 3 0 0 1 3-3h1ZM10 6a2 2 0 0 1 4 0Zm7 9a5 5 0 0 1-4 4.9V8h1a3 3 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}