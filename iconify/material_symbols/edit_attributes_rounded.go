package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditAttributesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-2.075 0-3.538-1.463T2 12q0-2.075 1.463-3.538T7 7h10q2.075 0 3.538 1.463T22 12q0 2.075-1.463 3.538T17 17H7Zm1.75-3.4l2.325-2.325q.225-.225.225-.525t-.225-.525Q10.85 10 10.55 10t-.525.225L8.05 12.2l-.475-.475Q7.35 11.5 7.05 11.5t-.525.225q-.225.225-.225.525t.225.525l.825.825q.3.3.7.3t.7-.3Z"/>`),
		g.Group(children),
	)
}