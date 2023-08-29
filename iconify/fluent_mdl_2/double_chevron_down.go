package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 91L1024 1115L0 91L91 0l933 933L1957 0l91 91zM1024 1829l933-933l91 91l-1024 1024L0 987l91-91l933 933z"/>`),
		g.Group(children),
	)
}