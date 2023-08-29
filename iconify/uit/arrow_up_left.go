package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.868 7.096h9.584a.548.548 0 0 0 0-1.096H6.548A.548.548 0 0 0 6 6.548v10.904a.548.548 0 1 0 1.096 0V7.867l9.97 9.97a.543.543 0 0 0 .773 0a.545.545 0 0 0-.001-.77l-9.97-9.971z"/>`),
		g.Group(children),
	)
}