package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitScreenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9V6h-3V4h5v5h-2ZM2 9V4h5v2H4v3H2Zm15 11v-2h3v-3h2v5h-5ZM2 20v-5h2v3h3v2H2Zm4-4V8h12v8H6Z"/>`),
		g.Group(children),
	)
}