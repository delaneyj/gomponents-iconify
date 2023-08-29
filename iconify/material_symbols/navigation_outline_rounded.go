package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18l-6.45 2.75q-.325.125-.612.063t-.488-.263q-.2-.2-.263-.5t.063-.625L11.075 4.05q.125-.3.388-.45T12 3.45q.275 0 .537.15t.388.45l6.825 15.375q.125.325.062.625t-.262.5q-.2.2-.488.263t-.612-.063L12 18Zm-4.9-.1l4.9-2.1l4.9 2.1l-4.9-11l-4.9 11Zm4.9-2.1Z"/>`),
		g.Group(children),
	)
}