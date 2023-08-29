package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartColumnIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 15h15.938v.969H1zM15.906.065L11.15.94l1.609 1.609l-4.78 4.057L3.9 3.274a.5.5 0 1 0-.658.753l4.745 3.914l5.479-4.685L15.15 4.94l.756-4.875zM12 7h2.951v6.878H12zM2 8h2.982v5.878H2zm5 2h2.951v3.892H7z"/>`),
		g.Group(children),
	)
}