package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 6.78V3a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v3.78A3 3 0 0 0 6 9v6a3 3 0 0 0 1 2.22V21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3.78A3 3 0 0 0 18 15V9a3 3 0 0 0-1-2.22ZM9 4h6v2H9Zm6 16H9v-2h6Zm1-5a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}