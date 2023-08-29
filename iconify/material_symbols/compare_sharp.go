package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 23v-2H3V3h7V1h2v22h-2Zm-5-5h5v-6l-5 6Zm9 3v-9l5 6V5h-5V3h7v18h-7Z"/>`),
		g.Group(children),
	)
}