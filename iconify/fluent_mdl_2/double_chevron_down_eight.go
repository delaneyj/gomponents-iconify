package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronDownEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 949L262 187L443 6l581 581L1605 6l181 181l-762 762zm581 81l181 181l-762 762l-762-762l181-181l581 581l581-581z"/>`),
		g.Group(children),
	)
}