package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Empathize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 5.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0M12 21.368l5.095-5.096a3.088 3.088 0 1 0-4.367-4.367l-.728.727l-.728-.727a3.088 3.088 0 1 0-4.367 4.367L12 21.368z"/>`),
		g.Group(children),
	)
}