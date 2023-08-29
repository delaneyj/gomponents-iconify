package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InHomeModeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20V10H1l11-9l11 9h-3v1h-2V8.475l-6-4.9l-6 4.9V18h2v2Zm10.2 2L10 17.8l1.4-1.4l2.8 2.8l5.9-5.9l1.4 1.4Z"/>`),
		g.Group(children),
	)
}