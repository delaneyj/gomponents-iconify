package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.452 6H6.547a.548.548 0 0 0 0 1.096h9.585l-9.97 9.97a.545.545 0 1 0 .772.772l9.97-9.971v9.586a.548.548 0 0 0 1.096 0V6.546A.548.548 0 0 0 17.452 6z"/>`),
		g.Group(children),
	)
}