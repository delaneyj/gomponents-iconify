package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleSoda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11v9a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-9a2 2 0 0 1 .6-1.42C11.1 7.89 11 4 11 4h-1V2h4v2h-1s-.1 3.89 1.4 5.58A2 2 0 0 1 15 11Z"/>`),
		g.Group(children),
	)
}