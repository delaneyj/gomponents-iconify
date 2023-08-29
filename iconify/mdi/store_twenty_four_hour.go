package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreTwentyFourHour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 12h-1v-2h-2V7h1v2h1V7h1m-5 3H9v1h2v1H8V9h2V8H8V7h3m8 0V4H5v3H2v13h8v-4h4v4h8V7h-3Z"/>`),
		g.Group(children),
	)
}