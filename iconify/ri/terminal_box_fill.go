package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm9 12v2h6v-2h-6Zm-3.586-3l-2.828 2.828L7 16.243L11.243 12L7 7.757L5.586 9.172L8.414 12Z"/>`),
		g.Group(children),
	)
}