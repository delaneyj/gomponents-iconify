package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15h-2v-2h2m1 7v2h-4v-2h1v-2h-1v-2h3v4m-.08-15a1.5 1.5 0 0 0-1.42-1h-11a1.5 1.5 0 0 0-1.42 1L3 11v8a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1h5a7 7 0 0 1 7-7a6.84 6.84 0 0 1 3 .68V11M6.5 15A1.5 1.5 0 1 1 8 13.5A1.5 1.5 0 0 1 6.5 15M5 10l1.5-4.5h11L19 10Z"/>`),
		g.Group(children),
	)
}