package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronLeftEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1461 1024l581 581l-181 181l-762-762l762-762l181 181l-581 581zm-443-581l-581 581l581 581l-181 181l-762-762l762-762l181 181z"/>`),
		g.Group(children),
	)
}