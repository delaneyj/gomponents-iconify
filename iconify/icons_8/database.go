package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 4v24h20V4H6zm2 2h16v5H8V6zm0 7h16v6H8v-6zm0 8h16v5H8v-5z"/>`),
		g.Group(children),
	)
}