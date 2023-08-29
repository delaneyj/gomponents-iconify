package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftEndSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1265 1024l776 775l-242 242L783 1024L1799 7l242 242l-776 775zM0 0h341v2048H0V0z"/>`),
		g.Group(children),
	)
}