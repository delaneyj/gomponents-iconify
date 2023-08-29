package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotaryLastLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0 0v6M12.5 9.5L6 3m5 0H6v5"/>`),
		g.Group(children),
	)
}