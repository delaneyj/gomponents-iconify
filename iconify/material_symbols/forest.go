package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22v-3h4v3h-4Zm-6 0v-4H0l3.85-6H2L9 2l7 10h-1.85l3.875 6H11v4H7Zm12.25-4L16 13h1.925l-5.3-7.575L15 2l7 10h-1.85L24 18h-4.75Z"/>`),
		g.Group(children),
	)
}