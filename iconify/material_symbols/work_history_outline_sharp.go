package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkHistoryOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19V8v11v-.375V19Zm-2 2V6h6V2h8v4h6v6.275q-.45-.325-.963-.563T20 11.3V8H4v11h7.075q.075.525.225 1.025t.375.975H2Zm8-15h4V4h-4v2Zm8 17q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm.5-5.2V15h-1v3.2l2.15 2.15l.7-.7l-1.85-1.85Z"/>`),
		g.Group(children),
	)
}