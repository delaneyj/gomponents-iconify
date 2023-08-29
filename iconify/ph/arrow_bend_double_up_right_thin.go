package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDoubleUpRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.83 106.83l-48 48a4 4 0 0 1-5.66-5.66L218.34 104l-45.17-45.17a4 4 0 0 1 5.66-5.66l48 48a4 4 0 0 1 0 5.66Zm-48-5.66l-48-48a4 4 0 1 0-5.66 5.66L166.34 100H128A100.11 100.11 0 0 0 28 200a4 4 0 0 0 8 0a92.1 92.1 0 0 1 92-92h38.34l-41.17 41.17a4 4 0 0 0 5.66 5.66l48-48a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}