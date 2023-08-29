package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.2 7l-4-4H21v4H10.2M20 8h-8.8l8.8 8.8V8m0 11.35v-.01L8.66 8l-1-1l-5.27-5.27L1.11 3L3 4.89V7h2.11l1 1H4v13h15.11l1.73 1.73l1.27-1.27L20 19.35Z"/>`),
		g.Group(children),
	)
}