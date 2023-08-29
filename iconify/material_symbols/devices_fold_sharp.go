package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesFoldSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21V2.975l7-3.05V3h5v18H10Zm4.675-2H20V5h-3v13.025L14.675 19ZM2 5V3h2v2H2Zm0 16v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4V7h2v2H2Zm4-4V3h2v2H6Zm0 16v-2h2v2H6Z"/>`),
		g.Group(children),
	)
}