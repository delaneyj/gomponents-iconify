package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instapaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M10 5a1 1 0 0 0-1-1H8V2h8v2h-1a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h1v2H8v-2h1a1 1 0 0 0 1-1V5z" fill="currentColor"/>`),
		g.Group(children),
	)
}