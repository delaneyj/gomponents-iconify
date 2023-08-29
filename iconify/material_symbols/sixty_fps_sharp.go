package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixtyFpsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V5h8v3H5v2h6v9H2Zm3-3h3v-3H5v3Zm10 0h4V8h-4v8Zm-3 3V5h10v14H12Z"/>`),
		g.Group(children),
	)
}