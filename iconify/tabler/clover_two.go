package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloverTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 11L7.603 7.56a2.104 2.104 0 0 1 0-2.95a2.04 2.04 0 0 1 2.912 0L11 5l.485-.39a2.04 2.04 0 0 1 2.912 0a2.104 2.104 0 0 1 0 2.95L11 11zm0 0l-3.397 3.44a2.104 2.104 0 0 0 0 2.95a2.04 2.04 0 0 0 2.912 0L11 17l.485.39a2.04 2.04 0 0 0 2.912 0a2.104 2.104 0 0 0 0-2.95L11 11zm3.44-3.397a2.104 2.104 0 0 1 2.95 0a2.04 2.04 0 0 1 0 2.912L17 11l.39.485a2.04 2.04 0 0 1 0 2.912a2.104 2.104 0 0 1-2.95 0M7.56 7.603a2.104 2.104 0 0 0-2.95 0a2.04 2.04 0 0 0 0 2.912L5 11l-.39.485a2.04 2.04 0 0 0 0 2.912a2.104 2.104 0 0 0 2.95 0M15 15l6 6"/>`),
		g.Group(children),
	)
}