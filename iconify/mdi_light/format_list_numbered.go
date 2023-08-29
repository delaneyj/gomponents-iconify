package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11v-1h3v.9L3.2 13H5v1H2v-.9L3.8 11H2m1-3V5H2V4h2v4H3m-1 9v-1h3v4H2v-1h2v-.5H3v-1h1V17H2M20 6v1H7V6h13m0 6v1H7v-1h13m0 6v1H7v-1h13Z"/>`),
		g.Group(children),
	)
}