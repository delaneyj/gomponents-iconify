package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewArrayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 7v9h-5V7h5m6-2h-3v13h3V5m-4 0H8v13h9V5M7 5H4v13h3V5Z"/>`),
		g.Group(children),
	)
}