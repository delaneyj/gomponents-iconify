package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKitLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 1a1 1 0 0 1 1 1v3h4a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h4V2a1 1 0 0 1 1-1h8Zm4 6H4v12h16V7Zm-7 2v3h3v2h-3.001L13 17h-2l-.001-3H8v-2h3V9h2Zm2-6H9v2h6V3Z"/>`),
		g.Group(children),
	)
}