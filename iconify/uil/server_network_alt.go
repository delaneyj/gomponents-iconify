package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerNetworkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 6h3a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2Zm8 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm5 9h-6.18A3 3 0 0 0 13 17.18V15h4a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2A3 3 0 0 0 20 6V4a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 4 10v2a3 3 0 0 0 3 3h4v2.18A3 3 0 0 0 9.18 19H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2ZM6 4a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1Zm1 9a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1Zm5 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm-1-11H8a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}