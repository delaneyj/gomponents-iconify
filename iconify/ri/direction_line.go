package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.515L3.515 12L12 20.485L20.485 12L12 3.515Zm.707-2.122l9.9 9.9a1 1 0 0 1 0 1.414l-9.9 9.9a1 1 0 0 1-1.414 0l-9.9-9.9a1 1 0 0 1 0-1.414l9.9-9.9a1 1 0 0 1 1.414 0ZM13 10V7.5l3.5 3.5l-3.5 3.5V12h-3v3H8v-4a1 1 0 0 1 1-1h4Z"/>`),
		g.Group(children),
	)
}