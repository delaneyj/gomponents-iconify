package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentWidthTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2h8a2 2 0 0 1 2 2v2.336a1.497 1.497 0 0 0-.98 1.662h-.52a1.5 1.5 0 0 0-.001 3h.522A1.498 1.498 0 0 0 16 12.664V16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3.336a1.497 1.497 0 0 0 .98-1.662h.52a1.5 1.5 0 0 0 .001-3H4.98A1.498 1.498 0 0 0 4 6.335V4a2 2 0 0 1 2-2ZM3.876 7.42a.5.5 0 0 1-.047.706L2.831 9H5.5a.5.5 0 1 1 0 1H2.83l1 .874a.5.5 0 0 1-.66.752l-2-1.75a.5.5 0 0 1 0-.752l2-1.75a.5.5 0 0 1 .706.047ZM17.17 10l-.998.874a.5.5 0 0 0 .658.752l2-1.75a.5.5 0 0 0 0-.752l-2-1.75a.5.5 0 0 0-.658.752L17.17 9l-2.669-.003a.5.5 0 0 0 0 1l2.67.003Z"/>`),
		g.Group(children),
	)
}