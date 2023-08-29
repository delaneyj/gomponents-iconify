package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardToInboxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 23l-1.4-1.4l1.575-1.6H15v-2h4.175l-1.6-1.6L19 15l4 4l-4 4ZM2 20V4h20v9.8q-.675-.4-1.438-.6T19 13q-2.5 0-4.25 1.75T13 19v1H2Zm10-7l8-5V6l-8 5l-8-5v2l8 5Z"/>`),
		g.Group(children),
	)
}