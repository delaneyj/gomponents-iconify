package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.618 5.968l1.453-1.453l1.414 1.414l-1.453 1.453A9 9 0 1 1 12 4c2.125 0 4.078.736 5.618 1.968ZM12 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14ZM11 8h2v6h-2V8ZM8 1h8v2H8V1Z"/>`),
		g.Group(children),
	)
}