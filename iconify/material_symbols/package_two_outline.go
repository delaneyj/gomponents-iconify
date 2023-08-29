package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19.425v-6.85L5 9.1v6.85l6 3.475Zm2 0l6-3.475V9.1l-6 3.475v6.85Zm-2 2.3L4 17.7q-.475-.275-.738-.725t-.262-1v-7.95q0-.55.263-1T4 6.3l7-4.025Q11.475 2 12 2t1 .275L20 6.3q.475.275.738.725t.262 1v7.95q0 .55-.263 1T20 17.7l-7 4.025Q12.525 22 12 22t-1-.275Zm5-13.2l1.925-1.1L12 4l-1.95 1.125l5.95 3.4Zm-4 2.325l1.95-1.125L8.025 6.3l-1.95 1.125L12 10.85Z"/>`),
		g.Group(children),
	)
}