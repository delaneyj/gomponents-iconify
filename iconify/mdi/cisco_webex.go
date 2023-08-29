package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CiscoWebex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 3a9 9 0 0 1 9 9a9 9 0 0 1-9 9a9 9 0 0 1-9-9a9 9 0 0 1 9-9M5.94 8.5a6.988 6.988 0 0 0 2.56 9.56c3.35 1.94 10.35-10.19 7-12.12A6.988 6.988 0 0 0 5.94 8.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}