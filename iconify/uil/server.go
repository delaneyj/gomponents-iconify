package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-6 0H6a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Zm9 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-6 0H6a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Zm9-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 2 11v2a3 3 0 0 0 .78 2A3 3 0 0 0 2 17v2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2v-2a3 3 0 0 0-.78-2A3 3 0 0 0 22 7Zm-2 14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-5-2a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM9 5H6a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}