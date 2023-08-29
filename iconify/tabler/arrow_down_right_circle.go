package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.464 8.464L18 18m-4 0h4v-4M8.414 8.414a2 2 0 1 0-2.828-2.828a2 2 0 0 0 2.828 2.828"/>`),
		g.Group(children),
	)
}