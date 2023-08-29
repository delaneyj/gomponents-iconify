package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoGame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0m6 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-6 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0m12 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 12h7m4 0h7M3 6h1m4 0h13M3 18h1m4 0h8m4 0h1M6 3v1m0 4v8m0 4v1m6-18v7m0 4v7m6-18v13m0 4v1"/>`),
		g.Group(children),
	)
}