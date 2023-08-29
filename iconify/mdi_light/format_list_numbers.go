package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListNumbers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 10.998v-1h3v.9l-1.8 2.1H5v1H2v-.9l1.8-2.1H2Zm1-3v-3H2v-1h2v4H3Zm-1 9v-1h3v4H2v-1h2v-.5H3v-1h1v-.5H2ZM20 6v1H7V6h13Zm0 6v1H7v-1h13Zm0 6v1H7v-1h13Z"/>`),
		g.Group(children),
	)
}