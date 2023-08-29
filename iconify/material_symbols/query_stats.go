package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueryStats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.625 18.175L1 17l5-8l3 3.5L13 6l2.725 4.075q-.575.025-1.088.138t-1.012.312l-.55-.825l-3.8 6.175L6.25 12.35l-3.625 5.825ZM21.575 23l-3.125-3.125q-.5.35-1.113.525t-1.262.175q-1.875 0-3.188-1.313t-1.312-3.187q0-1.875 1.313-3.188t3.187-1.312q1.875 0 3.188 1.313t1.312 3.187q0 .65-.175 1.263t-.525 1.137l3.125 3.1L21.575 23Zm-5.5-4.425q1.05 0 1.775-.725t.725-1.775q0-1.05-.725-1.775t-1.775-.725q-1.05 0-1.775.725t-.725 1.775q0 1.05.725 1.775t1.775.725Zm2.225-8q-.475-.2-.988-.325t-1.062-.15L21.375 2L23 3.175l-4.7 7.4Z"/>`),
		g.Group(children),
	)
}