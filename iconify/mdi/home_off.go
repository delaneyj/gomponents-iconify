package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L18.11 20H14v-4.11L12.11 14H10v6H5v-8H2l4.27-3.84L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M19 12h3L12 3L8.95 5.75L19 15.8V12Z"/>`),
		g.Group(children),
	)
}