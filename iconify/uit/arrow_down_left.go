package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.452 16.904H7.867l9.97-9.97a.546.546 0 1 0-.77-.772l-9.971 9.97V6.548a.548.548 0 0 0-1.096 0v10.904c0 .303.245.548.548.548h10.904a.548.548 0 0 0 0-1.096z"/>`),
		g.Group(children),
	)
}