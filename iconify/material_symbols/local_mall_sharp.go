package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalMallSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V6h4q0-2.075 1.463-3.538T12 1q2.075 0 3.538 1.463T17 6h4v16H3Zm9-8q2.075 0 3.538-1.463T17 9h-2q0 1.25-.875 2.125T12 12q-1.25 0-2.125-.875T9 9H7q0 2.075 1.463 3.538T12 14ZM9 6h6q0-1.25-.875-2.125T12 3q-1.25 0-2.125.875T9 6Z"/>`),
		g.Group(children),
	)
}