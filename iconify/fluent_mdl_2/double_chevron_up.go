package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1957L1024 933l1024 1024l-91 91l-933-933l-933 933l-91-91zM1024 219L91 1152l-91-91L1024 37l1024 1024l-91 91l-933-933z"/>`),
		g.Group(children),
	)
}