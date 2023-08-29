package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14h14v1H5v-1m16 3V8H3v9h18M1 5h22v14H1V5m4 5h7v2H5v-2Z"/>`),
		g.Group(children),
	)
}