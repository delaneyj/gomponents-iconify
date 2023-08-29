package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListChecks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18v1H7v-1h13m0-6v1H7v-1h13m0-6v1H7V6h13M2 5h3v3H2V5m1 1v1h1V6H3m-1 5h3v3H2v-3m1 1v1h1v-1H3m-1 5h3v3H2v-3m1 1v1h1v-1H3Z"/>`),
		g.Group(children),
	)
}