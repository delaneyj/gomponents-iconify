package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m91 0l1024 1024L91 2048l-91-91l933-933L0 91L91 0zm896 0l1024 1024L987 2048l-91-91l933-933L896 91l91-91z"/>`),
		g.Group(children),
	)
}