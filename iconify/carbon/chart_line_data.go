package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.06 17.88A4.25 4.25 0 0 0 11 18a4 4 0 0 0 2.23-.68l3.22 2.87a3.88 3.88 0 0 0-.2 3.17A4 4 0 1 0 22.62 19l2.54-5.09a3.78 3.78 0 0 0 2.91-.53a4 4 0 1 0-4.69-.38l-2.54 5.09A3.86 3.86 0 0 0 20 18a4 4 0 0 0-2.23.68l-3.22-2.87a3.88 3.88 0 0 0 .2-3.17a4 4 0 1 0-6.45 4.29L4 25V2H2v26a2 2 0 0 0 2 2h26v-2H4.67ZM26 8a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm-4 14a2 2 0 1 1-2-2a2 2 0 0 1 2 2ZM11 12a2 2 0 1 1-2 2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}