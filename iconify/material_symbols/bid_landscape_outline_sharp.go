package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BidLandscapeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-4.05V19h14v-8.75L13.05 17L9 12.95l-4 4Zm0-2.85l4-4l3.95 3.95L19 7.25V5H5v9.1Zm0-3.85v-3v6.8v-3.95v6.85v-4V17v-6.75Zm0 3.85V5v9.05v-3.95v4Zm0 2.85v-4V17v-6.75V19v-2.05Z"/>`),
		g.Group(children),
	)
}