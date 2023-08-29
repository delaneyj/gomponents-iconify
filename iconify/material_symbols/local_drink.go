package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalDrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.975 22q-.775 0-1.337-.513T5 20.226L3 2h18l-2 18.225q-.075.75-.638 1.263T17.026 22H6.975ZM5.65 8h12.7l.4-4H5.25l.4 4ZM12 19q1.2 0 2.025-.825t.825-2.025q0-1.025-.663-2.225T12 11q-1.525 1.725-2.188 2.925T9.15 16.15q0 1.2.825 2.025T12 19Z"/>`),
		g.Group(children),
	)
}