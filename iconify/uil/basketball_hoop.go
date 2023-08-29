package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballHoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 10H6a1 1 0 0 0 0 2h1v9a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-9h1a1 1 0 0 0 0-2Zm-7 8H9v-2h2Zm0-4H9v-2h2Zm4 4h-2v-2h2Zm0-4h-2v-2h2Zm5-12H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3a1 1 0 0 0 0-2a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1a1 1 0 0 0 0 2a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm-6 6a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}