package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.8 21H3v-5.8q1.2 0 2.1-.762T6 12.5q0-1.175-.9-1.937T3 9.8V4h6q0-1.05.725-1.775T11.5 1.5q1.05 0 1.775.725T14 4h6v6q1.05 0 1.775.725T22.5 12.5q0 1.05-.725 1.775T20 15v6h-5.8q0-1.25-.787-2.125T11.5 18q-1.125 0-1.913.875T8.8 21Z"/>`),
		g.Group(children),
	)
}