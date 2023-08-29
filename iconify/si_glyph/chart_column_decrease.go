package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartColumnDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 16h15.938v.969H0zm14.906-7.058l-4.756-.875l1.609-1.609l-4.78-4.057L2.9 5.733a.5.5 0 1 1-.658-.753l4.745-3.914l5.479 4.685l1.684-1.684l.756 4.875z"/><path d="M6 6h2.951v8.878H6zM1 9h2.982v5.878H1zm10 2h2.951v3.892H11z"/></g>`),
		g.Group(children),
	)
}