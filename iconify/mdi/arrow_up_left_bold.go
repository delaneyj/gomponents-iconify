package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 21C9.36 21 6 17.64 6 13.5V11H2l6-7l6 7h-4v2.5c0 1.93 1.57 3.5 3.5 3.5H21v4h-7.5Z"/>`),
		g.Group(children),
	)
}