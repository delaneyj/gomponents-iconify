package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevationDecline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H3v-9.75L9.45 15l3.77-2.2L21 17.29V21M3 8.94V6.75l6.45 3.75l3.77-2.2L21 12.79V15l-7.78-4.5l-3.77 2.17L3 8.94Z"/>`),
		g.Group(children),
	)
}