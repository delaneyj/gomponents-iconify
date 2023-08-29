package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronSmallUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.582 12.141a.695.695 0 0 1-.978 0a.68.68 0 0 1 0-.969l3.908-3.83a.697.697 0 0 1 .979 0l3.908 3.83a.68.68 0 0 1 0 .969a.697.697 0 0 1-.979 0L10 9l-3.418 3.141z"/>`),
		g.Group(children),
	)
}