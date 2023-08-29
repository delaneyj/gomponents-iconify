package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotationAltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.6 22.8L3.75 13H6.6l7 7l5-5H16v-2h6v6h-2v-2.6l-6.4 6.4ZM2 11V5h2v2.6l6.4-6.4l9.85 9.8H17.4l-7-7l-5 5H8v2H2Z"/>`),
		g.Group(children),
	)
}