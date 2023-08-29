package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthernetCable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3v4h2V3h-2M8 4v7h8V4h-2v4h-4V4H8m2 8v10h4V12h-4Z"/>`),
		g.Group(children),
	)
}