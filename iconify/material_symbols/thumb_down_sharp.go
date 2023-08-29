package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 16v-4.4L4.65 3H16v13l-7 7l-1.85-1.85L8.45 16H1Zm17 0V3h4v13h-4Z"/>`),
		g.Group(children),
	)
}