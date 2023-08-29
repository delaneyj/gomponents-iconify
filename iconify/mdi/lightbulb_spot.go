package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbSpot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19h2v3H9v-3m4 3h2v-3h-2v3M2 2v2h20V2H2m7 12v3h6v-3c2.5-1.43 5-3 5-8H4c0 5 2.5 6.57 5 8Z"/>`),
		g.Group(children),
	)
}