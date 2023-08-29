package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h6v2H4v2h3v2H4v4H2V7m14 0h4a2 2 0 0 1 2 2v2c0 1.11-.89 2-2 2h-2v2h4v2h-6v-4c0-1.1.9-2 2-2h2V9h-4V7m-6 0h4v10h-2V9h-2V7Z"/>`),
		g.Group(children),
	)
}