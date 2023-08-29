package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteRow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1a5 5 0 1 1-8 0H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h16Zm-7 10v2h6v-2h-6Zm6-8H5v4h14V7Z"/>`),
		g.Group(children),
	)
}