package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricalServicesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 15v-2h3v2h-3Zm0 4v-2h3v2h-3Zm-6 1v-2h-2v-4h2v-2h5v8h-5Zm-9-3V9h7V6H4V4h8v7H5v4h4v2H3Z"/>`),
		g.Group(children),
	)
}