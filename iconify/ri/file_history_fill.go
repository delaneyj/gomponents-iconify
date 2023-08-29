package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHistoryFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 4.999v14.01a.993.993 0 0 1-.993.991H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16Zm-3 7h-2v6h5v-2h-3V9Z"/>`),
		g.Group(children),
	)
}