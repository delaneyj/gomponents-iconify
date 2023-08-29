package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownEndSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1265L7 249L249 7l775 776L1799 7l242 242l-1017 1016zm1024 442v341H0v-341h2048z"/>`),
		g.Group(children),
	)
}