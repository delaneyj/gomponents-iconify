package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5q1.9 0 3.7.463t3.2 1.35Q20.3 7.7 21.15 9t.85 3H2q0-1.7.85-3T5.1 6.812q1.4-.887 3.2-1.35T12 5ZM9 19q-2.35 0-4.188-1.388T2.3 14h19.4q-.675 2.225-2.512 3.613T15 19H9Z"/>`),
		g.Group(children),
	)
}