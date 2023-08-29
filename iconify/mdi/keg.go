package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22v-2h1v-4H5v-2h1v-3H5V7h6V3h-1V2h4v1h-1v4h6v4h-1v3h1v2h-1v4h1v2H5M17 9a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}