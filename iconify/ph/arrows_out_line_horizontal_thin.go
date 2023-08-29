package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutLineHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 40v176a4 4 0 0 1-8 0V40a4 4 0 0 1 8 0Zm-36 84H25.66l25.17-25.17a4 4 0 0 0-5.66-5.66l-32 32a4 4 0 0 0 0 5.66l32 32a4 4 0 1 0 5.66-5.66L25.66 132H96a4 4 0 0 0 0-8Zm146.83 1.17l-32-32a4 4 0 0 0-5.66 5.66L230.34 124H160a4 4 0 0 0 0 8h70.34l-25.17 25.17a4 4 0 0 0 5.66 5.66l32-32a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}