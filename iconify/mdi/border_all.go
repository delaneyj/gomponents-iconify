package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11h-6V5h6m0 14h-6v-6h6m-8-2H5V5h6m0 14H5v-6h6m-8 8h18V3H3v18Z"/>`),
		g.Group(children),
	)
}