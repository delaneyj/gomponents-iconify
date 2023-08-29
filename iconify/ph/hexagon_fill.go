package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 80.18v95.64a16 16 0 0 1-8.32 14l-88 48.17a15.88 15.88 0 0 1-15.36 0l-88-48.17a16 16 0 0 1-8.32-14V80.18a16 16 0 0 1 8.32-14l88-48.17a15.88 15.88 0 0 1 15.36 0l88 48.17a16 16 0 0 1 8.32 14Z"/>`),
		g.Group(children),
	)
}