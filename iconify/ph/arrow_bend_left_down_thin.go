package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendLeftDownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 32a4 4 0 0 1-4 4a92.1 92.1 0 0 0-92 92v86.34l41.17-41.17a4 4 0 0 1 5.66 5.66l-48 48a4 4 0 0 1-5.66 0l-48-48a4 4 0 0 1 5.66-5.66L100 214.34V128A100.11 100.11 0 0 1 200 28a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}