package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.364 13.121c.924.924 1.12 2.3.586 3.415l1.535 1.535a1 1 0 0 1-1.414 1.414l-1.535-1.535a3.001 3.001 0 0 1-3.415-4.829a3 3 0 0 1 4.243 0ZM12.95 15.95a1 1 0 1 0-1.414-1.414a1 1 0 0 0 1.414 1.414Z" clip-rule="evenodd"/><path d="M8 5h8v2H8V5Zm8 4H8v2h8V9Z"/><path fill-rule="evenodd" d="M4 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V4Zm3-1h10a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}