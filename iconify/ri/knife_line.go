package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnifeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.343 1.408L22.374 19.44a1.5 1.5 0 1 1-2.121 2.122l-4.596-4.596L12.12 20.5l-7.778-7.778A8 8 0 0 1 4.17 1.587l.173-.179Zm.241 3.07l-.051.11a6.005 6.005 0 0 0 1.048 6.535l.176.185l6.364 6.364l2.828-2.829L4.584 4.478Z"/>`),
		g.Group(children),
	)
}