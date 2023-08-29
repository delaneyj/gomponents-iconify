package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.464 15.536L18 6m0 4V6h-4m-5.586 9.586a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828"/>`),
		g.Group(children),
	)
}