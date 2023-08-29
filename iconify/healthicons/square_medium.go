package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M9 10a1 1 0 0 1 1-1h28a1 1 0 0 1 1 1v28a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1V10Z"/>`),
		g.Group(children),
	)
}