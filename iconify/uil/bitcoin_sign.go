package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9a4 4 0 0 0-4-4V3a1 1 0 0 0-2 0v2h-2V3a1 1 0 0 0-2 0v2H6a1 1 0 0 0 0 2h1v10H6a1 1 0 0 0 0 2h3v2a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-2a4 4 0 0 0 2.62-7A4 4 0 0 0 19 9Zm-4 8H9v-4h6a2 2 0 0 1 0 4Zm0-6H9V7h6a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}