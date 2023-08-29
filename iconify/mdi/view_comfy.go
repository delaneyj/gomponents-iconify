package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9h4V5H3v4m0 5h4v-4H3v4m5 0h4v-4H8v4m5 0h4v-4h-4v4M8 9h4V5H8v4m5-4v4h4V5h-4m5 9h4v-4h-4v4M3 19h4v-4H3v4m5 0h4v-4H8v4m5 0h4v-4h-4v4m5 0h4v-4h-4v4m0-14v4h4V5h-4Z"/>`),
		g.Group(children),
	)
}