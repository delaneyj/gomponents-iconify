package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevationRise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-3.71l7.78-4.49l3.77 2.2L21 11.25V21H3M21 8.94l-6.45 3.73l-3.77-2.17L3 15v-2.21l7.78-4.49l3.77 2.2L21 6.75v2.19Z"/>`),
		g.Group(children),
	)
}