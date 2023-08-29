package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 600C1200 268.641 931.359 0 600 0S0 268.641 0 600s268.641 600 600 600s600-268.641 600-600zm-297.141-66.656L600 987.656L297.141 533.344h163.184V212.32h279.352v321.023h163.182z"/>`),
		g.Group(children),
	)
}