package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThirtyFpsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-3h6v-2.5H2v-3h6V8H2V5h9v5.5L9.5 12l1.5 1.5V19H2Zm13-3h4V8h-4v8Zm-3 3V5h10v14H12Z"/>`),
		g.Group(children),
	)
}