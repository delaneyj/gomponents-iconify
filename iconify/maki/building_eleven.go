package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M2 1v9h4V6h2v4h1V1H2zm3 8H3V6h2v3zm0-4H3V3h2v2zm3 0H6V3h2v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}