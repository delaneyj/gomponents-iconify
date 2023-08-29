package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 10.18L2.39 4.52l-.89.87l5.59 5.71a1.18 1.18 0 0 0 .86.39a1.13 1.13 0 0 0 .85-.39l5.7-5.7l-.88-.89z"/>`),
		g.Group(children),
	)
}