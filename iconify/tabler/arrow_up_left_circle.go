package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.536 15.536L6 6m4 0H6v4m9.586 5.586a2 2 0 1 0 2.828 2.828a2 2 0 0 0-2.828-2.828"/>`),
		g.Group(children),
	)
}