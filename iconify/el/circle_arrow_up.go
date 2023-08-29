package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 600c0 331.359-268.641 600-600 600S0 931.359 0 600S268.641 0 600 0s600 268.641 600 600zm-297.141 66.633L600 212.32L297.141 666.633h163.184v321.023h279.352V666.633h163.182z"/>`),
		g.Group(children),
	)
}