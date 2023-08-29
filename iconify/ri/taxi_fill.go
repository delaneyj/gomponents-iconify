package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12v9a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1H5v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-9l2.48-5.788A2 2 0 0 1 6.32 5H9V3h6v2h2.681a2 2 0 0 1 1.838 1.212L22 12ZM4.176 12h15.648l-2.143-5H6.32l-2.143 5ZM6.5 17a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm11 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}