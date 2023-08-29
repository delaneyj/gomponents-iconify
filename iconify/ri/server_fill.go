package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16a1 1 0 0 1 1 1v7H3V4a1 1 0 0 1 1-1ZM3 13h18v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-7Zm4 3v2h3v-2H7ZM7 6v2h3V6H7Z"/>`),
		g.Group(children),
	)
}