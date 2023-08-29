package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendRightUpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204.24 84.24a6 6 0 0 1-8.48 0L158 46.49V128A102.12 102.12 0 0 1 56 230a6 6 0 0 1 0-12a90.1 90.1 0 0 0 90-90V46.49l-37.76 37.75a6 6 0 0 1-8.48-8.48l48-48a6 6 0 0 1 8.48 0l48 48a6 6 0 0 1 0 8.48Z"/>`),
		g.Group(children),
	)
}