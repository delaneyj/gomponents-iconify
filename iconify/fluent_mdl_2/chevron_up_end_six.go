package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpEndSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 783l1017 1016l-242 242l-775-776l-775 776L7 1799L1024 783zM2048 0v341H0V0h2048z"/>`),
		g.Group(children),
	)
}