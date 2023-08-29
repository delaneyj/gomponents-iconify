package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronUpEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 1099l762 762l-181 181l-581-581l-581 581l-181-181l762-762zm-581-81L262 837l762-762l762 762l-181 181l-581-581l-581 581z"/>`),
		g.Group(children),
	)
}