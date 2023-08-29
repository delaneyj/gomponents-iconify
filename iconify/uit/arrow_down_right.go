package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.453 6a.548.548 0 0 0-.549.548v9.585l-9.97-9.971a.546.546 0 0 0-.772.771l9.97 9.971H6.549a.548.548 0 0 0 0 1.096h10.904a.548.548 0 0 0 .548-.548V6.548A.548.548 0 0 0 17.453 6z"/>`),
		g.Group(children),
	)
}