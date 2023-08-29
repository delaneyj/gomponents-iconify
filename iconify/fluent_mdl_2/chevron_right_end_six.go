package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightEndSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m249 7l1016 1017L249 2041L7 1799l776-775L7 249L249 7zm1799-7v2048h-341V0h341z"/>`),
		g.Group(children),
	)
}