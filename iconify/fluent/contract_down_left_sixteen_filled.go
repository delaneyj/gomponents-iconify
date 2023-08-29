package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractDownLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3A1.5 1.5 0 0 0 3 4.5V8h3.23C7.209 8 8 8.792 8 9.77V13h3.5a1.5 1.5 0 0 0 1.5-1.5V9.27a.5.5 0 0 1 1 0v2.23a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7A2.5 2.5 0 0 1 4.5 2h2.23a.5.5 0 0 1 0 1H4.5Zm9 3h-2.793l3.147-3.146a.5.5 0 0 0-.708-.707L10 5.293V2.5a.5.5 0 0 0-1 0v4a.498.498 0 0 0 .52.5h3.98a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}