package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.536 8.464L6 18m0-4v4h4m5.586-9.586a2 2 0 1 0 2.828-2.828a2 2 0 0 0-2.828 2.828"/>`),
		g.Group(children),
	)
}